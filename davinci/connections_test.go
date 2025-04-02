package davinci_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/samir-gandhi/davinci-client-go/davinci"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// type property struct {
// 	EnvId struct {
// 		Value string `json:"value`
// 	} `json:"envId`
// }

var testDataConnections = map[string]any{
	"params": map[string]davinci.Params{
		"a": {"1", "10", nil},
		"b": {"1000", "10", nil},
		"c": {},
	},
	"connectionsCreate": map[string]davinci.Connection{
		"aCreate": {
			Name: func() *string {
				s := "connectionACreate"
				return &s
			}(),
			ConnectorID: func() *string {
				s := "pingOneMfaConnector"
				return &s
			}(),
		},
		"bCreate": {
			Name: func() *string {
				s := "connectionBCreate"
				return &s
			}(),
			ConnectorID: func() *string {
				s := "pingOneMfaConnector"
				return &s
			}(),
			CustomerID: func() *string {
				s := "1234"
				return &s
			}(),
		},
		"cCreateneg": {
			Name: func() *string {
				s := ""
				return &s
			}(),
			ConnectorID: func() *string {
				s := "pingOneMfaConnector"
				return &s
			}(),
		},
		"dCreateneg": {
			Name: func() *string {
				s := "connectionDCreate"
				return &s
			}(),
			ConnectorID: func() *string {
				s := ""
				return &s
			}(),
		},
	},
	"connectionsRead": map[string]davinci.Connection{
		"aRead": {
			Name: func() *string {
				s := "connectionARead"
				return &s
			}(),
			ConnectorID: func() *string {
				s := "pingOneMfaConnector"
				return &s
			}(),
		},
		"bRead": {
			Name: func() *string {
				s := "connectionBRead"
				return &s
			}(),
			ConnectorID: func() *string {
				s := "pingOneMfaConnector"
				return &s
			}(),
			CustomerID: func() *string {
				s := "1234"
				return &s
			}(),
		},
		"cReadneg": {
			Name: func() *string {
				s := ""
				return &s
			}(),
			ConnectorID: func() *string {
				s := "pingOneMfaConnector"
				return &s
			}(),
		},
		"dReadneg": {
			Name: func() *string {
				s := "connectionDRead"
				return &s
			}(),
			ConnectorID: func() *string {
				s := ""
				return &s
			}(),
		},
		"eRead": {
			Name: func() *string {
				s := "connectionERead"
				return &s
			}(),
			ConnectorID: func() *string {
				s := "genericConnector"
				return &s
			}(),
		},
	},
	"connectionsUpdate": map[string]davinci.Connection{
		"aUpdate": {
			Name: func() *string {
				s := "connectionAUpdate"
				return &s
			}(),
			ConnectorID: func() *string {
				s := "pingOneMfaConnector"
				return &s
			}(),
			Properties: map[string]davinci.ConnectionProperty{
				"envId": {
					Value: "1234",
				},
				"policyId": {
					Value: "1234",
				},
			},
		},
		"bUpdate": {
			Name: func() *string {
				s := "connectionBUpdate"
				return &s
			}(),
			ConnectorID: func() *string {
				s := "pingOneSSOConnector"
				return &s
			}(),
			Properties: map[string]davinci.ConnectionProperty{
				"envId": {
					Value: "1234",
				},
			},
		},
		"cUpdate": {
			Name: func() *string {
				s := "connectionCUpdate"
				return &s
			}(),
			ConnectorID: func() *string {
				s := "pingOneMfaConnector"
				return &s
			}(),
			Properties: map[string]davinci.ConnectionProperty{},
		},
		"dUpdateNeg": {
			Name: func() *string {
				s := "connectionCUpdateNeg"
				return &s
			}(),
			ConnectorID: func() *string {
				s := "pingOneMfaConnector"
				return &s
			}(),
			Properties: map[string]davinci.ConnectionProperty{
				"InvalidProperty": {
					Value: "Foo",
				},
			},
		},
	},
	"connectionsCreateInitialized": map[string]davinci.Connection{
		// "aCreateInitialized": {
		// 	Name:        "connectionACreateInitialized",
		// 	ConnectorID: "pingOneMfaConnector",
		// 	Properties: Properties{
		// 		"envId": struct {
		// 			Value string `json:"value"`
		// 		}{"1234"},
		// 		"policyId": struct {
		// 			Value string `json:"value"`
		// 		}{"1234"},
		// 	},
		// },
		// "bCreateInitialized": {
		// 	Name:        "connectionBCreateInitialized",
		// 	ConnectorID: "pingOneSSOConnector",
		// 	Properties: Properties{
		// 		"envId": struct {
		// 			Value string `json:"value"`
		// 		}{"1234"},
		// 	},
		// },
		// "cCreateInitialized": {
		// 	Name:        "connectionCCreateInitialized",
		// 	ConnectorID: "pingOneMfaConnector",
		// 	Properties:  Properties{},
		// },
		// "dCreateInitializedNeg": {
		// 	Name:        "connectionDCreateInitialized",
		// 	ConnectorID: "pingOneMfaConnector",
		// 	Properties: Properties{
		// 		"InvalidProperty": struct {
		// 			Value string `json:"value"`
		// 		}{"Foo"},
		// 	},
		// },
		"eCreateInitializedOidc": {
			Name: func() *string {
				s := "connectionECreateInitializedOidc"
				return &s
			}(),
			ConnectorID: func() *string {
				s := "genericConnector"
				return &s
			}(),
			Properties: map[string]davinci.ConnectionProperty{
				"customAuth": {
					Properties: map[string]davinci.ConnectionProperty{
						"providerName": {
							Value: "foooidc",
						},
					},
				},
			},
		},
	},
}

func TestConnections_Read(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	args, _ := testDataConnections["params"].(map[string]davinci.Params)
	for i := range args {
		thisArgs := args[i]
		fmt.Printf("args[i] is %q\n", thisArgs)
		resp, err := c.ReadConnections(companyID, &thisArgs)
		if err != nil {
			fmt.Println(err.Error())
			if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
				t.Fatal()
			}
		}
		fmt.Printf("resp is: %v\n", resp)
	}
}

func TestConnection_Create(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	if args, ok := testDataConnections["connectionsCreate"].(map[string]davinci.Connection); ok {
		for i := range args {
			thisArg := args[i]
			name := fmt.Sprintf("%v-%v", thisArg.Name, RandString(6))
			thisArg.Name = &name
			fmt.Printf("args[i] is %#v\n", thisArg)
			resp, err := c.CreateConnection(companyID, &thisArg)
			if err != nil {
				fmt.Println(err.Error())
				// if it's not a negative test, consider it an actual failure.
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					t.Fatal()
				}
				fmt.Printf("Connection from key %v failed, continuing to next iteration \n", i)
				continue
			}
			thisArg.ConnectionID = resp.ConnectionID
			args[i] = thisArg
			fmt.Printf("resp is: %#v\n", resp)
		}
	}
}

//TODO - instead of creating connection for test, read all connections for connectionId

func TestConnection_Read(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	if args, ok := testDataConnections["connectionsRead"].(map[string]davinci.Connection); ok {
		for i := range args {
			thisArg := args[i]
			name := fmt.Sprintf("%v-%v", thisArg.Name, RandString(6))
			thisArg.Name = &name
			fmt.Printf("thisArg is %#v\n", thisArg)
			// Reading connection requires generated ConnectionID
			resp, err := c.CreateConnection(companyID, &thisArg)
			if err != nil {
				fmt.Println(err.Error())
				// if it's not a negative test, consider it an actual failure.
				if !(strings.Contains(i, "neg")) {
					t.Fatal()
				}
				// stop iteration of loop if connection create failed.
				fmt.Printf("Connection from key %v creation failed, continuing to next iteration \n", i)
				continue
			}
			thisArg.ConnectionID = resp.ConnectionID
			args[i] = thisArg
			fmt.Println(args[i])
			fmt.Printf("resp is: %#v\n", resp)
			res, err := c.ReadConnection(companyID, *resp.ConnectionID)
			if err != nil {
				fmt.Println(err.Error())
				// if it's not a negative test, consider it an actual failure.
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					t.Fatal()
				}
			}
			fmt.Printf("res is: %v\n", res)
			if v, ok := res.Properties["customAuth"]; ok {
				fmt.Println("customAuth is not nil, will attempt to unmarshal")
				a, _ := json.Marshal(v)
				var customAuth davinci.CustomAuth
				err := json.Unmarshal(a, &customAuth)
				if err != nil {
					fmt.Println(err.Error())
					t.Fatal()
				}
				if *customAuth.Properties.ClientID.Value == "" {
					fmt.Println("customAuth.Properties.ClientID.Value is empty after unmarshal")
					t.Fatal()
				}
			}
		}
	}
}

func TestConnection_Update(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	if args, ok := testDataConnections["connectionsUpdate"].(map[string]davinci.Connection); ok {
		for i := range args {
			thisArg := args[i]
			name := fmt.Sprintf("%v-%v", thisArg.Name, RandString(6))
			thisArg.Name = &name
			fmt.Printf("args[i] is %#v\n", thisArg)
			// Reading connection requires generated ConnectionID
			// For testing, the connection is created to get ConnectionID from response
			resp, err := c.CreateConnection(companyID, &thisArg)
			if err != nil {
				fmt.Println(err.Error())
				// if it's not a negative test, consider it an actual failure.
				if !(strings.Contains(i, "neg")) {
					t.Fatal()
				}
				// stop iteration of loop if connection create failed.
				fmt.Printf("Connection from key %v creation failed, continuing to next iteration \n", i)
				continue
			}
			thisArg.ConnectionID = resp.ConnectionID
			args[i] = thisArg
			fmt.Printf("resp is: %#v\n", resp)
			//upon success, pull properties from test data
			resp.Properties = thisArg.Properties
			res, err := c.UpdateConnection(companyID, resp)
			if err != nil {
				fmt.Println(err.Error())
				// if it's not a negative test, consider it an actual failure.
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					t.Fatal()
				}
			}
			fmt.Printf("res is: %v\n", res)
		}
	}
}

func TestConnection_CreateInitialized(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	if args, ok := testDataConnections["connectionsCreateInitialized"].(map[string]davinci.Connection); ok {
		for i := range args {
			thisArg := args[i]
			name := fmt.Sprintf("%v-%v", thisArg.Name, RandString(6))
			thisArg.Name = &name
			fmt.Printf("args[i] is %#v\n", thisArg)
			// Reading connection requires generated ConnectionID
			// For testing, the connection is created to get ConnectionID from response
			resp, err := c.CreateInitializedConnection(companyID, &thisArg)
			if err != nil {
				fmt.Println(err.Error())
				// if it's not a negative test, consider it an actual failure.
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					t.Fatal()
				}
				// stop iteration of loop if connection create failed.
				fmt.Printf("Initialized Connection from key %v creation failed, continuing to next iteration \n", i)
				continue
			}
			thisArg.ConnectionID = resp.ConnectionID
			args[i] = thisArg
			fmt.Printf("resp is: %#v\n", resp)
			if v, ok := resp.Properties["customAuth"]; ok {
				fmt.Println("customAuth is not nil, will attempt to unmarshal")
				a, _ := json.Marshal(v)
				var customAuth davinci.CustomAuth
				err := json.Unmarshal(a, &customAuth)
				if err != nil {
					fmt.Println(err.Error())
					t.Fatal()
				}
				if *customAuth.Properties.ProviderName.Value == "" {
					fmt.Println("customAuth.Properties.ProviderName.Value is empty after unmarshal")
					t.Fatal()
				}
			}
		}
	}
}

func TestConnection_Delete(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")

	if args, ok := testDataConnections["connectionsCreateInitialized"].(map[string]davinci.Connection); ok {
		for i := range args {
			thisArg := args[i]
			name := fmt.Sprintf("%v-%v", thisArg.Name, RandString(6))
			thisArg.Name = &name
			fmt.Printf("args[i] is %#v\n", thisArg)
			// Reading connection requires generated ConnectionID
			// For testing, the connection is created to get ConnectionID from response
			res, err := c.CreateInitializedConnection(companyID, &thisArg)
			if err != nil {
				fmt.Println(err.Error())
				// if it's not a negative test, consider it an actual failure.
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					t.Fatal()
				}
				// stop iteration of loop if connection create failed.
				fmt.Printf("Initialized Connection from key %v creation failed, continuing to next iteration \n", i)
				continue
			}

			fmt.Printf("res is: %v\n", res)
			resp, err := c.DeleteConnection(companyID, *thisArg.ConnectionID)
			if err != nil {
				fmt.Println(err.Error())
				// if it's not a negative test, consider it an actual failure.
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					fmt.Println("Failed Successfully")
					t.Fatal()
				}
				// stop iteration of loop if connection create failed.
				fmt.Printf("Connection from key %v failed, continuing to next iteration \n", i)
				continue
			}
			fmt.Printf("resp is: %#v\n", resp)
		}
	} else {
		t.Fatalf("No tests")
	}
}
