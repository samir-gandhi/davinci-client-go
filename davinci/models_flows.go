package davinci

type FlowInfo struct {
	Flow Flow `json:"flowInfo"`
}
type FlowsInfo struct {
	Flow []Flow `json:"flowsInfo,omitempty"`
}

type FlowImport struct {
	Name            string            `json:"name,omitempty"`
	Description     string            `json:"description,omitempty"`
	FlowInfo        Flow              `json:"flowInfo,omitempty"`
	FlowNameMapping map[string]string `json:"flowNameMapping,omitempty"`
}
type FlowsImport struct {
	Name            string            `json:"name,omitempty"`
	Description     string            `json:"description,omitempty"`
	FlowInfo        Flows             `json:"flowInfo,omitempty"`
	FlowNameMapping map[string]string `json:"flowNameMapping,omitempty"`
}

type Flows struct {
	Flow []Flow `json:"flows,omitempty"`
}

// Used specifically for PUTs to existing flows.
type FlowUpdate struct {
	CurrentVersion int           `json:"currentVersion,omitempty"`
	Name           string        `json:"name,omitempty"`
	Description    string        `json:"description,omitempty"`
	Settings       interface{}   `json:"settings,omitempty"`
	Trigger        Trigger       `json:"trigger,omitempty"`
	GraphData      GraphData     `json:"graphData,omitempty"`
	InputSchema    []interface{} `json:"inputSchema,omitempty"`
	OutputSchema   OutputSchema  `json:"outputSchema,omitempty"`
}

type OutputSchema struct {
	Output interface{} `json:"output,omitempty"`
}

//	type ShowContinueButton struct {
//		Value bool `json:"value,omitempty"`
//	}
type SubFlowValue struct {
	Label string `json:"label,omitempty" mapstructure:"label"`
	Value string `json:"value,omitempty" mapstructure:"value"`
}
type SubFlowID struct {
	Value SubFlowValue `json:"value,omitempty" mapstructure:"value"`
}
type SubFlowVersionID struct {
	Value string `json:"value,omitempty" mapstructure:"value"`
}

// Used for type assertion on Properties of a Node Data
type SubFlowProperties struct {
	SubFlowID        SubFlowID        `json:"subFlowId,omitempty" mapstructure:"subFlowId"`
	SubFlowVersionID SubFlowVersionID `json:"subFlowVersionId,omitempty" mapstructure:"subFlowVersionId"`
}
