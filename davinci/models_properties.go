package davinci

import (
	"encoding/json"
	"maps"
)

type _Properties Properties
type Properties struct {
	AdditionalProperties map[string]any     `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Form                 *Form              `json:"form,omitempty" davinci:"form,config,omitempty"`
	FormData             *FormData          `json:"formData,omitempty" davinci:"formData,config,omitempty"`
	SubFlowID            *SubFlowID         `json:"subFlowId,omitempty" davinci:"subFlowId,config,omitempty"`
	SubFlowVersionID     *SubFlowVersionID  `json:"subFlowVersionId,omitempty" davinci:"subFlowVersionId,config,omitempty"`
	SaveFlowVariables    *SaveFlowVariables `json:"saveFlowVariables,omitempty" davinci:"saveFlowVariables,*,omitempty"`
	SaveVariables        *SaveFlowVariables `json:"saveVariables,omitempty" davinci:"saveVariables,*,omitempty"`
}

func (o Properties) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Properties) ToMap() (map[string]any, error) {

	result := map[string]any{}

	if o.Form != nil {
		result["form"] = o.Form
	}

	if o.FormData != nil {
		result["formData"] = o.FormData
	}

	if o.SubFlowID != nil {
		result["subFlowId"] = o.SubFlowID
	}

	if o.SubFlowVersionID != nil {
		result["subFlowVersionId"] = o.SubFlowVersionID
	}

	if o.SaveFlowVariables != nil {
		result["saveFlowVariables"] = o.SaveFlowVariables
	}

	if o.SaveVariables != nil {
		result["saveVariables"] = o.SaveVariables
	}

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *Properties) UnmarshalJSON(bytes []byte) (err error) {
	varProperties := _Properties{}

	if err = json.Unmarshal(bytes, &varProperties); err == nil {
		*o = Properties(varProperties)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "form")
		delete(additionalProperties, "formData")
		delete(additionalProperties, "subFlowId")
		delete(additionalProperties, "subFlowVersionId")
		delete(additionalProperties, "saveFlowVariables")
		delete(additionalProperties, "saveVariables")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
