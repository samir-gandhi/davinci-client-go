package davinci

import (
	"encoding/json"
	"maps"
)

type _SaveFlowVariable SaveFlowVariable
type SaveFlowVariable struct {
	AdditionalProperties map[string]any `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Key                  *float64       `json:"key,omitempty" davinci:"key,config,omitempty"`
	Label                *string        `json:"label,omitempty" davinci:"label,config,omitempty"`
	Name                 string         `json:"name" davinci:"name,config"`
	Type                 string         `json:"type" davinci:"type,config"`
	Value                interface{}    `json:"value,omitempty" davinci:"value,config,omitempty"`
	NameDefault          *string        `json:"nameDefault,omitempty" davinci:"nameDefault,config,omitempty"`
}

func (o SaveFlowVariable) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o SaveFlowVariable) ToMap() (map[string]any, error) {

	result := map[string]any{}

	result["key"] = o.Key
	result["label"] = o.Label
	result["name"] = o.Name
	result["type"] = o.Type
	result["value"] = o.Value
	result["nameDefault"] = o.NameDefault

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *SaveFlowVariable) UnmarshalJSON(bytes []byte) (err error) {
	varSaveFlowVariable := _SaveFlowVariable{}

	if err = json.Unmarshal(bytes, &varSaveFlowVariable); err == nil {
		*o = SaveFlowVariable(varSaveFlowVariable)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "label")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		delete(additionalProperties, "nameDefault")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
