package davinci_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/samir-gandhi/davinci-client-go/davinci"
)

func TestTrigger_Configuration_RoundTrip(t *testing.T) {
	enabled := true
	triggerType := "AUTHENTICATION"

	testCases := map[string]func(t *testing.T){
		"trigger-type-only": func(t *testing.T) {
			trigger := davinci.Trigger{
				TriggerType: &triggerType,
			}

			b, err := json.Marshal(trigger)
			if err != nil {
				t.Fatalf("json.Marshal failed: %v", err)
			}
			jsonStr := string(b)

			if strings.Contains(jsonStr, `"configuration"`) {
				t.Errorf("marshaled JSON unexpectedly contains \"configuration\" key: %s", jsonStr)
			}
			if !strings.Contains(jsonStr, `"type"`) {
				t.Errorf("marshaled JSON missing \"type\" key: %s", jsonStr)
			}

			var got davinci.Trigger
			if err := json.Unmarshal(b, &got); err != nil {
				t.Fatalf("json.Unmarshal failed: %v", err)
			}

			if got.TriggerType == nil || *got.TriggerType != triggerType {
				t.Errorf("expected TriggerType=%q, got %v", triggerType, got.TriggerType)
			}
			if _, ok := got.AdditionalProperties["configuration"]; ok {
				t.Errorf("\"configuration\" key present in AdditionalProperties: %v", got.AdditionalProperties)
			}
		},
		"trigger-with-configuration-lands-in-additional-properties": func(t *testing.T) {
			raw := `{"type":"AUTHENTICATION","configuration":{"mfa":{"enabled":true}}}`

			var got davinci.Trigger
			if err := json.Unmarshal([]byte(raw), &got); err != nil {
				t.Fatalf("json.Unmarshal failed: %v", err)
			}

			if got.TriggerType == nil || *got.TriggerType != triggerType {
				t.Errorf("expected TriggerType=%q, got %v", triggerType, got.TriggerType)
			}
			if _, ok := got.AdditionalProperties["configuration"]; !ok {
				t.Errorf("expected \"configuration\" in AdditionalProperties, but it was absent: %v", got.AdditionalProperties)
			}
		},
		"policy-trigger-mfa-enabled": func(t *testing.T) {
			pt := davinci.PolicyTrigger{
				Type: &triggerType,
				Configuration: &davinci.TriggerConfiguration{
					MFA: &davinci.TriggerConfigurationMFA{
						Enabled: &enabled,
					},
				},
			}

			b, err := json.Marshal(pt)
			if err != nil {
				t.Fatalf("json.Marshal failed: %v", err)
			}
			jsonStr := string(b)

			if !strings.Contains(jsonStr, `"configuration"`) {
				t.Errorf("marshaled JSON missing \"configuration\" key: %s", jsonStr)
			}
			if !strings.Contains(jsonStr, `"mfa"`) {
				t.Errorf("marshaled JSON missing \"mfa\" key: %s", jsonStr)
			}
			if !strings.Contains(jsonStr, `"enabled":true`) {
				t.Errorf("marshaled JSON missing \"enabled\":true under mfa: %s", jsonStr)
			}

			var got davinci.PolicyTrigger
			if err := json.Unmarshal(b, &got); err != nil {
				t.Fatalf("json.Unmarshal failed: %v", err)
			}

			if got.Configuration == nil {
				t.Fatal("unmarshaled PolicyTrigger.Configuration is nil")
			}
			if got.Configuration.MFA == nil {
				t.Fatal("unmarshaled PolicyTrigger.Configuration.MFA is nil")
			}
			if got.Configuration.MFA.Enabled == nil {
				t.Fatal("unmarshaled PolicyTrigger.Configuration.MFA.Enabled is nil")
			}
			if *got.Configuration.MFA.Enabled != true {
				t.Errorf("expected PolicyTrigger.Configuration.MFA.Enabled = true, got %v", *got.Configuration.MFA.Enabled)
			}
		},
	}

	for name, run := range testCases {
		t.Run(name, func(t *testing.T) {
			run(t)
		})
	}
}
