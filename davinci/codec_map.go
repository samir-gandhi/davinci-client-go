package davinci

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

var _ ValueEncoder = &MapCodec{}
var _ ValueDecoder = &MapCodec{}

type MapCodec struct {
	dCtx *DecoderContext
	eCtx *EncoderContext
}

func NewMapDecoder(dCtx *DecoderContext) ValueDecoder {
	return MapCodec{
		dCtx: dCtx,
	}
}

func NewMapEncoder(eCtx *EncoderContext) ValueEncoder {
	return MapCodec{
		eCtx: eCtx,
	}
}

func (MapCodec) String() string {
	return "davinci.MapCodec"
}

func (d MapCodec) DecodeValue(data []byte, v reflect.Value) error {
	if !v.IsValid() || !v.CanSet() || v.Kind() != reflect.Map {
		return fmt.Errorf("invalid map value to decode")
	}

	// Unmarshal the data into a map
	var tempMap map[string]any
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	valType := v.Type()
	valKeyType := valType.Key()
	valElemType := valType.Elem()
	mapType := reflect.MapOf(valKeyType, valElemType)
	vMap := reflect.MakeMapWithSize(mapType, 0)

	for key, item := range tempMap {
		// Create a new value of the map element type
		elem := reflect.New(valElemType).Elem()

		// Convert the map value to a JSON byte map
		jsonValueBytes, err := json.Marshal(item)
		if err != nil {
			return err
		}

		// Decode the value into the new element
		if err := d.dCtx.Decode(jsonValueBytes, elem.Addr().Interface()); err != nil {
			return err
		}

		if elem.Kind() == reflect.Struct {
			// Check if the struct is nil, in which case we should skip it
			if reflect.DeepEqual(elem.Interface(), reflect.Zero(elem.Type()).Interface()) {
				continue
			}
		}

		// Append the new element to the map
		vMap.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(elem.Interface()))
	}

	// Set the map value
	v.Set(vMap)

	return nil
}

func (d MapCodec) EncodeValue(v reflect.Value) ([]byte, error) {

	if !v.IsValid() || v.Kind() != reflect.Map {
		return nil, fmt.Errorf("invalid map value to encode")
	}

	// Create a new map to hold the encoded values
	var encodedFields []string

	// Iterate over the map and encode each value
	for _, key := range v.MapKeys() {
		// Get the value at the current index
		val := v.MapIndex(key)

		// Encode the key
		jsonValueMapKeyBytes, err := d.eCtx.Encode(key.Interface())
		if err != nil {
			return nil, err
		}

		// Encode the value
		jsonValueMapValueBytes, err := d.eCtx.Encode(val.Interface())
		if err != nil {
			return nil, err
		}

		encodedFields = append(encodedFields, fmt.Sprintf(`%s: %s`, string(jsonValueMapKeyBytes), string(jsonValueMapValueBytes)))
	}

	result := fmt.Sprintf("{%s}", strings.Join(encodedFields, ","))

	return []byte(result), nil
}
