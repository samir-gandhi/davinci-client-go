package davinci

import (
	"encoding/json"
	"maps"
)

type _NodeData NodeData
type NodeData struct {
	AdditionalProperties map[string]any `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	CapabilityName       *string        `json:"capabilityName,omitempty" davinci:"capabilityName,config,omitempty"`
	ConnectionID         *string        `json:"connectionId,omitempty" davinci:"connectionId,config,omitempty"`
	ConnectorID          *string        `json:"connectorId,omitempty" davinci:"connectorId,config,omitempty"`
	ID                   *string        `json:"id,omitempty" davinci:"id,config,omitempty"`
	Label                *string        `json:"label,omitempty" davinci:"label,config,omitempty"`
	Name                 *string        `json:"name,omitempty" davinci:"name,config,omitempty"`
	NodeType             *string        `json:"nodeType,omitempty" davinci:"nodeType,config,omitempty"`
	Properties           *Properties    `json:"properties" davinci:"properties,*"`
	Source               *string        `json:"source,omitempty" davinci:"source,config,omitempty"`
	Status               *string        `json:"status,omitempty" davinci:"status,config,omitempty"`
	Target               *string        `json:"target,omitempty" davinci:"target,config,omitempty"`
	Type                 *string        `json:"type,omitempty" davinci:"type,config,omitempty"`
}

func (o NodeData) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o NodeData) ToMap() (map[string]any, error) {

	result := map[string]any{}

	if o.CapabilityName != nil {
		result["capabilityName"] = o.CapabilityName
	}

	if o.ConnectionID != nil {
		result["connectionId"] = o.ConnectionID
	}

	if o.ConnectorID != nil {
		result["connectorId"] = o.ConnectorID
	}

	if o.ID != nil {
		result["id"] = o.ID
	}

	if o.Label != nil {
		result["label"] = o.Label
	}

	if o.Name != nil {
		result["name"] = o.Name
	}

	if o.NodeType != nil {
		result["nodeType"] = o.NodeType
	}

	if o.Properties != nil {
		result["properties"] = o.Properties
	}

	if o.Source != nil {
		result["source"] = o.Source
	}

	if o.Status != nil {
		result["status"] = o.Status
	}

	if o.Source != nil {
		result["target"] = o.Target
	}

	if o.Type != nil {
		result["type"] = o.Type
	}

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *NodeData) UnmarshalJSON(bytes []byte) (err error) {
	varNodeData := _NodeData{}

	if err = json.Unmarshal(bytes, &varNodeData); err == nil {
		*o = NodeData(varNodeData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "capabilityName")
		delete(additionalProperties, "connectionId")
		delete(additionalProperties, "connectorId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "label")
		delete(additionalProperties, "name")
		delete(additionalProperties, "nodeType")
		delete(additionalProperties, "properties")
		delete(additionalProperties, "source")
		delete(additionalProperties, "status")
		delete(additionalProperties, "target")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
