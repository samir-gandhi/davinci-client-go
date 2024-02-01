package davinci

import "encoding/json"

type _Pan Pan
type Pan struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	X                    float64                `json:"x,omitempty"`
	Y                    float64                `json:"y,omitempty"`
}

func (o Pan) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Pan) ToMap() (map[string]interface{}, error) {

	// Marshal and unmarshal the metadata
	jsonPan, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err = json.Unmarshal(jsonPan, &result); err != nil {
		return nil, err
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *Pan) UnmarshalJSON(bytes []byte) (err error) {
	varPan := _Pan{}

	if err = json.Unmarshal(bytes, &varPan); err == nil {
		*o = Pan(varPan)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
