package davinci

import "encoding/json"

type _Data Data
type Data struct {
	AdditionalProperties map[string]interface{} `json:"-"` // used to capture all other properties that are not explicitly defined in the model
	CapabilityName       string                 `json:"capabilityName,omitempty"`
	ConnectionID         string                 `json:"connectionId,omitempty"`
	ConnectorID          string                 `json:"connectorId,omitempty"`
	ID                   string                 `json:"id,omitempty"`
	Label                string                 `json:"label,omitempty"`
	MultiValueSourceId   string                 `json:"multiValueSourceId,omitempty"`
	Name                 string                 `json:"name,omitempty"`
	NodeType             string                 `json:"nodeType,omitempty"`
	Properties           Properties             `json:"properties,omitempty"`
	Source               string                 `json:"source,omitempty"`
	Status               string                 `json:"status,omitempty"`
	Target               string                 `json:"target,omitempty"`
	Type                 string                 `json:"type,omitempty"`
}

func (o Data) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Data) ToMap() (map[string]interface{}, error) {

	// Marshal and unmarshal the metadata
	jsonData, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err = json.Unmarshal(jsonData, &result); err != nil {
		return nil, err
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *Data) UnmarshalJSON(bytes []byte) (err error) {
	varData := _Data{}

	if err = json.Unmarshal(bytes, &varData); err == nil {
		*o = Data(varData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "capabilityName")
		delete(additionalProperties, "connectionId")
		delete(additionalProperties, "connectorId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "label")
		delete(additionalProperties, "multiValueSourceId")
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
