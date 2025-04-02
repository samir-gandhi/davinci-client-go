package davinci

import (
	"encoding/json"
	"maps"
)

type _Edge Edge
type Edge struct {
	AdditionalProperties map[string]any `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Data                 *Data          `json:"data,omitempty" davinci:"data,*,omitempty"`
	Position             *Position      `json:"position,omitempty" davinci:"position,*,omitempty"`
	Group                *string        `json:"group,omitempty" davinci:"group,designercue,omitempty"`
	Removed              *bool          `json:"removed,omitempty" davinci:"removed,designercue,omitempty"`
	Selected             *bool          `json:"selected,omitempty" davinci:"selected,designercue,omitempty"`
	Selectable           *bool          `json:"selectable,omitempty" davinci:"selectable,designercue,omitempty"`
	Locked               *bool          `json:"locked,omitempty" davinci:"locked,designercue,omitempty"`
	Grabbable            *bool          `json:"grabbable,omitempty" davinci:"grabbable,designercue,omitempty"`
	Pannable             *bool          `json:"pannable,omitempty" davinci:"pannable,designercue,omitempty"`
	Classes              *string        `json:"classes,omitempty" davinci:"classes,config,omitempty"`
}

func (o Edge) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Edge) ToMap() (map[string]any, error) {

	result := map[string]any{}

	if o.Data != nil {
		result["data"] = o.Data
	}

	if o.Position != nil {
		result["position"] = o.Position
	}

	if o.Group != nil {
		result["group"] = o.Group
	}

	if o.Removed != nil {
		result["removed"] = o.Removed
	}

	if o.Selected != nil {
		result["selected"] = o.Selected
	}

	if o.Selectable != nil {
		result["selectable"] = o.Selectable
	}

	if o.Locked != nil {
		result["locked"] = o.Locked
	}

	if o.Grabbable != nil {
		result["grabbable"] = o.Grabbable
	}

	if o.Pannable != nil {
		result["pannable"] = o.Pannable
	}

	if o.Classes != nil {
		result["classes"] = o.Classes
	}

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *Edge) UnmarshalJSON(bytes []byte) (err error) {
	varEdge := _Edge{}

	if err = json.Unmarshal(bytes, &varEdge); err == nil {
		*o = Edge(varEdge)
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
