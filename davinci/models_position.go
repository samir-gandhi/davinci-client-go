package davinci

import "encoding/json"

type _Position Position
type Position struct {
	AdditionalProperties map[string]interface{} `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	X                    *float64               `json:"x,omitempty" davinci:"x,designercue,omitempty"`
	Y                    *float64               `json:"y,omitempty" davinci:"y,designercue,omitempty"`
}

func (o Position) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Position) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	if o.X != nil {
		result["x"] = o.X
	}

	if o.Y != nil {
		result["y"] = o.Y
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *Position) UnmarshalJSON(bytes []byte) (err error) {
	varPosition := _Position{}

	if err = json.Unmarshal(bytes, &varPosition); err == nil {
		*o = Position(varPosition)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
