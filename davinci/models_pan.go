package davinci

import (
	"encoding/json"
	"maps"
)

type _Pan Pan
type Pan struct {
	AdditionalProperties map[string]any `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	X                    *float64       `json:"x,omitempty" davinci:"x,designercue,omitempty"`
	Y                    *float64       `json:"y,omitempty" davinci:"y,designercue,omitempty"`
}

func (o Pan) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Pan) ToMap() (map[string]any, error) {

	result := map[string]any{}

	if o.X != nil {
		result["x"] = o.X
	}

	if o.Y != nil {
		result["y"] = o.Y
	}

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *Pan) UnmarshalJSON(bytes []byte) (err error) {
	varPan := _Pan{}

	if err = json.Unmarshal(bytes, &varPan); err == nil {
		*o = Pan(varPan)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
