package davinci

import (
	"encoding/json"
	"fmt"
	"reflect"
	"slices"
	"strings"
)

var _ ValueEncoder = &StructCodec{}
var _ ValueDecoder = &StructCodec{}

type StructCodec struct {
	dCtx *DecoderContext
	eCtx *EncoderContext
}

func NewStructDecoder(dCtx *DecoderContext) ValueDecoder {
	return StructCodec{
		dCtx: dCtx,
	}
}

func NewStructEncoder(eCtx *EncoderContext) ValueEncoder {
	return StructCodec{
		eCtx: eCtx,
	}
}

func (StructCodec) String() string {
	return "davinci.StructCodec"
}

func (d StructCodec) DecodeValue(data []byte, v reflect.Value) error {
	if !v.IsValid() || !v.CanSet() || v.Kind() != reflect.Struct {
		return fmt.Errorf("invalid struct value to decode")
	}

	mappedFields := make([]string, 0)

	typ := v.Type()

	// Unmarshal the data into a map[string]any to work with.
	var tempMap map[string]any
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	var unmappedPropertiesField reflect.Value

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if typ.Field(i).Anonymous {
			// decode the value into the struct field
			if err := d.dCtx.Decode(data, field.Addr().Interface()); err != nil {
				return err
			}

			continue
		}

		fieldType := typ.Field(i)
		tagValue := fieldType.Tag.Get("davinci")

		// Split the tag value into JSON field name and purpose
		tagParts := strings.Split(tagValue, ",")
		if len(tagParts) < 2 {
			continue // Skip fields without a proper tag
		}

		jsonFieldName, fieldPurpose := tagParts[0], tagParts[1]

		mappedFields = append(mappedFields, jsonFieldName)

		if !slices.Contains([]string{"*", "-", "designercue", "environmentmetadata", "config", "flowmetadata", "flowvariables", "versionmetadata", "unmappedproperties"}, fieldPurpose) {
			return fmt.Errorf("davinci export field purpose %s is not recognised", fieldPurpose)
		}

		if fieldPurpose == "-" {
			continue
		}

		if fieldPurpose == "unmappedproperties" {
			if unmappedPropertiesField.Kind() != reflect.Invalid {
				return fmt.Errorf("multiple fields with purpose 'unmappedproperties' found")
			}

			if field.Kind() != reflect.Map || field.Type().Key().Kind() != reflect.String || field.Type().Elem().Kind() != reflect.Interface {
				return fmt.Errorf("field with purpose 'unmappedproperties' must be a map[string]any")
			}

			unmappedPropertiesField = field
		}

		if (fieldPurpose == "designercue" && !d.dCtx.Opts.IgnoreDesignerCues) ||
			(fieldPurpose == "environmentmetadata" && !d.dCtx.Opts.IgnoreEnvironmentMetadata) ||
			(fieldPurpose == "config" && !d.dCtx.Opts.IgnoreConfig) ||
			(fieldPurpose == "flowmetadata" && !d.dCtx.Opts.IgnoreFlowMetadata) ||
			(fieldPurpose == "flowvariables" && !d.dCtx.Opts.IgnoreFlowVariables) ||
			(fieldPurpose == "versionmetadata" && !d.dCtx.Opts.IgnoreVersionMetadata) ||
			fieldPurpose == "*" {

			if jsonValue, ok := tempMap[jsonFieldName]; ok {
				// Convert the map value to a JSON byte slice
				jsonValueBytes, err := json.Marshal(jsonValue)
				if err != nil {
					return err
				}

				if len(tagParts) > 2 && tagParts[2] == "omitempty" && string(jsonValueBytes) == "null" {
					continue
				}

				// Decode the value into the struct field
				if err := d.dCtx.Decode(jsonValueBytes, field.Addr().Interface()); err != nil {
					return err
				}
			}

			continue
		}
	}

	//Deal with additional / unmapped properties
	if !d.dCtx.Opts.IgnoreUnmappedProperties && unmappedPropertiesField.Kind() != reflect.Invalid {
		for _, mappedField := range mappedFields {
			delete(tempMap, mappedField)
		}

		if unmappedPropertiesField.IsValid() && len(tempMap) > 0 {
			unmappedPropertiesField.Set(reflect.ValueOf(tempMap))
		}

	}

	return nil
}

func (d StructCodec) EncodeValue(v reflect.Value) ([]byte, error) {

	var encodedFields []string

	var unmappedPropertiesField reflect.Value

	encodedAnonymousFields, err := d.encodePartValue(v, &unmappedPropertiesField)
	if err != nil {
		return nil, err
	}

	encodedFields = append(encodedFields, encodedAnonymousFields...)

	// Combine all the encoded fields into a single JSON object
	result := fmt.Sprintf("{%s}", strings.Join(encodedFields, ","))

	// Convert the JSON object to a byte slice
	resultBytes := []byte(result)

	// Unmarshal the data into a map[string]any to work with.
	var tempMap map[string]any
	if err := json.Unmarshal(resultBytes, &tempMap); err != nil {
		return nil, err
	}

	//Deal with additional / unmapped properties
	if !d.eCtx.Opts.IgnoreUnmappedProperties && unmappedPropertiesField.Kind() != reflect.Invalid {
		for k, v := range unmappedPropertiesField.Interface().(map[string]any) {
			tempMap[k] = v
		}
	}

	// Convert the map value to a JSON byte slice
	return json.Marshal(tempMap)
}

func (d StructCodec) encodePartValue(v reflect.Value, unmappedPropertiesField *reflect.Value) ([]string, error) {
	if !v.IsValid() || v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("invalid struct value to encode")
	}

	// Iterate over the struct fields, encoding each one by looking at the `davinci` tag, which tells the routine what name to give the field in the resulting byte slice
	var encodedFields []string

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)

		if fieldType.Anonymous {
			// encode the value into the struct field
			encodedAnonymousFields, err := d.encodePartValue(field, unmappedPropertiesField)
			if err != nil {
				return nil, err
			}

			encodedFields = append(encodedFields, encodedAnonymousFields...)

			continue
		}

		tagValue := fieldType.Tag.Get("davinci")
		tagParts := strings.Split(tagValue, ",")
		if len(tagParts) < 2 {
			continue
		}

		jsonFieldName := tagParts[0]
		fieldPurpose := tagParts[1]

		if !slices.Contains([]string{"*", "-", "designercue", "environmentmetadata", "config", "flowmetadata", "flowvariables", "versionmetadata", "unmappedproperties"}, fieldPurpose) {
			return nil, fmt.Errorf("davinci export field purpose %s is not recognised", fieldPurpose)
		}

		if fieldPurpose == "-" {
			continue
		}

		if fieldPurpose == "unmappedproperties" {
			if unmappedPropertiesField.Kind() != reflect.Invalid {
				return nil, fmt.Errorf("multiple fields with purpose 'unmappedproperties' found")
			}

			if field.Kind() != reflect.Map || field.Type().Key().Kind() != reflect.String || field.Type().Elem().Kind() != reflect.Interface {
				return nil, fmt.Errorf("field with purpose 'unmappedproperties' must be a map[string]any")
			}

			*unmappedPropertiesField = field
		}

		if (fieldPurpose == "designercue" && !d.eCtx.Opts.IgnoreDesignerCues) ||
			(fieldPurpose == "environmentmetadata" && !d.eCtx.Opts.IgnoreEnvironmentMetadata) ||
			(fieldPurpose == "config" && !d.eCtx.Opts.IgnoreConfig) ||
			(fieldPurpose == "flowmetadata" && !d.eCtx.Opts.IgnoreFlowMetadata) ||
			(fieldPurpose == "flowvariables" && !d.eCtx.Opts.IgnoreFlowVariables) ||
			(fieldPurpose == "versionmetadata" && !d.eCtx.Opts.IgnoreVersionMetadata) ||
			fieldPurpose == "*" {

			fieldValue := field.Interface()

			if fieldValue == nil || (reflect.TypeOf(fieldValue).Kind() == reflect.Ptr && reflect.ValueOf(fieldValue).IsNil()) {
				continue
			}

			// Encode the field value to JSON bytes
			jsonValueBytes, err := d.eCtx.Encode(fieldValue)
			if err != nil {
				return nil, err
			}

			// Add the encoded field to the resulting byte slice
			encodedFields = append(encodedFields, fmt.Sprintf(`"%s":%s`, jsonFieldName, string(jsonValueBytes)))
		}
	}

	return encodedFields, nil
}
