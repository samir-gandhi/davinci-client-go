package davinci

import (
	"encoding/json"
	"maps"
)

type _SubFlowID SubFlowID
type SubFlowID struct {
	AdditionalProperties map[string]any `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Value                *SubFlowValue  `json:"value,omitempty" davinci:"value,*,omitempty"`
}

func (o SubFlowID) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o SubFlowID) ToMap() (map[string]any, error) {

	result := map[string]any{}

	if o.Value != nil {
		result["value"] = o.Value
	}

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *SubFlowID) UnmarshalJSON(bytes []byte) (err error) {
	varSubFlowID := _SubFlowID{}

	if err = json.Unmarshal(bytes, &varSubFlowID); err == nil {
		*o = SubFlowID(varSubFlowID)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
