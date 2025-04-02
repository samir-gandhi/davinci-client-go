package davinci

import (
	"encoding/json"
	"maps"
)

type _SaveFlowVariables SaveFlowVariables
type SaveFlowVariables struct {
	AdditionalProperties map[string]any     `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Value                []SaveFlowVariable `json:"value,omitempty" davinci:"value,*,omitempty"`
}

func (o SaveFlowVariables) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o SaveFlowVariables) ToMap() (map[string]any, error) {

	result := map[string]any{}

	if o.Value != nil {
		result["value"] = o.Value
	}

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *SaveFlowVariables) UnmarshalJSON(bytes []byte) (err error) {
	varSaveFlowVariables := _SaveFlowVariables{}

	if err = json.Unmarshal(bytes, &varSaveFlowVariables); err == nil {
		*o = SaveFlowVariables(varSaveFlowVariables)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
