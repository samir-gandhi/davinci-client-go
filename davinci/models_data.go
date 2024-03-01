package davinci

type Data struct {
	AdditionalProperties map[string]interface{} `davinci:"-,-"` // used to capture all other properties that are not explicitly defined in the model
	CapabilityName       *string                `davinci:"capabilityName,config,omitempty"`
	ConnectionID         *string                `davinci:"connectionId,config,omitempty"`
	ConnectorID          *string                `davinci:"connectorId,config,omitempty"`
	ID                   *string                `davinci:"id,config,omitempty"`
	Label                *string                `davinci:"label,config,omitempty"`
	MultiValueSourceId   *string                `davinci:"multiValueSourceId,config,omitempty"`
	Name                 *string                `davinci:"name,config,omitempty"`
	NodeType             *string                `davinci:"nodeType,config,omitempty"`
	Properties           *Properties            `davinci:"properties,*,omitempty"`
	Source               *string                `davinci:"source,config,omitempty"`
	Status               *string                `davinci:"status,config,omitempty"`
	Target               *string                `davinci:"target,config,omitempty"`
	Type                 *string                `davinci:"type,config,omitempty"`
}
