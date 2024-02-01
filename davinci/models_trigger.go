package davinci

import "encoding/json"

type _Trigger Trigger
type Trigger struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	TriggerType          string                 `json:"type,omitempty"`
}

func (o Trigger) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Trigger) ToMap() (map[string]interface{}, error) {

	// Marshal and unmarshal the metadata
	jsonTrigger, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err = json.Unmarshal(jsonTrigger, &result); err != nil {
		return nil, err
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *Trigger) UnmarshalJSON(bytes []byte) (err error) {
	varTrigger := _Trigger{}

	if err = json.Unmarshal(bytes, &varTrigger); err == nil {
		*o = Trigger(varTrigger)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
