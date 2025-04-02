package davinci

import (
	"encoding/json"
	"maps"
)

type _Node Node
type Node struct {
	AdditionalProperties map[string]any `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Data                 *NodeData      `json:"data,omitempty" davinci:"data,*,omitempty"`
	Position             *Position      `json:"position,omitempty" davinci:"position,*,omitempty"`
	Group                string         `json:"group" davinci:"group,designercue"`
	Removed              bool           `json:"removed" davinci:"removed,designercue"`
	Selected             bool           `json:"selected" davinci:"selected,designercue"`
	Selectable           bool           `json:"selectable" davinci:"selectable,designercue"`
	Locked               bool           `json:"locked" davinci:"locked,designercue"`
	Grabbable            bool           `json:"grabbable" davinci:"grabbable,designercue"`
	Pannable             bool           `json:"pannable" davinci:"pannable,designercue"`
	Classes              string         `json:"classes" davinci:"classes,config"`
}

func (o Node) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Node) ToMap() (map[string]any, error) {

	result := map[string]any{}

	if o.Data != nil {
		result["data"] = o.Data
	}

	if o.Position != nil {
		result["position"] = o.Position
	}

	result["group"] = o.Group
	result["removed"] = o.Removed
	result["selected"] = o.Selected
	result["selectable"] = o.Selectable
	result["locked"] = o.Locked
	result["grabbable"] = o.Grabbable
	result["pannable"] = o.Pannable
	result["classes"] = o.Classes

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *Node) UnmarshalJSON(bytes []byte) (err error) {
	varNode := _Node{}

	if err = json.Unmarshal(bytes, &varNode); err == nil {
		*o = Node(varNode)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "position")
		delete(additionalProperties, "group")
		delete(additionalProperties, "removed")
		delete(additionalProperties, "selected")
		delete(additionalProperties, "selectable")
		delete(additionalProperties, "locked")
		delete(additionalProperties, "grabbable")
		delete(additionalProperties, "pannable")
		delete(additionalProperties, "classes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
