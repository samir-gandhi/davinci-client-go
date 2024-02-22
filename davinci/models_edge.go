package davinci

import "encoding/json"

var (
	_ DaVinciExportModel = Edge{}
)

type _Edge Edge
type Edge struct {
	AdditionalProperties map[string]interface{} `davinci:"-,unmapped"` // used to capture all other properties that are not explicitly defined in the model
	Data                 *Data                  `davinci:"data,unmapped,omitempty"`
	Position             *Position              `davinci:"position,unmapped,omitempty"`
	Group                *string                `davinci:"group,unmapped"`
	Removed              *bool                  `davinci:"removed,unmapped"`
	Selected             *bool                  `davinci:"selected,unmapped"`
	Selectable           *bool                  `davinci:"selectable,unmapped"`
	Locked               *bool                  `davinci:"locked,unmapped"`
	Grabbable            *bool                  `davinci:"grabbable,unmapped"`
	Pannable             *bool                  `davinci:"pannable,unmapped"`
	Classes              *string                `davinci:"classes,unmapped"`
}

func (o Edge) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Edge) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

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

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *Edge) UnmarshalJSON(bytes []byte) (err error) {
	varEdge := _Edge{}

	if err = json.Unmarshal(bytes, &varEdge); err == nil {
		*o = Edge(varEdge)
	}

	additionalProperties := make(map[string]interface{})

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

// DesignerCuesFields implements DaVinciExportModel.
func (o Edge) DesignerCuesFields() []string {
	return []string{
		"Group",
		"Removed",
		"Selected",
		"Selectable",
		"Locked",
		"Grabbable",
		"Pannable",
	}
}

// EnvironmentMetadataFields implements DaVinciExportModel.
func (o Edge) EnvironmentMetadataFields() []string {
	return []string{}
}

// FlowConfigFields implements DaVinciExportModel.
func (o Edge) FlowConfigFields() []string {
	return []string{
		"Classes",
	}
}

// FlowMetadataFields implements DaVinciExportModel.
func (o Edge) FlowMetadataFields() []string {
	return []string{}
}

// VersionMetadataFields implements DaVinciExportModel.
func (o Edge) VersionMetadataFields() []string {
	return []string{}
}

// SetAdditionalProperties implements DaVinciExportModel.
func (o Edge) SetAdditionalProperties(v map[string]interface{}) {
	o.AdditionalProperties = v
}
