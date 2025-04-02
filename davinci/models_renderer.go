package davinci

import (
	"encoding/json"
	"maps"
)

type _Renderer Renderer
type Renderer struct {
	AdditionalProperties map[string]any `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Name                 *string        `json:"name,omitempty" davinci:"name,designercue,omitempty"`
}

func (o Renderer) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Renderer) ToMap() (map[string]any, error) {

	result := map[string]any{}

	if o.Name != nil {
		result["name"] = o.Name
	}

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *Renderer) UnmarshalJSON(bytes []byte) (err error) {
	varRenderer := _Renderer{}

	if err = json.Unmarshal(bytes, &varRenderer); err == nil {
		*o = Renderer(varRenderer)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
