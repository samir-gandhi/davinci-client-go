package davinci

import (
	"encoding/json"
	"maps"
)

type _FlowSettings FlowSettings
type FlowSettings struct {
	AdditionalProperties            map[string]any           `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Csp                             *FlowSettingsStringValue `json:"csp,omitempty" davinci:"csp,config,omitempty"`
	Css                             *FlowSettingsStringValue `json:"css,omitempty" davinci:"css,config,omitempty"`
	CssLinks                        *[]string                `json:"cssLinks" davinci:"cssLinks,config"`
	CustomErrorScreenBrandLogoUrl   *string                  `json:"customErrorScreenBrandLogoUrl,omitempty" davinci:"customErrorScreenBrandLogoUrl,config,omitempty"`
	CustomErrorShowFooter           *bool                    `json:"customErrorShowFooter,omitempty" davinci:"customErrorShowFooter,config,omitempty"`
	CustomFaviconLink               *string                  `json:"customFaviconLink,omitempty" davinci:"customFaviconLink,config,omitempty"`
	CustomLogoURLSelection          *interface{}             `json:"customLogoURLSelection,omitempty" davinci:"customLogoURLSelection,config,omitempty"`
	CustomTitle                     *string                  `json:"customTitle,omitempty" davinci:"customTitle,config,omitempty"`
	DebugMode                       *bool                    `json:"debugMode,omitempty" davinci:"debugMode,environmentmetadata,omitempty"`
	DefaultErrorScreenBrandLogo     *bool                    `json:"defaultErrorScreenBrandLogo,omitempty" davinci:"defaultErrorScreenBrandLogo,config,omitempty"`
	DisplayNodeIDs                  *bool                    `json:"displayNodeIDs,omitempty" davinci:"displayNodeIDs,designercue,omitempty"`
	DoNotSubstituteUnreplacedFields *bool                    `json:"doNotSubstituteUnreplacedFields,omitempty" davinci:"doNotSubstituteUnreplacedFields,config,omitempty"`
	FlowHttpTimeoutInSeconds        *int32                   `json:"flowHttpTimeoutInSeconds,omitempty" davinci:"flowHttpTimeoutInSeconds,config,omitempty"`
	FlowTimeoutInSeconds            *int32                   `json:"flowTimeoutInSeconds,omitempty" davinci:"flowTimeoutInSeconds,config,omitempty"`
	IntermediateLoadingScreenCSS    *FlowSettingsStringValue `json:"intermediateLoadingScreenCSS,omitempty" davinci:"intermediateLoadingScreenCSS,config,omitempty"`
	IntermediateLoadingScreenHTML   *FlowSettingsStringValue `json:"intermediateLoadingScreenHTML,omitempty" davinci:"intermediateLoadingScreenHTML,config,omitempty"`
	JsLinks                         *[]interface{}           `json:"jsLinks" davinci:"jsLinks,config"`
	LogLevel                        *int32                   `json:"logLevel,omitempty" davinci:"logLevel,environmentmetadata,omitempty"`
	PingOneFlow                     *bool                    `json:"pingOneFlow,omitempty" davinci:"pingOneFlow,config,omitempty"`
	RequireAuthenticationToInitiate *bool                    `json:"requireAuthenticationToInitiate,omitempty" davinci:"requireAuthenticationToInitiate,config,omitempty"`
	ScrubSensitiveInfo              *bool                    `json:"scrubSensitiveInfo,omitempty" davinci:"scrubSensitiveInfo,config,omitempty"`
	SensitiveInfoFields             *interface{}             `json:"sensitiveInfoFields,omitempty" davinci:"sensitiveInfoFields,config,omitempty"`
	UseBetaAlgorithm                *bool                    `json:"useBetaAlgorithm,omitempty" davinci:"useBetaAlgorithm,config,omitempty"`
	UseCsp                          *bool                    `json:"useCsp,omitempty" davinci:"useCsp,config,omitempty"`
	UseCustomCSS                    *bool                    `json:"useCustomCSS,omitempty" davinci:"useCustomCSS,config,omitempty"`
	UseCustomScript                 *bool                    `json:"useCustomScript,omitempty" davinci:"useCustomScript,config,omitempty"`
	UseIntermediateLoadingScreen    *bool                    `json:"useIntermediateLoadingScreen,omitempty" davinci:"useIntermediateLoadingScreen,config,omitempty"`
}

func (o FlowSettings) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o FlowSettings) ToMap() (map[string]any, error) {

	result := map[string]any{}

	if o.Csp != nil {
		result["csp"] = o.Csp
	}

	if o.Css != nil {
		result["css"] = o.Css
	}

	if o.CssLinks != nil {
		if len(*o.CssLinks) == 0 {
			result["cssLinks"] = []string{}
		} else {
			result["cssLinks"] = o.CssLinks
		}
	}

	if o.CustomErrorScreenBrandLogoUrl != nil {
		result["customErrorScreenBrandLogoUrl"] = o.CustomErrorScreenBrandLogoUrl
	}

	if o.CustomErrorShowFooter != nil {
		result["customErrorShowFooter"] = o.CustomErrorShowFooter
	}

	if o.CustomFaviconLink != nil {
		result["customFaviconLink"] = o.CustomFaviconLink
	}

	if o.CustomLogoURLSelection != nil {
		result["customLogoURLSelection"] = o.CustomLogoURLSelection
	}

	if o.CustomTitle != nil {
		result["customTitle"] = o.CustomTitle
	}

	if o.DebugMode != nil {
		result["debugMode"] = o.DebugMode
	}

	if o.DefaultErrorScreenBrandLogo != nil {
		result["defaultErrorScreenBrandLogo"] = o.DefaultErrorScreenBrandLogo
	}

	if o.DisplayNodeIDs != nil {
		result["displayNodeIDs"] = o.DisplayNodeIDs
	}

	if o.DoNotSubstituteUnreplacedFields != nil {
		result["doNotSubstituteUnreplacedFields"] = o.DoNotSubstituteUnreplacedFields
	}

	if o.FlowHttpTimeoutInSeconds != nil {
		result["flowHttpTimeoutInSeconds"] = o.FlowHttpTimeoutInSeconds
	}

	if o.FlowTimeoutInSeconds != nil {
		result["flowTimeoutInSeconds"] = o.FlowTimeoutInSeconds
	}

	if o.IntermediateLoadingScreenCSS != nil {
		result["intermediateLoadingScreenCSS"] = o.IntermediateLoadingScreenCSS
	}

	if o.IntermediateLoadingScreenHTML != nil {
		result["intermediateLoadingScreenHTML"] = o.IntermediateLoadingScreenHTML
	}

	if o.JsLinks != nil {
		if len(*o.JsLinks) == 0 {
			result["jsLinks"] = []string{}
		} else {
			result["jsLinks"] = o.JsLinks
		}
	}

	if o.LogLevel != nil {
		result["logLevel"] = o.LogLevel
	}

	if o.PingOneFlow != nil {
		result["pingOneFlow"] = o.PingOneFlow
	}

	if o.RequireAuthenticationToInitiate != nil {
		result["requireAuthenticationToInitiate"] = o.RequireAuthenticationToInitiate
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

	if o.UseCsp != nil {
		result["useCsp"] = o.UseCsp
	}

	if o.UseCustomCSS != nil {
		result["useCustomCSS"] = o.UseCustomCSS
	}

	if o.UseCustomScript != nil {
		result["useCustomScript"] = o.UseCustomScript
	}

	if o.UseIntermediateLoadingScreen != nil {
		result["useIntermediateLoadingScreen"] = o.UseIntermediateLoadingScreen
	}

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *FlowSettings) UnmarshalJSON(bytes []byte) (err error) {
	varFlowSettings := _FlowSettings{}

	if err = json.Unmarshal(bytes, &varFlowSettings); err == nil {
		*o = FlowSettings(varFlowSettings)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "csp")
		delete(additionalProperties, "css")
		delete(additionalProperties, "cssLinks")
		delete(additionalProperties, "customErrorScreenBrandLogoUrl")
		delete(additionalProperties, "customErrorShowFooter")
		delete(additionalProperties, "customFaviconLink")
		delete(additionalProperties, "customLogoURLSelection")
		delete(additionalProperties, "customTitle")
		delete(additionalProperties, "debugMode")
		delete(additionalProperties, "defaultErrorScreenBrandLogo")
		delete(additionalProperties, "displayNodeIDs")
		delete(additionalProperties, "doNotSubstituteUnreplacedFields")
		delete(additionalProperties, "flowHttpTimeoutInSeconds")
		delete(additionalProperties, "flowTimeoutInSeconds")
		delete(additionalProperties, "intermediateLoadingScreenCSS")
		delete(additionalProperties, "intermediateLoadingScreenHTML")
		delete(additionalProperties, "jsLinks")
		delete(additionalProperties, "logLevel")
		delete(additionalProperties, "pingOneFlow")
		delete(additionalProperties, "requireAuthenticationToInitiate")
		delete(additionalProperties, "scrubSensitiveInfo")
		delete(additionalProperties, "sensitiveInfoFields")
		delete(additionalProperties, "useBetaAlgorithm")
		delete(additionalProperties, "useCsp")
		delete(additionalProperties, "useCustomCSS")
		delete(additionalProperties, "useCustomScript")
		delete(additionalProperties, "useIntermediateLoadingScreen")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
