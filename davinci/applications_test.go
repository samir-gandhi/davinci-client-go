package davinci_test

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/samir-gandhi/davinci-client-go/davinci"
)

var testDataApps = map[string]any{
	"params": map[string]davinci.Params{
		"a": {
			Page:  "0",
			Limit: "1",
		},
		"bNeg": {
			Page:  "100",
			Limit: "10",
		},
		"c": {},
	},
	"appsCreate": map[string]davinci.App{
		"aCreate": {
			Name: "aCreate",
		},
		"bCreate": {
			Name: "B CREATE",
		},
		"cCreateNeg": {
			Name: "",
		},
	},
	"appsRead": map[string]davinci.App{
		"aCreate": {
			Name: "aCreate",
		},
		"bCreate": {
			Name: "B CREATE",
		},
		"cCreateNeg": {
			Name: "",
		},
	},
	"appsUpdate": map[string]davinci.AppUpdate{
		"aCreate": {
			Name: "aCreate",
			Oauth: &davinci.Oauth{
				Enabled: true,
				Values: &davinci.OauthValues{
					Enabled:       true,
					AllowedScopes: []string{"openid", "profile", "flow_analytics"},
				},
			},
			Policies: []davinci.Policy{{
				PolicyFlows: []davinci.PolicyFlow{{
					FlowID:    "1764a19731a067d8b56f0c2d250cd9ea",
					VersionID: -1,
					Weight: func() *int {
						i := 100
						return &i
					}(),
				}},
				Name: func() *string {
					s := "aCreatePolicy"
					return &s
				}(),
			}},
			UserPortal: &davinci.UserPortal{
				Values: &davinci.UserPortalValues{
					UpTitle: "thisIsTitle",
				},
			},
		},
		"bCreate": {
			Name: "BCREATE",
			Saml: &davinci.Saml{
				Values: &davinci.SamlValues{
					Enabled: true,
					RedirectURI: func() *string {
						s := "https://example.com"
						return &s
					}(),
				},
			},
		},
		"cCreateNeg": {
			Name: "",
		},
	},
}

func TestApplications_Read(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	args, _ := testDataApps["params"].(map[string]davinci.Params)
	for i, thisArg := range args {
		testName := i
		t.Run(testName, func(t *testing.T) {

			resp, err := c.ReadApplications(companyID, &thisArg)
			if err != nil {
				fmt.Println(err.Error())
				log.Printf("Failed Successfully\n")
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					log.Printf("failed to get flows with params: %v \n Error is: %v", args, err)
					t.Fatal()
				}
			}
			if resp != nil {
				//log.Printf("Apps Returned Successfully\n appId is: %+v \n", *resp[0].AppID)
				log.Printf("Apps Returned Successfully\n appId is: %#v \n", resp)
			}

		})
	}
}

func TestApplication_Create(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	args, _ := testDataApps["appsCreate"].(map[string]davinci.App)
	for i, thisArg := range args {
		testName := i
		t.Run(testName, func(t *testing.T) {
			name := fmt.Sprintf("%v-%v", thisArg.Name, RandString(6))
			thisArg.Name = name

			resp, err := c.CreateApplication(companyID, thisArg.Name)
			if err != nil {
				fmt.Println(err.Error())
				log.Printf("Failed Successfully\n")
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					log.Printf("failed with params: %v \n Error is: %v", args, err)
					t.Fatal()
				}
			}
			if resp != nil {
				log.Printf("Apps Returned Successfully\n appId is: %+v \n", resp.AppID)
			}
		})
	}
}

func TestApplication_Read(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	args, _ := testDataApps["appsRead"].(map[string]davinci.App)
	for i, thisArg := range args {
		testName := i
		t.Run(testName, func(t *testing.T) {
			name := fmt.Sprintf("%v-%v", thisArg.Name, RandString(6))
			thisArg.Name = name

			resp, err := c.CreateApplication(companyID, thisArg.Name)
			if err != nil {
				fmt.Println(err.Error())
				log.Printf("Failed Successfully\n")
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					log.Printf("failed with params: %v \n Error is: %v", args, err)
					t.Fatal()
				}
			}
			if resp != nil {
				res, err := c.ReadApplication(companyID, *resp.AppID)
				if err != nil {
					fmt.Println(err.Error())
					log.Printf("Failed Successfully\n")
					if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
						log.Printf("failed with params: %v \n Error is: %v", args, err)
						t.Fatal()
					}
				}
				if res != nil {
					log.Printf("Apps Returned Successfully\n appId is: %+v \n", res.AppID)
				}
			}

		})
	}
}

func TestApplication_Update(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	args, _ := testDataApps["appsUpdate"].(map[string]davinci.AppUpdate)
	for i, thisArg := range args {
		testName := i
		t.Run(testName, func(t *testing.T) {
			name := fmt.Sprintf("%v-%v", thisArg.Name, RandString(6))
			thisArg.Name = name

			resp, err := c.CreateApplication(companyID, thisArg.Name)
			if err != nil {
				fmt.Println(err.Error())
				log.Printf("Failed Successfully\n")
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					log.Printf("App Create failed with params: %v \n Error is: %v", args, err)
					t.Fatal()
				}
			}
			if resp != nil {
				thisArg.AppID = resp.AppID
				res, err := c.UpdateApplication(companyID, &thisArg)
				if err != nil {
					fmt.Println(err.Error())
					log.Printf("Failed Successfully\n")
					if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
						log.Printf("failed with params: %v \n Error is: %v", args, err)
						t.Fatal()
					}
				}
				if res != nil {
					log.Printf("Apps Returned Successfully\n appId is: %+v \n", resp.AppID)
				}
			}

		})
	}
}

func TestApplication_Delete(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	args, _ := testDataApps["appsUpdate"].(map[string]davinci.AppUpdate)
	for i, thisArg := range args {
		testName := i
		t.Run(testName, func(t *testing.T) {
			name := fmt.Sprintf("%v-%v", thisArg.Name, RandString(6))
			thisArg.Name = name

			resp, err := c.CreateApplication(companyID, thisArg.Name)
			if err != nil {
				fmt.Println(err.Error())
				log.Printf("Failed Successfully\n")
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					log.Printf("App Create failed with params: %v \n Error is: %v", args, err)
					t.Fatal()
				}
			}
			if resp != nil {
				thisArg.AppID = resp.AppID
				res, err := c.DeleteApplication(companyID, *thisArg.AppID)
				fmt.Println(res)
				if err != nil && *res.Message != "App successfully removed" {
					fmt.Println(err.Error())
					log.Printf("Failed Successfully\n")
					if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
						log.Printf("failed with params: %v \n Error is: %v", args, err)
						t.Fatal()
					}
				}
				if res != nil {
					log.Printf("App Deleted Successfully\n message is: %v \n", res)
				}
			}

		})
	}
}
