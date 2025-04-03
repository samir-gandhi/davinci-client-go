package davinci

import (
	"encoding/json"
	"maps"
)

type _Flow Flow
type Flow struct {
	AdditionalProperties map[string]any `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	FlowConfiguration
	FlowEnvironmentMetadata
	FlowMetadata
	FlowVersionMetadata
}

type FlowConfiguration struct {
	FlowUpdateConfiguration
	FlowColor *string `json:"flowColor,omitempty" davinci:"flowColor,designercue,omitempty"`
}

type FlowUpdateConfiguration struct {
	GraphData    *GraphData    `json:"graphData,omitempty" davinci:"graphData,*,omitempty"`
	InputSchema  []interface{} `json:"inputSchema,omitempty" davinci:"inputSchema,config,omitempty"`
	OutputSchema *OutputSchema `json:"outputSchema,omitempty" davinci:"outputSchema,*,omitempty"`
	Settings     *FlowSettings `json:"settings,omitempty" davinci:"settings,*,omitempty"`
	Trigger      *Trigger      `json:"trigger,omitempty" davinci:"trigger,*,omitempty"`
}

type FlowEnvironmentMetadata struct {
	CompanyID    string    `json:"companyId" davinci:"companyId,environmentmetadata"`
	CreatedDate  EpochTime `json:"createdDate" davinci:"createdDate,environmentmetadata"`
	CustomerID   string    `json:"customerId" davinci:"customerId,environmentmetadata"`
	FlowID       string    `json:"flowId" davinci:"flowId,environmentmetadata"`
	ParentFlowID *string   `json:"parentFlowId,omitempty" davinci:"parentFlowId,environmentmetadata,omitempty"`
}

type FlowMetadata struct {
	AuthTokenExpireIds   []interface{}  `json:"authTokenExpireIds" davinci:"authTokenExpireIds,flowmetadata"`
	Connections          []interface{}  `json:"connections,omitempty" davinci:"connections,flowmetadata,omitempty"`
	ConnectorIds         []string       `json:"connectorIds" davinci:"connectorIds,flowmetadata"`
	Description          *string        `json:"description,omitempty" davinci:"description,flowmetadata,omitempty"`
	EnabledGraphData     interface{}    `json:"enabledGraphData,omitempty" davinci:"enabledGraphData,flowmetadata,omitempty"`
	FunctionConnectionID interface{}    `json:"functionConnectionId,omitempty" davinci:"functionConnectionId,flowmetadata,omitempty"`
	InputSchemaCompiled  interface{}    `json:"inputSchemaCompiled,omitempty" davinci:"inputSchemaCompiled,flowmetadata,omitempty"`
	IsInputSchemaSaved   *bool          `json:"isInputSchemaSaved,omitempty" davinci:"isInputSchemaSaved,flowmetadata,omitempty"`
	IsOutputSchemaSaved  bool           `json:"isOutputSchemaSaved" davinci:"isOutputSchemaSaved,flowmetadata"`
	Name                 string         `json:"name" davinci:"name,flowmetadata"`
	Orx                  *string        `json:"orx,omitempty" davinci:"orx,flowmetadata,omitempty"`
	OutputSchemaCompiled *OutputSchema  `json:"outputSchemaCompiled,omitempty" davinci:"outputSchemaCompiled,*,omitempty"` //compiled is used in exported flow json, must be converted to JUST output when updating flow.
	Timeouts             interface{}    `json:"timeouts,omitempty" davinci:"timeouts,flowmetadata,omitempty"`
	Variables            []FlowVariable `json:"variables,omitempty" davinci:"variables,*,omitempty"`
}

type FlowVersionMetadata struct {
	CurrentVersion   *int32      `json:"currentVersion,omitempty" davinci:"currentVersion,versionmetadata,omitempty"`
	DeployedDate     *EpochTime  `json:"deployedDate,omitempty" davinci:"deployedDate,versionmetadata,omitempty"`
	FlowStatus       string      `json:"flowStatus" davinci:"flowStatus,versionmetadata"`
	PublishedVersion *int32      `json:"publishedVersion,omitempty" davinci:"publishedVersion,versionmetadata,omitempty"`
	SavedDate        EpochTime   `json:"savedDate" davinci:"savedDate,versionmetadata"`
	UpdatedDate      *EpochTime  `json:"updatedDate,omitempty" davinci:"updatedDate,versionmetadata,omitempty"`
	VersionID        int32       `json:"versionId" davinci:"versionId,versionmetadata"`
	VersionInfo      interface{} `json:"versionInfo,omitempty" davinci:"versionInfo,versionmetadata,omitempty"`
}

func (o *Flow) UnmarshalDavinci(bytes []byte, opts ExportCmpOpts) (err error) {
	varFlow := _Flow{}

	if err = Unmarshal(bytes, &varFlow, opts); err != nil {
		return err
	}

	*o = Flow(varFlow)

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err != nil {
		return err
	}

	delete(additionalProperties, "authTokenExpireIds")
	delete(additionalProperties, "companyId")
	delete(additionalProperties, "connections")
	delete(additionalProperties, "connectorIds")
	delete(additionalProperties, "createdDate")
	delete(additionalProperties, "currentVersion")
	delete(additionalProperties, "customerId")
	delete(additionalProperties, "deployedDate")
	delete(additionalProperties, "description")
	delete(additionalProperties, "enabledGraphData")
	delete(additionalProperties, "flowColor")
	delete(additionalProperties, "flowId")
	delete(additionalProperties, "flowStatus")
	delete(additionalProperties, "functionConnectionId")
	delete(additionalProperties, "graphData")
	delete(additionalProperties, "inputSchema")
	delete(additionalProperties, "inputSchemaCompiled")
	delete(additionalProperties, "isInputSchemaSaved")
	delete(additionalProperties, "isOutputSchemaSaved")
	delete(additionalProperties, "name")
	delete(additionalProperties, "orx")
	delete(additionalProperties, "outputSchema")
	delete(additionalProperties, "outputSchemaCompiled")
	delete(additionalProperties, "parentFlowId")
	delete(additionalProperties, "publishedVersion")
	delete(additionalProperties, "savedDate")
	delete(additionalProperties, "settings")
	delete(additionalProperties, "timeouts")
	delete(additionalProperties, "trigger")
	delete(additionalProperties, "updatedDate")
	delete(additionalProperties, "variables")
	delete(additionalProperties, "versionId")
	delete(additionalProperties, "versionInfo")
	o.AdditionalProperties = additionalProperties

	return nil
}

func (o Flow) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o Flow) ToMap() (map[string]any, error) {

	result := map[string]any{}

	if o.Description != nil {
		result["authTokenExpireIds"] = o.AuthTokenExpireIds
	}

	result["companyId"] = o.CompanyID

	if o.Connections != nil {
		result["connections"] = o.Connections
	}

	result["connectorIds"] = o.ConnectorIds
	result["createdDate"] = o.CreatedDate

	if o.CurrentVersion != nil {
		result["currentVersion"] = o.CurrentVersion
	}

	result["customerId"] = o.CustomerID

	if o.DeployedDate != nil {
		result["deployedDate"] = o.DeployedDate
	}

	if o.Description != nil {
		result["description"] = o.Description
	}

	if o.EnabledGraphData != nil {
		result["enabledGraphData"] = o.EnabledGraphData
	}

	if o.FlowColor != nil {
		result["flowColor"] = o.FlowColor
	}

	result["flowId"] = o.FlowID

	if o.ParentFlowID != nil {
		result["parentFlowId"] = o.ParentFlowID
	}

	result["flowStatus"] = o.FlowStatus

	if o.FunctionConnectionID != nil {
		result["functionConnectionId"] = o.FunctionConnectionID
	}

	result["graphData"] = o.GraphData

	if o.InputSchema != nil {
		result["inputSchema"] = o.InputSchema
	}

	if o.InputSchemaCompiled != nil {
		result["inputSchemaCompiled"] = o.InputSchemaCompiled
	}

	if o.IsInputSchemaSaved != nil {
		result["isInputSchemaSaved"] = o.IsInputSchemaSaved
	}

	result["isOutputSchemaSaved"] = o.IsOutputSchemaSaved

	result["name"] = o.Name

	if o.Orx != nil {
		result["orx"] = o.Orx
	}

	if o.OutputSchema != nil {
		result["outputSchema"] = o.OutputSchema
	}

	if o.OutputSchemaCompiled != nil {
		result["outputSchemaCompiled"] = o.OutputSchemaCompiled
	}

	if o.PublishedVersion != nil {
		result["publishedVersion"] = o.PublishedVersion
	}

	result["savedDate"] = o.SavedDate

	if o.Settings != nil {
		result["settings"] = o.Settings
	}

	result["timeouts"] = o.Timeouts

	if o.Trigger != nil {
		result["trigger"] = o.Trigger
	}

	if o.UpdatedDate != nil {
		result["updatedDate"] = o.UpdatedDate
	}

	if o.Variables != nil {
		result["variables"] = o.Variables
	}

	result["versionId"] = o.VersionID

	if o.VersionInfo != nil {
		result["versionInfo"] = o.VersionInfo
	}

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *Flow) UnmarshalJSON(bytes []byte) (err error) {
	varFlow := _Flow{}

	if err = json.Unmarshal(bytes, &varFlow); err == nil {
		*o = Flow(varFlow)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "authTokenExpireIds")
		delete(additionalProperties, "companyId")
		delete(additionalProperties, "connections")
		delete(additionalProperties, "connectorIds")
		delete(additionalProperties, "createdDate")
		delete(additionalProperties, "currentVersion")
		delete(additionalProperties, "customerId")
		delete(additionalProperties, "deployedDate")
		delete(additionalProperties, "description")
		delete(additionalProperties, "enabledGraphData")
		delete(additionalProperties, "flowColor")
		delete(additionalProperties, "flowId")
		delete(additionalProperties, "flowStatus")
		delete(additionalProperties, "functionConnectionId")
		delete(additionalProperties, "graphData")
		delete(additionalProperties, "inputSchema")
		delete(additionalProperties, "inputSchemaCompiled")
		delete(additionalProperties, "isInputSchemaSaved")
		delete(additionalProperties, "isOutputSchemaSaved")
		delete(additionalProperties, "name")
		delete(additionalProperties, "orx")
		delete(additionalProperties, "outputSchema")
		delete(additionalProperties, "outputSchemaCompiled")
		delete(additionalProperties, "parentFlowId")
		delete(additionalProperties, "publishedVersion")
		delete(additionalProperties, "savedDate")
		delete(additionalProperties, "settings")
		delete(additionalProperties, "timeouts")
		delete(additionalProperties, "trigger")
		delete(additionalProperties, "updatedDate")
		delete(additionalProperties, "variables")
		delete(additionalProperties, "versionId")
		delete(additionalProperties, "versionInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
