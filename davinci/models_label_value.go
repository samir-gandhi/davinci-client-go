package davinci

import (
	"encoding/json"
	"maps"
)

type _LabelValue LabelValue
type LabelValue struct {
	AdditionalProperties map[string]any `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Label                *string        `json:"label,omitempty" davinci:"label,config,omitempty"`
	Value                *string        `json:"value,omitempty" davinci:"value,config,omitempty"`
}

func (o LabelValue) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o LabelValue) ToMap() (map[string]any, error) {

	result := map[string]any{}

	if o.Label != nil {
		result["label"] = o.Label
	}

	if o.Value != nil {
		result["value"] = o.Value
	}

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *LabelValue) UnmarshalJSON(bytes []byte) (err error) {
	varLabelValue := _LabelValue{}

	if err = json.Unmarshal(bytes, &varLabelValue); err == nil {
		*o = LabelValue(varLabelValue)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
