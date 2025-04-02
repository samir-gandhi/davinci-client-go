package davinci

import (
	"encoding/json"
	"maps"
)

type _FlowVariable FlowVariable
type FlowVariable struct {
	AdditionalProperties map[string]any      `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	ID                   *string             `json:"id,omitempty" davinci:"id,environmentmetadata,omitempty"`
	CompanyID            *string             `json:"companyId,omitempty" davinci:"companyId,environmentmetadata,omitempty"`
	Context              *string             `json:"context,omitempty" davinci:"context,config,omitempty"`
	CreatedDate          *EpochTime          `json:"createdDate,omitempty" davinci:"createdDate,flowvariables,omitempty"`
	CustomerID           *string             `json:"customerId,omitempty" davinci:"customerId,environmentmetadata,omitempty"`
	Fields               *FlowVariableFields `json:"fields,omitempty" davinci:"fields,*,omitempty"`
	FlowID               *string             `json:"flowId,omitempty" davinci:"flowId,environmentmetadata,omitempty"`
	Key                  *float64            `json:"key,omitempty" davinci:"key,flowmetadata,omitempty"`
	Label                *string             `json:"label,omitempty" davinci:"label,flowvariables,omitempty"`
	Name                 string              `json:"name" davinci:"name,config"`
	Type                 string              `json:"type" davinci:"type,config"`
	UpdatedDate          *EpochTime          `json:"updatedDate,omitempty" davinci:"updatedDate,flowvariables,omitempty"`
	Value                interface{}         `json:"value,omitempty" davinci:"value,flowvariables,omitempty"`
	Visibility           *string             `json:"visibility,omitempty" davinci:"visibility,flowvariables,omitempty"`
}

func (o FlowVariable) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o FlowVariable) ToMap() (map[string]any, error) {

	result := map[string]any{}

	result["id"] = o.ID
	result["companyId"] = o.CompanyID
	result["context"] = o.Context
	result["createdDate"] = o.CreatedDate
	result["customerId"] = o.CustomerID
	result["fields"] = o.Fields
	result["flowId"] = o.FlowID
	result["key"] = o.Key
	result["label"] = o.Label
	result["name"] = o.Name
	result["type"] = o.Type

	if o.UpdatedDate != nil {
		result["updatedDate"] = o.UpdatedDate
	}

	result["value"] = o.Value
	result["visibility"] = o.Visibility

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *FlowVariable) UnmarshalJSON(bytes []byte) (err error) {
	varFlowVariable := _FlowVariable{}

	if err = json.Unmarshal(bytes, &varFlowVariable); err == nil {
		*o = FlowVariable(varFlowVariable)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "companyId")
		delete(additionalProperties, "context")
		delete(additionalProperties, "createdDate")
		delete(additionalProperties, "customerId")
		delete(additionalProperties, "fields")
		delete(additionalProperties, "flowId")
		delete(additionalProperties, "key")
		delete(additionalProperties, "label")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "updatedDate")
		delete(additionalProperties, "value")
		delete(additionalProperties, "visibility")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
