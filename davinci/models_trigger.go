package davinci

import (
	"encoding/json"
	"maps"
)

type _Trigger Trigger
type Trigger struct {
	AdditionalProperties map[string]any `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	TriggerType          *string        `json:"type,omitempty" davinci:"type,config,omitempty"`
}

func (o Trigger) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Trigger) ToMap() (map[string]any, error) {

	result := map[string]any{}

	if o.TriggerType != nil {
		result["type"] = o.TriggerType
	}

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *Trigger) UnmarshalJSON(bytes []byte) (err error) {
	varTrigger := _Trigger{}

	if err = json.Unmarshal(bytes, &varTrigger); err == nil {
		*o = Trigger(varTrigger)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
