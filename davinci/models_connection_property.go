package davinci

import (
	"encoding/json"
	"maps"
)

type _ConnectionProperty ConnectionProperty
type ConnectionProperty struct {
	AdditionalProperties map[string]any                `json:"additionalProperties,omitempty"`
	CompanyId            *string                       `json:"companyId,omitempty"`
	ConstructItems       []string                      `json:"constructItems,omitempty"`
	CreatedDate          *EpochTime                    `json:"createdDate,omitempty"`
	CustomerId           *string                       `json:"customerId,omitempty"`
	DisplayName          *string                       `json:"displayName,omitempty"`
	Info                 *string                       `json:"info,omitempty"`
	Placeholder          *string                       `json:"placeholder,omitempty"`
	PreferredControlType *string                       `json:"preferredControlType,omitempty"`
	Properties           map[string]ConnectionProperty `json:"properties,omitempty"`
	Required             *bool                         `json:"required,omitempty"`
	Sections             []string                      `json:"sections,omitempty"`
	Secure               *bool                         `json:"secure,omitempty"`
	Type                 *string                       `json:"type,omitempty"`
	Value                interface{}                   `json:"value,omitempty"`
}

func (o ConnectionProperty) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o ConnectionProperty) ToMap() (map[string]any, error) {

	result := map[string]any{}

	if o.CompanyId != nil {
		result["companyId"] = o.CompanyId
	}

	if o.ConstructItems != nil {
		result["constructItems"] = o.ConstructItems
	}

	if o.CreatedDate != nil {
		result["createdDate"] = o.CreatedDate
	}

	if o.CustomerId != nil {
		result["customerId"] = o.CustomerId
	}

	if o.DisplayName != nil {
		result["displayName"] = o.DisplayName
	}

	if o.Info != nil {
		result["info"] = o.Info
	}

	if o.Placeholder != nil {
		result["placeholder"] = o.Placeholder
	}

	if o.PreferredControlType != nil {
		result["preferredControlType"] = o.PreferredControlType
	}

	if o.Properties != nil {
		result["properties"] = o.Properties
	}

	if o.Required != nil {
		result["required"] = o.Required
	}

	if o.Sections != nil {
		result["sections"] = o.Sections
	}

	if o.Secure != nil {
		result["secure"] = o.Secure
	}

	if o.Type != nil {
		result["type"] = o.Type
	}

	if o.Value != nil {
		result["value"] = o.Value
	}

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *ConnectionProperty) UnmarshalJSON(bytes []byte) (err error) {
	varConnectionProperty := _ConnectionProperty{}

	if err = json.Unmarshal(bytes, &varConnectionProperty); err == nil {
		*o = ConnectionProperty(varConnectionProperty)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "companyId")
		delete(additionalProperties, "constructItems")
		delete(additionalProperties, "createdDate")
		delete(additionalProperties, "customerId")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "info")
		delete(additionalProperties, "placeholder")
		delete(additionalProperties, "preferredControlType")
		delete(additionalProperties, "properties")
		delete(additionalProperties, "required")
		delete(additionalProperties, "sections")
		delete(additionalProperties, "secure")
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
