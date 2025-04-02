package davinci

import "encoding/json"

type _FlowSettings FlowSettings
type FlowSettings struct {
	AdditionalProperties          map[string]interface{} `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Csp                           *string                `json:"csp,omitempty" davinci:"csp,config,omitempty"`
	Css                           *string                `json:"css,omitempty" davinci:"css,config,omitempty"`
	CssLinks                      *[]string              `json:"cssLinks,omitempty" davinci:"cssLinks,config,omitempty"`
	DebugMode                     *bool                  `json:"debugMode,omitempty" davinci:"debugMode,environmentmetadata,omitempty"`
	FlowHttpTimeoutInSeconds      *int32                 `json:"flowHttpTimeoutInSeconds,omitempty" davinci:"flowHttpTimeoutInSeconds,config,omitempty"`
	IntermediateLoadingScreenCSS  *string                `json:"intermediateLoadingScreenCSS,omitempty" davinci:"intermediateLoadingScreenCSS,config,omitempty"`
	IntermediateLoadingScreenHTML *string                `json:"intermediateLoadingScreenHTML,omitempty" davinci:"intermediateLoadingScreenHTML,config,omitempty"`
	LogLevel                      *int32                 `json:"logLevel,omitempty" davinci:"logLevel,environmentmetadata,omitempty"`
	PingOneFlow                   *bool                  `json:"pingOneFlow,omitempty" davinci:"pingOneFlow,config,omitempty"`
	ScrubSensitiveInfo            *bool                  `json:"scrubSensitiveInfo,omitempty" davinci:"scrubSensitiveInfo,config,omitempty"`
	SensitiveInfoFields           *interface{}           `json:"sensitiveInfoFields,omitempty" davinci:"sensitiveInfoFields,config,omitempty"`
	UseBetaAlgorithm              *bool                  `json:"useBetaAlgorithm,omitempty" davinci:"useBetaAlgorithm,config,omitempty"`
	UseCustomCSS                  *bool                  `json:"useCustomCSS,omitempty" davinci:"useCustomCSS,config,omitempty"`
}

func (o FlowSettings) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o FlowSettings) ToMap() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	if o.Csp != nil {
		result["csp"] = o.Csp
	}

	if o.Css != nil {
		result["css"] = o.Css
	}

	if o.CssLinks != nil {
		result["cssLinks"] = o.CssLinks
	}

	if o.DebugMode != nil {
		result["debugMode"] = o.DebugMode
	}

	if o.FlowHttpTimeoutInSeconds != nil {
		result["flowHttpTimeoutInSeconds"] = o.FlowHttpTimeoutInSeconds
	}

	if o.IntermediateLoadingScreenCSS != nil {
		result["intermediateLoadingScreenCSS"] = o.IntermediateLoadingScreenCSS
	}

	if o.IntermediateLoadingScreenHTML != nil {
		result["intermediateLoadingScreenHTML"] = o.IntermediateLoadingScreenHTML
	}

	if o.LogLevel != nil {
		result["logLevel"] = o.LogLevel
	}

	if o.PingOneFlow != nil {
		result["pingOneFlow"] = o.PingOneFlow
	}

	if o.ScrubSensitiveInfo != nil {
		result["scrubSensitiveInfo"] = o.ScrubSensitiveInfo
	}

	if o.SensitiveInfoFields != nil {
		result["sensitiveInfoFields"] = o.SensitiveInfoFields
	}

	if o.UseBetaAlgorithm != nil {
		result["useBetaAlgorithm"] = o.UseBetaAlgorithm
	}

	if o.UseCustomCSS != nil {
		result["useCustomCSS"] = o.UseCustomCSS
	}

	for k, v := range o.AdditionalProperties {
		result[k] = v
	}

	return result, nil
}

func (o *FlowSettings) UnmarshalJSON(bytes []byte) (err error) {
	varFlowSettings := _FlowSettings{}

	if err = json.Unmarshal(bytes, &varFlowSettings); err == nil {
		*o = FlowSettings(varFlowSettings)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "csp")
		delete(additionalProperties, "css")
		delete(additionalProperties, "cssLinks")
		delete(additionalProperties, "debugMode")
		delete(additionalProperties, "flowHttpTimeoutInSeconds")
		delete(additionalProperties, "intermediateLoadingScreenCSS")
		delete(additionalProperties, "intermediateLoadingScreenHTML")
		delete(additionalProperties, "logLevel")
		delete(additionalProperties, "pingOneFlow")
		delete(additionalProperties, "scrubSensitiveInfo")
		delete(additionalProperties, "sensitiveInfoFields")
		delete(additionalProperties, "useBetaAlgorithm")
		delete(additionalProperties, "useCustomCSS")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
