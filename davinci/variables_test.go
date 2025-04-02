package davinci_test

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/samir-gandhi/davinci-client-go/davinci"
	"github.com/samir-gandhi/davinci-client-go/tools"
)

var testDataVars = map[string]any{
	"params": map[string]davinci.Params{
		"a": {
			Page:  "0",
			Limit: "1",
		},
		"bNeg": {Page: "100",
			Limit: "10",
		},
		"c": {},
	},
}

func TestVariables_Read(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	args, _ := testDataVars["params"].(map[string]davinci.Params)
	for i, thisArg := range args {
		testName := i
		t.Run(testName, func(t *testing.T) {

			resp, err := c.ReadVariables(companyID, &thisArg)
			if err != nil {
				log.Printf("Failed Successfully: %v\n", err.Error())
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					log.Printf("failed to get flows with params: %v \n Error is: %v", args, err)
					t.Fatal()
				}
			}
			if resp != nil {
				var l []string
				for j := range resp {
					l = append(l, j)
				}
				log.Printf("Vars Returned Successfully\n vars returned is: %v\n", l)
			}

		})
	}
}

func testReadVariable_inputs() []string {
	a := []string{"flow##SK##flowInstance"}
	return a
}

func TestVariable_Read(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	args := testReadVariable_inputs()
	for _, thisArg := range args {
		t.Run(thisArg, func(t *testing.T) {

			resp, err := c.ReadVariable(companyID, thisArg)
			if err != nil {
				log.Printf("Failed Successfully: %v\n", err.Error())
				if !(strings.Contains(thisArg, "neg")) && !(strings.Contains(thisArg, "Neg")) {
					log.Printf("failed to get flows with params: %v \n Error is: %v", args, err)
					t.Fatal()
				}
			}
			if resp != nil {
				var l []string
				for j := range resp {
					l = append(l, j)
				}
				log.Printf("Vars Returned Successfully\n vars returned is: %v\n", l)
			}

		})
	}
}

func testCreateVariables_inputs(resourceName string) (varMap map[string]davinci.VariablePayload) {
	varMap = map[string]davinci.VariablePayload{
		"company": {
			Name: func() *string {
				s := fmt.Sprintf("company-%s", resourceName)
				return &s
			}(),
			Context: "company",
			Value: func() *string {
				s := fmt.Sprintf(`value-%s`, resourceName)
				return &s
			}(),
			Description: func() *string {
				s := fmt.Sprintf("description-%s", resourceName)
				return &s
			}(),
			Type: "string",
			Mutable: func() *bool {
				b := true
				return &b
			}(),
		},
		"flow": {
			Name: func() *string {
				s := fmt.Sprintf("flow-%s", resourceName)
				return &s
			}(),
			Context: "flowInstance",
			Value: func() *string {
				s := fmt.Sprintf(`value-%s`, resourceName)
				return &s
			}(),
			Description: func() *string {
				s := fmt.Sprintf("description-%s", resourceName)
				return &s
			}(),
			Type: "string",
			Mutable: func() *bool {
				b := true
				return &b
			}(),
		},
		"flowNeg": {
			Name: func() *string {
				s := fmt.Sprintf("flow-%s", resourceName)
				return &s
			}(),
			Context: "badContext",
			Value: func() *string {
				s := fmt.Sprintf(`value-%s`, resourceName)
				return &s
			}(),
			Description: func() *string {
				s := fmt.Sprintf("description-%s", resourceName)
				return &s
			}(),
			Type: "string",
			Mutable: func() *bool {
				b := true
				return &b
			}(),
		},
	}
	return varMap
}

func TestVariable_Create(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	resourceName := tools.RandomString(10)
	args := testCreateVariables_inputs(resourceName)
	for i, thisArg := range args {
		testName := i
		t.Run(testName, func(t *testing.T) {

			resp, err := c.CreateVariable(companyID, &thisArg)
			if err != nil {
				log.Printf("Failed Successfully: %v\n", err.Error())
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					log.Printf("failed to get flows with params: %v \n Error is: %v", args, err)
					t.Fatal()
				}
			}
			if resp != nil {
				log.Printf("Var created successfully\n vars returned is: %v\n", resp)
			}

		})
	}
}

func testUpdateVariables_inputs(resourceName string) (varMap map[string]davinci.VariablePayload) {
	varMap = map[string]davinci.VariablePayload{
		"company": {
			Name: func() *string {
				s := fmt.Sprintf("company-%s", resourceName)
				return &s
			}(),
			Context: "company",
			Value: func() *string {
				s := fmt.Sprintf(`value-%s`, resourceName)
				return &s
			}(),
			Description: func() *string {
				s := fmt.Sprintf("description-%s", resourceName)
				return &s
			}(),
			Type: "string",
			Mutable: func() *bool {
				b := true
				return &b
			}(),
		},
		"flow": {
			Name: func() *string {
				s := fmt.Sprintf("flow-%s", resourceName)
				return &s
			}(),
			Context: "flowInstance",
			Value: func() *string {
				s := fmt.Sprintf(`value-%s`, resourceName)
				return &s
			}(),
			Description: func() *string {
				s := fmt.Sprintf("description-%s", resourceName)
				return &s
			}(),
			Type: "string",
			Mutable: func() *bool {
				b := true
				return &b
			}(),
		},
		"flowNeg": {
			Name: func() *string {
				s := fmt.Sprintf("flow-%s", resourceName)
				return &s
			}(),
			Context: "badContext",
			Value: func() *string {
				s := fmt.Sprintf(`value-%s`, resourceName)
				return &s
			}(),
			Description: func() *string {
				s := fmt.Sprintf("description-%s", resourceName)
				return &s
			}(),
			Type: "string",
			Mutable: func() *bool {
				b := true
				return &b
			}(),
		},
	}
	return varMap
}

func TestVariable_Update(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	resourceName := tools.RandomString(10)
	args := testUpdateVariables_inputs(resourceName)
	for i, thisArg := range args {
		testName := i
		t.Run(testName, func(t *testing.T) {

			resp, err := c.CreateVariable(companyID, &thisArg)
			if err != nil {
				log.Printf("Failed Successfully: %v\n", err.Error())
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					log.Printf("failed to get flows with params: %v \n Error is: %v", args, err)
					t.Fatal()
				}
			}
			if resp != nil {
				// log.Printf("Var created successfully\n vars returned is: %v\n", resp)
				description := *thisArg.Description + tools.RandomString(10)
				thisArg.Description = &description
				update, err := c.UpdateVariable(companyID, &thisArg)
				if err != nil {
					log.Printf("Failed Successfully: %v\n", err.Error())
					if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
						log.Printf("failed to get flows with params: %v \n Error is: %v", args, err)
						t.Fatal()
					}
				}
				if update != nil {
					log.Printf("Var updated successfully\n var returned is: %v\n", resp)
				}
			}

		})
	}
}

func testDeleteVariables_inputs(resourceName string) (varMap map[string]davinci.VariablePayload) {
	varMap = map[string]davinci.VariablePayload{
		"company": {
			Name: func() *string {
				s := fmt.Sprintf("company-%s", resourceName)
				return &s
			}(),
			Context: "company",
			Value: func() *string {
				s := fmt.Sprintf(`value-%s`, resourceName)
				return &s
			}(),
			Description: func() *string {
				s := fmt.Sprintf("description-%s", resourceName)
				return &s
			}(),
			Type: "string",
			Mutable: func() *bool {
				b := true
				return &b
			}(),
		},
		"flow": {
			Name: func() *string {
				s := fmt.Sprintf("flow-%s", resourceName)
				return &s
			}(),
			Context: "flowInstance",
			Value: func() *string {
				s := fmt.Sprintf(`value-%s`, resourceName)
				return &s
			}(),
			Description: func() *string {
				s := fmt.Sprintf("description-%s", resourceName)
				return &s
			}(),
			Type: "string",
			Mutable: func() *bool {
				b := true
				return &b
			}(),
		},
		"flowNeg": {
			Name: func() *string {
				s := fmt.Sprintf("flow-%s", resourceName)
				return &s
			}(),
			Context: "badContext",
			Value: func() *string {
				s := fmt.Sprintf(`value-%s`, resourceName)
				return &s
			}(),
			Description: func() *string {
				s := fmt.Sprintf("description-%s", resourceName)
				return &s
			}(),
			Type: "string",
			Mutable: func() *bool {
				b := true
				return &b
			}(),
		},
	}
	return varMap
}

func TestVariable_Delete(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	companyID := os.Getenv("PINGONE_TARGET_ENVIRONMENT_ID")
	resourceName := tools.RandomString(10)
	args := testDeleteVariables_inputs(resourceName)
	for i, thisArg := range args {
		testName := i
		t.Run(testName, func(t *testing.T) {

			resp, err := c.CreateVariable(companyID, &thisArg)
			if err != nil {
				log.Printf("Failed Successfully: %v\n", err.Error())
				if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
					log.Printf("failed to get flows with params: %v \n Error is: %v", args, err)
					t.Fatal()
				}
			}
			if resp != nil {
				// log.Printf("Var created successfully\n vars returned is: %v\n", resp)
				update, err := c.DeleteVariable(companyID, fmt.Sprintf(`%s##SK##%s`, *thisArg.Name, thisArg.Context))
				if err != nil {
					log.Printf("Failed Successfully: %v\n", err.Error())
					if !(strings.Contains(i, "neg")) && !(strings.Contains(i, "Neg")) {
						log.Printf("failed to get flows with params: %v \n Error is: %v", args, err)
						t.Fatal()
					}
				}
				if update != nil {
					log.Printf("Var updated successfully\n var returned is: %v\n", resp)
				}
			}

		})
	}
}
