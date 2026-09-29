package davinci_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/samir-gandhi/davinci-client-go/davinci"
)

// Timeout error screen settings attributes (pingone/terraform-provider-pingone#1360).
func TestFlowSettings_TimeoutErrorScreen_RoundTrip(t *testing.T) {
	message := "The session timed out. Please try again."
	html := "<div class=\"timeout-error\">The flow has timed out</div>"
	css := "body { background-color: #f0f0f0; }"

	testCases := map[string]func(t *testing.T){
		"marshal-and-unmarshal-lands-in-struct-fields": func(t *testing.T) {
			useCustom := true
			settings := davinci.FlowSettings{
				CustomTimeoutErrorScreenMessage: &davinci.FlowSettingsStringValue{ValueString: &message},
				CustomTimeoutErrorScreenHTML:    &davinci.FlowSettingsStringValue{ValueString: &html},
				CustomTimeoutErrorScreenCSS:     &davinci.FlowSettingsStringValue{ValueString: &css},
				UseCustomTimeoutErrorScreen:     &useCustom,
			}

			b, err := davinci.Marshal(settings, davinci.ExportCmpOpts{})
			if err != nil {
				t.Fatalf("davinci.Marshal failed: %v", err)
			}
			jsonStr := string(b)

			for _, key := range []string{
				"customTimeoutErrorScreenMessage",
				"customTimeoutErrorScreenHTML",
				"customTimeoutErrorScreenCSS",
				"useCustomTimeoutErrorScreen",
			} {
				if !strings.Contains(jsonStr, "\""+key+"\"") {
					t.Errorf("marshaled JSON missing %q key: %s", key, jsonStr)
				}
			}

			var got davinci.FlowSettings
			if err := davinci.Unmarshal(b, &got, davinci.ExportCmpOpts{}); err != nil {
				t.Fatalf("davinci.Unmarshal failed: %v", err)
			}

			if got.CustomTimeoutErrorScreenMessage == nil || got.CustomTimeoutErrorScreenMessage.ValueString == nil || *got.CustomTimeoutErrorScreenMessage.ValueString != message {
				t.Errorf("expected CustomTimeoutErrorScreenMessage=%q, got %+v", message, got.CustomTimeoutErrorScreenMessage)
			}
			if got.CustomTimeoutErrorScreenHTML == nil || got.CustomTimeoutErrorScreenHTML.ValueString == nil || *got.CustomTimeoutErrorScreenHTML.ValueString != html {
				t.Errorf("expected CustomTimeoutErrorScreenHTML=%q, got %+v", html, got.CustomTimeoutErrorScreenHTML)
			}
			if got.CustomTimeoutErrorScreenCSS == nil || got.CustomTimeoutErrorScreenCSS.ValueString == nil || *got.CustomTimeoutErrorScreenCSS.ValueString != css {
				t.Errorf("expected CustomTimeoutErrorScreenCSS=%q, got %+v", css, got.CustomTimeoutErrorScreenCSS)
			}
			if got.UseCustomTimeoutErrorScreen == nil || !*got.UseCustomTimeoutErrorScreen {
				t.Errorf("expected UseCustomTimeoutErrorScreen=true, got %v", got.UseCustomTimeoutErrorScreen)
			}
			for _, key := range []string{
				"customTimeoutErrorScreenMessage",
				"customTimeoutErrorScreenHTML",
				"customTimeoutErrorScreenCSS",
				"useCustomTimeoutErrorScreen",
			} {
				if _, ok := got.AdditionalProperties[key]; ok {
					t.Errorf("expected %q to be classified as config, found in AdditionalProperties: %v", key, got.AdditionalProperties)
				}
			}
		},
		"empty-object-value-tolerated": func(t *testing.T) {
			// The DaVinci API may return `{}` for screen value fields (as observed for the analogous intermediate loading screen fields).
			raw := `{"customTimeoutErrorScreenHTML":{},"useCustomTimeoutErrorScreen":true}`

			var got davinci.FlowSettings
			if err := json.Unmarshal([]byte(raw), &got); err != nil {
				t.Fatalf("json.Unmarshal failed: %v", err)
			}

			if got.CustomTimeoutErrorScreenHTML == nil {
				t.Fatalf("expected CustomTimeoutErrorScreenHTML to be set, got nil: %+v", got)
			}
			if got.UseCustomTimeoutErrorScreen == nil || !*got.UseCustomTimeoutErrorScreen {
				t.Errorf("expected UseCustomTimeoutErrorScreen=true, got %v", got.UseCustomTimeoutErrorScreen)
			}
		},
		"flow-json-with-fields-passes-validation": func(t *testing.T) {
			flowJson := `{
				"name": "simple",
				"flowId": "8f93840f61b58b043a0a38439a1c6640",
				"companyId": "2c6123ae-108f-4d11-bcc2-6c8f4dfa9fdb",
				"customerId": "db5f4450b2bd8a56ce076dec0c358a9a",
				"flowStatus": "enabled",
				"isOutputSchemaSaved": false,
				"savedDate": 1707837216592,
				"versionId": 4,
				"createdDate": 1707837216607,
				"authTokenExpireIds": [],
				"connectorIds": ["errorConnector"],
				"settings": {
					"useCustomTimeoutErrorScreen": true,
					"customTimeoutErrorScreenMessage": "The session timed out. Please try again.",
					"customTimeoutErrorScreenHTML": "<div class=\"timeout-error\">The flow has timed out</div>",
					"customTimeoutErrorScreenCSS": "body { background-color: #f0f0f0; }",
					"logLevel": 2
				},
				"graphData": {"elements": {"nodes": [
					{"data": {"id": "n1", "nodeType": "CONNECTION", "connectionId": "53ab83a4a4ab919d9f2cb02d9e111ac8", "connectorId": "errorConnector", "name": "Error Message", "label": "Error Message", "status": "configured", "capabilityName": "customErrorMessage", "type": "action", "properties": {"errorMessage": {"value": "This is an error"}}}}
				]}}
			}`

			err := davinci.ValidFlowExport([]byte(flowJson), davinci.ExportCmpOpts{
				IgnoreConfig:              false,
				IgnoreDesignerCues:        false,
				IgnoreEnvironmentMetadata: true,
				IgnoreUnmappedProperties:  false,
				IgnoreVersionMetadata:     true,
				IgnoreFlowMetadata:        true,
				IgnoreFlowVariables:       true,
			})
			if err != nil {
				t.Errorf("ValidFlowExport rejected timeout error screen settings: %v", err)
			}
		},
		"equality-detects-config-drift": func(t *testing.T) {
			useTrue := true
			useFalse := false
			flowA := davinci.Flow{FlowConfiguration: davinci.FlowConfiguration{FlowUpdateConfiguration: davinci.FlowUpdateConfiguration{
				Settings: &davinci.FlowSettings{UseCustomTimeoutErrorScreen: &useTrue},
			}}}
			flowB := davinci.Flow{FlowConfiguration: davinci.FlowConfiguration{FlowUpdateConfiguration: davinci.FlowUpdateConfiguration{
				Settings: &davinci.FlowSettings{UseCustomTimeoutErrorScreen: &useFalse},
			}}}

			if !davinci.Equal(flowA, flowB, davinci.ExportCmpOpts{IgnoreConfig: true, IgnoreUnmappedProperties: false}) {
				t.Errorf("expected flows to be equal when config is ignored")
			}
			if davinci.Equal(flowA, flowB, davinci.ExportCmpOpts{IgnoreConfig: false, IgnoreUnmappedProperties: false}) {
				t.Errorf("expected flows to be unequal when config drift is present")
			}
		},
	}

	for name, tc := range testCases {
		t.Run(name, tc)
	}
}
