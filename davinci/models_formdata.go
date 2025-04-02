package davinci

import (
	"encoding/json"
	"maps"
)

type _FormData FormData
type FormData struct {
	AdditionalProperties map[string]any  `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Value                []FormDataValue `json:"value,omitempty" davinci:"value,config,omitempty"`
}

func (o FormData) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o FormData) ToMap() (map[string]any, error) {

	result := map[string]any{}

	if o.Value != nil {
		result["value"] = o.Value
	}

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *FormData) UnmarshalJSON(bytes []byte) (err error) {
	varFormData := _FormData{}

	if err = json.Unmarshal(bytes, &varFormData); err == nil {
		*o = FormData(varFormData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
