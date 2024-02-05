package davinci

import "encoding/json"

type _Elements Elements
type Elements struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	Nodes                []Node                 `json:"nodes,omitempty"`
	Edges                []Edge                 `json:"edges,omitempty"`
}

func (o Elements) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Elements) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	result["nodes"] = o.Nodes
	result["edges"] = o.Edges

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *Elements) UnmarshalJSON(bytes []byte) (err error) {
	varElements := _Elements{}

	if err = json.Unmarshal(bytes, &varElements); err == nil {
		*o = Elements(varElements)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "nodes")
		delete(additionalProperties, "edges")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
