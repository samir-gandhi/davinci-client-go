package davinci

import (
	"encoding/json"
	"maps"
)

type _SubFlowProperties SubFlowProperties
type SubFlowProperties struct {
	AdditionalProperties map[string]any     `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	SubFlowID            *SubFlowID         `json:"subFlowId,omitempty" davinci:"subFlowId,*,omitempty"`
	SubFlowProperties    *SubFlowProperties `json:"subFlowVersionId,omitempty" davinci:"subFlowVersionId,*,omitempty"`
}

func (o SubFlowProperties) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o SubFlowProperties) ToMap() (map[string]any, error) {

	result := map[string]any{}

	if o.SubFlowID != nil {
		result["subFlowId"] = o.SubFlowID
	}

	if o.SubFlowProperties != nil {
		result["subFlowVersionId"] = o.SubFlowProperties
	}

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *SubFlowProperties) UnmarshalJSON(bytes []byte) (err error) {
	varSubFlowProperties := _SubFlowProperties{}

	if err = json.Unmarshal(bytes, &varSubFlowProperties); err == nil {
		*o = SubFlowProperties(varSubFlowProperties)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "subFlowId")
		delete(additionalProperties, "subFlowVersionId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
