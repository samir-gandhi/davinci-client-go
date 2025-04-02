package davinci

import (
	"encoding/json"
	"maps"
)

type _FlowVariableFields FlowVariableFields
type FlowVariableFields struct {
	AdditionalProperties map[string]any `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Type                 *string        `json:"type,omitempty" davinci:"type,config,omitempty"`
	DisplayName          *string        `json:"displayName,omitempty" davinci:"displayName,flowvariables,omitempty"`
	Mutable              *bool          `json:"mutable,omitempty" davinci:"mutable,flowvariables,omitempty"`
	Value                interface{}    `json:"value,omitempty" davinci:"value,flowvariables,omitempty"`
	Min                  *int32         `json:"min,omitempty" davinci:"min,flowvariables,omitempty"`
	Max                  *int32         `json:"max,omitempty" davinci:"max,flowvariables,omitempty"`
}

func (o FlowVariableFields) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o FlowVariableFields) ToMap() (map[string]any, error) {

	result := map[string]any{}

	if o.Type != nil {
		result["type"] = o.Type
	}

	if o.DisplayName != nil {
		result["displayName"] = o.DisplayName
	}

	if o.Mutable != nil {
		result["mutable"] = o.Mutable
	}

	if o.Value != nil {
		result["value"] = o.Value
	}

	if o.Min != nil {
		result["min"] = o.Min
	}

	if o.Max != nil {
		result["max"] = o.Max
	}

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *FlowVariableFields) UnmarshalJSON(bytes []byte) (err error) {
	varFlowVariableFields := _FlowVariableFields{}

	if err = json.Unmarshal(bytes, &varFlowVariableFields); err == nil {
		*o = FlowVariableFields(varFlowVariableFields)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "mutable")
		delete(additionalProperties, "value")
		delete(additionalProperties, "min")
		delete(additionalProperties, "max")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
