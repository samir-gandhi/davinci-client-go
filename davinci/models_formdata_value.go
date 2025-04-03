package davinci

import (
	"encoding/json"
	"maps"
)

type _FormDataValue FormDataValue
type FormDataValue struct {
	AdditionalProperties map[string]any `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Key                  *string        `json:"key,omitempty" davinci:"key,config,omitempty"`
	Value                *string        `json:"value,omitempty" davinci:"value,config,omitempty"`
}

func (o FormDataValue) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o FormDataValue) ToMap() (map[string]any, error) {

	result := map[string]any{}

	if o.Key != nil {
		result["key"] = o.Key
	}

	if o.Value != nil {
		result["value"] = o.Value
	}

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *FormDataValue) UnmarshalJSON(bytes []byte) (err error) {
	varFormDataValue := _FormDataValue{}

	if err = json.Unmarshal(bytes, &varFormDataValue); err == nil {
		*o = FormDataValue(varFormDataValue)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
