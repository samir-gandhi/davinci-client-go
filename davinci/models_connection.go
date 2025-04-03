package davinci

import (
	"encoding/json"
	"maps"
)

type _Connection Connection
type Connection struct {
	AdditionalProperties map[string]any                `json:"additionalProperties,omitempty"`
	CustomerID           *string                       `json:"customerId,omitempty"`
	ConnectorID          *string                       `json:"connectorId,omitempty"`
	Name                 *string                       `json:"name,omitempty"`
	CreatedDate          *EpochTime                    `json:"createdDate,omitempty"`
	Properties           map[string]ConnectionProperty `json:"properties,omitempty"`
	UpdatedDate          *EpochTime                    `json:"updatedDate,omitempty"`
	ConnectionID         *string                       `json:"connectionId,omitempty"`
	CompanyID            *string                       `json:"companyId,omitempty"`
}

func (o Connection) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Connection) ToMap() (map[string]any, error) {

	result := map[string]any{}

	if o.CustomerID != nil {
		result["customerId"] = o.CustomerID
	}

	if o.ConnectorID != nil {
		result["connectorId"] = o.ConnectorID
	}

	if o.Name != nil {
		result["name"] = o.Name
	}

	if o.CreatedDate != nil {
		result["createdDate"] = o.CreatedDate
	}

	if o.Properties != nil {
		result["properties"] = o.Properties
	}

	if o.UpdatedDate != nil {
		result["updatedDate"] = o.UpdatedDate
	}

	if o.ConnectionID != nil {
		result["connectionId"] = o.ConnectionID
	}

	if o.CompanyID != nil {
		result["companyId"] = o.CompanyID
	}

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *Connection) UnmarshalJSON(bytes []byte) (err error) {
	varConnection := _Connection{}

	if err = json.Unmarshal(bytes, &varConnection); err == nil {
		*o = Connection(varConnection)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "customerId")
		delete(additionalProperties, "connectorId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "createdDate")
		delete(additionalProperties, "properties")
		delete(additionalProperties, "updatedDate")
		delete(additionalProperties, "connectionId")
		delete(additionalProperties, "companyId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
