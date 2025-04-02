package davinci_test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/samir-gandhi/davinci-client-go/davinci"
	"github.com/samir-gandhi/davinci-client-go/davinci/test"
	"github.com/samir-gandhi/davinci-client-go/davinci/test/data"
)

/* TEST PLAN:
* - Read all
* - Create new
* - Read new
* - Update new
* - Delete new
 */

// testData for Roles functions
var testDataFlows = map[string]any{
	"params": map[string]davinci.Params{
		"limitTen": {Limit: "10"},
		"limitTwo": {Limit: "2"},
		"limitNil": {},
		// Flows doesn't allow page arg
		"pageNeg": {Page: "2"},
	},
	// "flowsCreate": Flow{
	// 	AdditionalProperties: map[string]any{
	// 		"custom-unimplemented-1": "custom-unimplemented-1",
	// 		"custom-unimplemented-2": "custom-unimplemented-2",
	// 	},
	// 	AuthTokenExpireIds: []interface{}{},
	// 	CompanyID:          "2c6123ae-108f-4d11-bcc2-6c8f4dfa9fdb",
	// 	Connections:        []interface{}{},
	// 	CreatedDate:        time.Unix(1706708769850, 0),
	// 	CurrentVersion:     8,
	// 	CustomerID:         "db5f4450b2bd8a56ce076dec0c358a9a",
	// 	DeployedDate:       time.Unix(1706709739837, 0),
	// 	FlowID:             "c7062a8857740ee2185694bb855f8f21",
	// 	PublishedVersion:   8,
	// 	SavedDate:          1706708769645,
	// 	UpdatedDate:        1706709739837,
	// 	VersionID:          8,
	// 	FlowConfigProperties: FlowConfig{
	// 		ConnectorIds: []string{
	// 			"httpConnector",
	// 			"functionsConnector",
	// 			"errorConnector",
	// 			"flowConnector",
	// 			"variablesConnector",
	// 		},
	// 		Description:          "",
	// 		EnabledGraphData:     nil,
	// 		FlowColor:            "#AFD5FF",
	// 		FunctionConnectionID: nil,
	// 		GraphData:            GraphData{},
	// 	},
	// },
	"flowsCreateJson": map[string]any{
		"properImport": func() string {
			v, err := test.ReadJsonFile("flows/flows-create-1.json")
			if err != nil {
				panic(fmt.Errorf("Failed to read file: %v", err))
			}
			return v
		}(),
		"directExport": func() string {
			v, err := test.ReadJsonFile("flows/flows-create-2.json")
			if err != nil {
				panic(fmt.Errorf("Failed to read file: %v", err))
			}
			return v
		}(),
		"directExportNeg": func() string {
			v, err := test.ReadJsonFile("flows/flows-create-3.json")
			if err != nil {
				panic(fmt.Errorf("Failed to read file: %v", err))
			}
			return v
		}(),
		"arrayEscapedImport": data.FLOW_ESCAPED,
		"arrayImport": func() string {
			v, err := test.ReadJsonFile("flows/flows-create-4.json")
			if err != nil {
				panic(fmt.Errorf("Failed to read file: %v", err))
			}
			return v
		}(),
	},
}

// Gets an array of all roles for a company
func TestFlows_Read(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")

	if args, ok := testDataFlows["flowsCreateJson"].(map[string]any); ok {
		for i, thisArg := range args {
			testName := i
			t.Run(testName, func(t *testing.T) {
				//fmt.Printf("thisArg is %q\n", thisArg)

				if payloadJson, ok := thisArg.(string); ok {
					resp, err := c.CreateFlow(companyID, &payloadJson)
					if err != nil {
						fmt.Println(err.Error())
						log.Printf("Failed Successfully\n")
						// if it's not a negative test, consider it an actual failure.
						if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
							log.Printf("Failed. Error is: %v", err)
							t.Fatal()
						}
					}
					if resp != nil {
						log.Printf("Flows Created Successfully\n resp.FlowId is: %v \n", resp.FlowID)
					}

				}
			})
		}
	}

	args, _ := testDataFlows["params"].(map[string]davinci.Params)
	for i := range args {
		testName := i
		t.Run(testName, func(t *testing.T) {
			thisArgs := args[i]
			fmt.Printf("Test Args are: %q\n", thisArgs)

			resp, err := c.ReadFlows(companyID, &thisArgs)
			if err != nil {
				fmt.Println(err.Error())
				log.Printf("Failed Successfully\n")
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					log.Printf("failed to get flows with params: %v \n Error is: %v", args, err)
					t.Fatal()
				}
			}
			if resp != nil {
				log.Printf("Flows Returned Successfully\n resp[0].FlowId is: %+v \n", resp[0])
			}
			// Too verbose to print all Flows.

		})
	}
}

func TestFlow_Create(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	if args, ok := testDataFlows["flowsCreateJson"].(map[string]any); ok {
		for i, thisArg := range args {
			testName := i
			t.Run(testName, func(t *testing.T) {
				//fmt.Printf("thisArg is %q\n", thisArg)

				if payloadJson, ok := thisArg.(string); ok {
					resp, err := c.CreateFlow(companyID, &payloadJson)
					if err != nil {
						fmt.Println(err.Error())
						log.Printf("Failed Successfully\n")
						// if it's not a negative test, consider it an actual failure.
						if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
							log.Printf("Failed. Error is: %v", err)
							t.Fatal()
						}
					}
					if resp != nil {
						log.Printf("Flows Created Successfully\n resp.FlowId is: %v \n", resp.FlowID)
					}

				}
			})
		}
	}
}

func TestFlow_CreateWithJson(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	if args, ok := testDataFlows["flowsCreateJson"].(map[string]any); ok {
		for i, thisArg := range args {
			testName := i
			t.Run(testName, func(t *testing.T) {
				fmt.Printf("thisArg is %q\n", thisArg)

				if payloadJson, ok := thisArg.(string); ok {
					resp, err := c.CreateFlowWithJson(companyID, &payloadJson)
					if err != nil {
						fmt.Println(err.Error())
						log.Printf("Failed Successfully\n")
						// if it's not a negative test, consider it an actual failure.
						if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
							log.Printf("Failed with params: %v \n Error is: %v", args, err)
							t.Fatal()
						}
					}
					if resp != nil {
						log.Printf("Flows Created Successfully\n resp.FlowId is: %v \n", resp.FlowID)
					}

				}
			})
		}
	}
}

func TestFlow_Read(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	flows, err := c.ReadFlows(companyID, &davinci.Params{Limit: "3"})
	if err != nil {
		t.Fatal()
	}
	for _, testVal := range flows {
		testName := testVal.Name
		t.Run(testName, func(t *testing.T) {

			resp, err := c.ReadFlowVersion(companyID, testVal.FlowID, nil)
			if err != nil {
				fmt.Println(err.Error())
				t.Fatal()
				log.Printf("Failed with error: %v \n", err.Error())
			}
			if resp.Flow.Name == "" {
				t.Fatal()
				log.Printf("Failed to get flow with name: %v \n Returned empty flow.", testName)
			}
			if resp.Flow.Name != "" {
				log.Printf("Flow Returned Successfully\n resp.FlowName is: %+v \n", resp.Flow.Name)
			}

		})
	}
}

//TODO - Use data specific to update for this.

func TestFlow_Update(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	if args, ok := testDataFlows["flowsCreateJson"].(map[string]any); ok {
		for i, thisArg := range args {
			testName := i
			t.Run(testName, func(t *testing.T) {
				// fmt.Printf("thisArg is %q\n", thisArg)

				if payloadJson, ok := thisArg.(string); ok {
					resp, err := c.CreateFlow(companyID, &payloadJson)
					if err != nil {
						fmt.Println(err.Error())
						log.Printf("Failed Successfully\n")
						// if it's not a negative test, consider it an actual failure.
						if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
							log.Printf("Failed with params error is: %v", err)
							t.Fatal()
						}
					}
					if resp != nil {
						fmt.Printf("Flows Created Successfully\n resp.FlowId is: %v \n", resp.FlowID)
						resp.Name = fmt.Sprintf("%v-UPDATED", resp.Name)
						payload, err := json.Marshal(resp)
						if err != nil {
							t.Fatal()
							return
						}

						var data davinci.FlowUpdate
						err = davinci.Unmarshal([]byte(payload[:]), &data, davinci.ExportCmpOpts{})
						if err != nil {
							t.Fatalf(err.Error())
							return
						}

						resp, err := c.UpdateFlow(companyID, resp.FlowID, data)
						if err != nil {
							log.Printf("Update failed")
							t.Fatal()
						}
						fmt.Printf("Update resp.Name: %v", resp.Name)
					}

				}
			})
		}
	}
}

func TestFlow_Delete(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	if args, ok := testDataFlows["flowsCreateJson"].(map[string]any); ok {
		for i, thisArg := range args {
			testName := i
			t.Run(testName, func(t *testing.T) {
				fmt.Printf("thisArg is %q\n", thisArg)

				if payloadJson, ok := thisArg.(string); ok {
					resp, err := c.CreateFlowWithJson(companyID, &payloadJson)
					if err != nil {
						fmt.Println(err.Error())
						log.Printf("Failed Successfully\n")

						// if it's not a negative test, consider it an actual failure.
						if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
							log.Printf("Failed with params: %v \n Error is: %v", args, err)

							t.Fatal()
							return
						}
					}
					if resp != nil {
						fmt.Printf("Flows Created Successfully\n resp.FlowId is: %v \n", resp.FlowID)
						resp, err := c.DeleteFlow(companyID, resp.FlowID)
						if err != nil {
							log.Printf("Delete failed")

							t.Fatal()
						}
						fmt.Printf("Deleted Successfully: %v", resp)
					}

				}
			})
		}
	}
}

func TestFlow_Deploy(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	if args, ok := testDataFlows["flowsCreateJson"].(map[string]any); ok {
		for i, thisArg := range args {
			testName := i
			t.Run(testName, func(t *testing.T) {
				fmt.Printf("thisArg is %q\n", thisArg)

				if payloadJson, ok := thisArg.(string); ok {
					resp, err := c.CreateFlowWithJson(companyID, &payloadJson)
					if err != nil {
						fmt.Println(err.Error())
						log.Printf("Failed Successfully\n")

						// if it's not a negative test, consider it an actual failure.
						if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
							log.Printf("Failed with params: %v \n Error is: %v", args, err)

							t.Fatal()
							return
						}
					}
					if resp != nil {
						fmt.Printf("Flows Created Successfully\n resp.FlowId is: %v \n", resp.FlowID)
						resp, err := c.DeployFlow(companyID, resp.FlowID)
						if err != nil {
							log.Printf("Delete failed")

							t.Fatal()
						}
						fmt.Printf("Deleted Successfully: %v", resp)
					}

				}
			})
		}
	}
}
