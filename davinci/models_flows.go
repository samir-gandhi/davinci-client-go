package davinci

type FlowInfo struct {
	Flow Flow `json:"flowInfo" davinci:"flowInfo,*"`
}
type FlowsInfo struct {
	Flow []Flow `json:"flowsInfo,omitempty" davinci:"flowsInfo,*"`
}

type FlowImport struct {
	Name            *string           `json:"name,omitempty" davinci:"name,flowmetadata,omitempty"`
	Description     *string           `json:"description,omitempty" davinci:"description,flowmetadata,omitempty"`
	FlowInfo        *Flow             `json:"flowInfo,omitempty" davinci:"flowInfo,*,omitempty"`
	FlowNameMapping map[string]string `json:"flowNameMapping,omitempty" davinci:"flowNameMapping,*,omitempty"`
}
type FlowsImport struct {
	Name            *string           `json:"name,omitempty"`
	Description     *string           `json:"description,omitempty"`
	FlowInfo        *Flows            `json:"flowInfo,omitempty"`
	FlowNameMapping map[string]string `json:"flowNameMapping,omitempty"`
}

type Flows struct {
	Flow []Flow `json:"flows,omitempty" davinci:"flows,*"`
}

// Used specifically for PUTs to existing flows.
type FlowUpdate struct {
	FlowUpdateConfiguration
	CurrentVersion *int32  `json:"currentVersion,omitempty" davinci:"currentVersion,versionmetadata,omitempty"`
	Name           *string `json:"name,omitempty" davinci:"name,flowmetadata"`
	Description    *string `json:"description,omitempty" davinci:"description,flowmetadata,omitempty"`
}

//	type ShowContinueButton struct {
//		Value *bool `json:"value,omitempty"`
//	}

type AdditionalProperties map[string]any
