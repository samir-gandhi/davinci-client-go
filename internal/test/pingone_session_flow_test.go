package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingOneSessionFlowApp(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	apps := makeTestApps(RandString(10))
	app := apps["plain"]
	policy := testDataAppPolicy["basePolicy"]
	app.Policies = append(app.Policies, policy)
	// using flow that is manually created in feature flag env.
	app.Policies[0].PolicyFlows[0].FlowID = "5c32a89d4093b0eba7292ecafdd6b0e9"
	resp, err := c.CreateInitializedApplication(companyID, &app)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, "AUTHENTICATION", resp.Policies[0].Trigger.Type)
}

func TestNoPolicyApp(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	apps := makeTestApps(RandString(10))
	app := apps["noPolicy"]
	_, err = c.CreateInitializedApplication(companyID, &app)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
