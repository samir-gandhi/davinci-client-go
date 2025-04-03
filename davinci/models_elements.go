package davinci

import (
	"encoding/json"
	"maps"
)

type _Elements Elements
type Elements struct {
	AdditionalProperties map[string]any `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Nodes                []Node         `json:"nodes,omitempty" davinci:"nodes,*,omitempty"`
	Edges                []Edge         `json:"edges,omitempty" davinci:"edges,*,omitempty"`
}

func (o Elements) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Elements) ToMap() (map[string]any, error) {

	result := map[string]any{}

	if o.Nodes != nil {
		result["nodes"] = o.Nodes
	} else {
		result["nodes"] = make([]Node, 0)
	}

	if o.Edges != nil {
		result["edges"] = o.Edges
	} else {
		result["edges"] = make([]Edge, 0)
	}

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *Elements) UnmarshalJSON(bytes []byte) (err error) {
	varElements := _Elements{}

	if err = json.Unmarshal(bytes, &varElements); err == nil {
		*o = Elements(varElements)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "nodes")
		delete(additionalProperties, "edges")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
