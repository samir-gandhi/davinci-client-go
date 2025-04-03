package davinci

import (
	"encoding/json"
	"maps"
)

type _OutputSchema OutputSchema
type OutputSchema struct {
	AdditionalProperties map[string]any `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Output               interface{}    `json:"output,omitempty" davinci:"output,config,omitempty"`
}

func (o OutputSchema) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o OutputSchema) ToMap() (map[string]any, error) {

	result := map[string]any{}

	if o.Output != nil {
		result["output"] = o.Output
	}

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *OutputSchema) UnmarshalJSON(bytes []byte) (err error) {
	varOutputSchema := _OutputSchema{}

	if err = json.Unmarshal(bytes, &varOutputSchema); err == nil {
		*o = OutputSchema(varOutputSchema)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "output")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
