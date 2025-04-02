package davinci

import (
	"encoding/json"
	"maps"
)

type _SubFlowVersionID SubFlowVersionID
type SubFlowVersionID struct {
	AdditionalProperties map[string]any         `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Value                *SubFlowVersionIDValue `json:"value,omitempty" davinci:"value,*,omitempty"`
}

func (o SubFlowVersionID) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o SubFlowVersionID) ToMap() (map[string]any, error) {

	result := map[string]any{}

	if o.Value != nil {
		result["value"] = o.Value
	}

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *SubFlowVersionID) UnmarshalJSON(bytes []byte) (err error) {
	varSubFlowVersionID := _SubFlowVersionID{}

	if err = json.Unmarshal(bytes, &varSubFlowVersionID); err == nil {
		*o = SubFlowVersionID(varSubFlowVersionID)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
