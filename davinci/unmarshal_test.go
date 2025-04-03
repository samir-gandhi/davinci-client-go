package davinci_test

import (
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/samir-gandhi/davinci-client-go/davinci"
	"github.com/samir-gandhi/davinci-client-go/davinci/test"
)

func TestUnmarshal_Simple(t *testing.T) {

	t.Run("unmarshal simple flow object", func(t *testing.T) {

		x := float64(1)
		y := float64(50)

		jsonString := fmt.Sprintf(`{"x":%f,"y":%f}`, x, y)

		newObj := davinci.Position{}

		err := davinci.Unmarshal([]byte(jsonString), &newObj, davinci.ExportCmpOpts{})
		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(newObj.X, &x) {
			t.Fatalf("Value X Equality failure (-want, +got):\n-%f, +%#v", x, newObj.X)
		}

		if !cmp.Equal(newObj.Y, &y) {
			t.Fatalf("Value Y Equality failure (-want, +got):\n-%f, +%#v", y, newObj.Y)
		}
	})

}

func TestUnmarshal_Nested(t *testing.T) {

	t.Run("unmarshal flow object with nested properties", func(t *testing.T) {

		companyId := "1234"
		context := "test"
		typeName := "variable"
		displayName := "My Variable"

		jsonString := fmt.Sprintf(`{"companyId":"%s","context":"%s","fields":{"type":"%s","displayName":"%s"}}`, companyId, context, typeName, displayName)

		newObj := davinci.FlowVariable{}

		err := davinci.Unmarshal([]byte(jsonString), &newObj, davinci.ExportCmpOpts{})
		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(*newObj.CompanyID, companyId) {
			t.Fatalf("Value newObj.CompanyID Equality failure (-want, +got):\n-%s, +%#v", companyId, *newObj.CompanyID)
		}

		if !cmp.Equal(*newObj.Context, context) {
			t.Fatalf("Value newObj.Context Equality failure (-want, +got):\n-%s, +%#v", context, *newObj.Context)
		}

		if !cmp.Equal(*newObj.Fields.Type, typeName) {
			t.Fatalf("Value newObj.Fields.Type Equality failure (-want, +got):\n-%s, +%#v", typeName, *newObj.Fields.Type)
		}

		if !cmp.Equal(*newObj.Fields.DisplayName, displayName) {
			t.Fatalf("Value Y Equality failure (-want, +got):\n-%s, +%#v", displayName, *newObj.Fields.DisplayName)
		}
	})

}

func TestUnmarshal_AdditionalProperties(t *testing.T) {

	t.Run("no additional properties", func(t *testing.T) {

		x := float64(1)
		y := float64(50)

		jsonString := fmt.Sprintf(`{"x":%f,"y":%f}`, x, y)

		newObj := davinci.Position{}

		err := davinci.Unmarshal([]byte(jsonString), &newObj, davinci.ExportCmpOpts{})
		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(newObj.AdditionalProperties, map[string]any{}, cmpopts.EquateEmpty()) {
			t.Fatalf("Additional Properties Equality failure (-want, +got):\n%s", cmp.Diff(map[string]any{}, newObj.AdditionalProperties))
		}
	})

	t.Run("additional properites present", func(t *testing.T) {

		x := float64(1)
		y := float64(50)

		jsonString := fmt.Sprintf(`{"x":%f,"y":%f,"custom-attribute-1":"custom-value-1","custom-attribute-2":"custom-value-2"}`, x, y)

		newObj := davinci.Position{}

		err := davinci.Unmarshal([]byte(jsonString), &newObj, davinci.ExportCmpOpts{})
		if err != nil {
			t.Fatal(err)
		}

		additionalProperties := map[string]any{
			"custom-attribute-1": "custom-value-1",
			"custom-attribute-2": "custom-value-2",
		}

		if !cmp.Equal(newObj.AdditionalProperties, additionalProperties, cmpopts.EquateEmpty()) {
			t.Fatalf("Additional Properties Equality failure (-want, +got):\n%s", cmp.Diff(additionalProperties, newObj.AdditionalProperties))
		}

	})

	t.Run("filter only additional properites", func(t *testing.T) {

		x := float64(1)
		y := float64(50)

		jsonString := fmt.Sprintf(`{"x":%f,"y":%f,"custom-attribute-1":"custom-value-1","custom-attribute-2":"custom-value-2"}`, x, y)

		newObj := davinci.Position{}

		err := davinci.Unmarshal([]byte(jsonString), &newObj, davinci.ExportCmpOpts{
			IgnoreConfig:              true,
			IgnoreDesignerCues:        true,
			IgnoreEnvironmentMetadata: true,
			IgnoreUnmappedProperties:  false,
			IgnoreVersionMetadata:     true,
			IgnoreFlowMetadata:        true,
		})
		if err != nil {
			t.Fatal(err)
		}

		expectedObj := davinci.Position{
			AdditionalProperties: map[string]any{
				"custom-attribute-1": "custom-value-1",
				"custom-attribute-2": "custom-value-2",
			},
		}

		if !cmp.Equal(newObj, expectedObj, cmpopts.EquateEmpty()) {
			t.Fatalf("Additional Properties Equality failure (-want, +got):\n%s", cmp.Diff(expectedObj, newObj))
		}

	})
}

func TestUnmarshal_Nested_AdditionalProperties(t *testing.T) {

	t.Run("no additional properties", func(t *testing.T) {

		companyId := "1234"
		context := "test"
		typeName := "variable"
		displayName := "My Variable"

		jsonString := fmt.Sprintf(`{"companyId":"%s","context":"%s","fields":{"type":"%s","displayName":"%s"}}`, companyId, context, typeName, displayName)

		newObj := davinci.FlowVariable{}

		err := davinci.Unmarshal([]byte(jsonString), &newObj, davinci.ExportCmpOpts{})
		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(newObj.Fields.AdditionalProperties, map[string]any{}, cmpopts.EquateEmpty()) {
			t.Fatalf("Additional Properties Equality failure (-want, +got):\n%s", cmp.Diff(map[string]any{}, newObj.AdditionalProperties))
		}
	})

	t.Run("additional properites present", func(t *testing.T) {

		companyId := "1234"
		context := "test"
		typeName := "variable"
		displayName := "My Variable"

		jsonString := fmt.Sprintf(`{"companyId":"%s","context":"%s","fields":{"type":"%s","displayName":"%s","custom-attribute-1":"custom-value-1","custom-attribute-2":"custom-value-2"}}`, companyId, context, typeName, displayName)

		newObj := davinci.FlowVariable{}

		err := davinci.Unmarshal([]byte(jsonString), &newObj, davinci.ExportCmpOpts{})
		if err != nil {
			t.Fatal(err)
		}

		additionalProperties := map[string]any{
			"custom-attribute-1": "custom-value-1",
			"custom-attribute-2": "custom-value-2",
		}

		if !cmp.Equal(newObj.Fields.AdditionalProperties, additionalProperties, cmpopts.EquateEmpty()) {
			t.Fatalf("Additional Properties Equality failure (-want, +got):\n%s", cmp.Diff(additionalProperties, newObj.AdditionalProperties))
		}

	})

	t.Run("filter only additional properites", func(t *testing.T) {

		companyId := "1234"
		context := "test"
		typeName := "variable"
		displayName := "My Variable"

		jsonString := fmt.Sprintf(`{"custom-attribute-1":"custom-value-1","custom-attribute-2":"custom-value-2","companyId":"%s","context":"%s","fields":{"type":"%s","displayName":"%s","custom-attribute-1":"custom-value-1","custom-attribute-2":"custom-value-2"}}`, companyId, context, typeName, displayName)

		newObj := davinci.FlowVariable{}

		err := davinci.Unmarshal([]byte(jsonString), &newObj, davinci.ExportCmpOpts{
			IgnoreConfig:              true,
			IgnoreDesignerCues:        true,
			IgnoreEnvironmentMetadata: true,
			IgnoreUnmappedProperties:  false,
			IgnoreVersionMetadata:     true,
			IgnoreFlowMetadata:        true,
		})
		if err != nil {
			t.Fatal(err)
		}

		expectedObj := davinci.FlowVariable{
			AdditionalProperties: map[string]any{
				"custom-attribute-1": "custom-value-1",
				"custom-attribute-2": "custom-value-2",
			},
			Fields: &davinci.FlowVariableFields{
				AdditionalProperties: map[string]any{
					"custom-attribute-1": "custom-value-1",
					"custom-attribute-2": "custom-value-2",
				},
			},
		}

		if !cmp.Equal(newObj, expectedObj, cmpopts.EquateEmpty()) {
			t.Fatalf("Additional Properties Equality failure (-want, +got):\n%s", cmp.Diff(expectedObj, newObj))
		}

	})
}

func TestUnmarshal_Filter(t *testing.T) {

	t.Run("ignore config", func(t *testing.T) {

		flowFile := "./test/data/test-encoding.json"

		jsonFile, err := os.Open(flowFile)
		if err != nil {
			t.Errorf("Failed to open file: %v", err)
		}

		jsonBytes, err := io.ReadAll(jsonFile)
		if err != nil {
			t.Errorf("Failed to read file: %v", err)
		}

		newObj := test.TestModel{}

		err = davinci.Unmarshal(jsonBytes, &newObj, davinci.ExportCmpOpts{
			IgnoreConfig:              true,
			IgnoreDesignerCues:        false,
			IgnoreEnvironmentMetadata: false,
			IgnoreUnmappedProperties:  false,
			IgnoreVersionMetadata:     false,
			IgnoreFlowMetadata:        false,
		})
		if err != nil {
			t.Fatal(err)
		}

		expectedObj := test.TestModel{
			AdditionalProperties: map[string]any{
				"custom-attribute-1": "custom-value-1",
				"custom-attribute-2": "custom-value-2",
			},
			Test1: func() *string { s := "test1Value"; return &s }(),
			//Test2: func() *string { s := "test2Value"; return &s }(),
			Test3: &davinci.EpochTime{time.UnixMilli(1707837216607)},
			Test4: func() *string { s := "test4Value"; return &s }(),
			Test5: func() *test.TestModel2 {
				s := test.TestModel2{
					AdditionalProperties: map[string]any{
						"custom-attribute-1": "custom-value-1",
						"custom-attribute-2": "custom-value-2",
					},
					Test1: func() *string { s := "test1SubValue"; return &s }(),
					//Test2:  func() *string { s := "test2SubValue"; return &s }(),
					Test3: &davinci.EpochTime{time.UnixMilli(1707837221226)},
					Test4: func() *string { s := "test4SubValue"; return &s }(),
					Test7: func() *float64 { s := 1e+50; return &s }(),
					//Test8:  func() *string { s := "test8SubValue"; return &s }(),
					Test9: func() *string { s := "test9SubValue"; return &s }(),
					//Test10: func() *string { s := "test10SubValue"; return &s }(),
					Test11: func() *string { s := "test11SubValue"; return &s }(),
				}
				return &s
			}(),
			Test7: func() *float64 { s := 1e-50; return &s }(),
			//Test8:  func() *string { s := "test8Value"; return &s }(),
			Test9: func() *string { s := "test9Value"; return &s }(),
			//Test10: func() *string { s := "test10Value"; return &s }(),
			Test11: func() *string { s := "test11Value"; return &s }(),
		}

		if !cmp.Equal(newObj, expectedObj, cmpopts.EquateEmpty()) {
			t.Fatalf("Filter Equality failure (-want, +got):\n%s", cmp.Diff(expectedObj, newObj))
		}

	})

	t.Run("ignore designer cues", func(t *testing.T) {

		flowFile := "./test/data/test-encoding.json"

		jsonFile, err := os.Open(flowFile)
		if err != nil {
			t.Errorf("Failed to open file: %v", err)
		}

		jsonBytes, err := io.ReadAll(jsonFile)
		if err != nil {
			t.Errorf("Failed to read file: %v", err)
		}

		newObj := test.TestModel{}

		err = davinci.Unmarshal(jsonBytes, &newObj, davinci.ExportCmpOpts{
			IgnoreConfig:              false,
			IgnoreDesignerCues:        true,
			IgnoreEnvironmentMetadata: false,
			IgnoreUnmappedProperties:  false,
			IgnoreVersionMetadata:     false,
			IgnoreFlowMetadata:        false,
		})
		if err != nil {
			t.Fatal(err)
		}

		expectedObj := test.TestModel{
			AdditionalProperties: map[string]any{
				"custom-attribute-1": "custom-value-1",
				"custom-attribute-2": "custom-value-2",
			},
			Test1: func() *string { s := "test1Value"; return &s }(),
			Test2: func() *string { s := "test2Value"; return &s }(),
			Test3: &davinci.EpochTime{time.UnixMilli(1707837216607)},
			Test4: func() *string { s := "test4Value"; return &s }(),
			Test5: func() *test.TestModel2 {
				s := test.TestModel2{
					AdditionalProperties: map[string]any{
						"custom-attribute-1": "custom-value-1",
						"custom-attribute-2": "custom-value-2",
					},
					Test1: func() *string { s := "test1SubValue"; return &s }(),
					Test2: func() *string { s := "test2SubValue"; return &s }(),
					Test3: &davinci.EpochTime{time.UnixMilli(1707837221226)},
					Test4: func() *string { s := "test4SubValue"; return &s }(),
					Test7: func() *float64 { s := 1e+50; return &s }(),
					Test8: func() *string { s := "test8SubValue"; return &s }(),
					//Test9: func() *string { s := "test9SubValue"; return &s }(),
					Test10: func() *string { s := "test10SubValue"; return &s }(),
					//Test11: func() *string { s := "test11SubValue"; return &s }(),
				}
				return &s
			}(),
			Test7: func() *float64 { s := 1e-50; return &s }(),
			Test8: func() *string { s := "test8Value"; return &s }(),
			//Test9: func() *string { s := "test9Value"; return &s }(),
			Test10: func() *string { s := "test10Value"; return &s }(),
			//Test11: func() *string { s := "test11Value"; return &s }(),
		}

		if !cmp.Equal(newObj, expectedObj, cmpopts.EquateEmpty()) {
			t.Fatalf("Filter Equality failure (-want, +got):\n%s", cmp.Diff(expectedObj, newObj))
		}

	})

	t.Run("ignore environment metadata", func(t *testing.T) {

		flowFile := "./test/data/test-encoding.json"

		jsonFile, err := os.Open(flowFile)
		if err != nil {
			t.Errorf("Failed to open file: %v", err)
		}

		jsonBytes, err := io.ReadAll(jsonFile)
		if err != nil {
			t.Errorf("Failed to read file: %v", err)
		}

		newObj := test.TestModel{}

		err = davinci.Unmarshal(jsonBytes, &newObj, davinci.ExportCmpOpts{
			IgnoreConfig:              false,
			IgnoreDesignerCues:        false,
			IgnoreEnvironmentMetadata: true,
			IgnoreUnmappedProperties:  false,
			IgnoreVersionMetadata:     false,
			IgnoreFlowMetadata:        false,
		})
		if err != nil {
			t.Fatal(err)
		}

		expectedObj := test.TestModel{
			AdditionalProperties: map[string]any{
				"custom-attribute-1": "custom-value-1",
				"custom-attribute-2": "custom-value-2",
			},
			//Test1: func() *string { s := "test1Value"; return &s }(),
			Test2: func() *string { s := "test2Value"; return &s }(),
			Test3: &davinci.EpochTime{time.UnixMilli(1707837216607)},
			//Test4: func() *string { s := "test4Value"; return &s }(),
			Test5: func() *test.TestModel2 {
				s := test.TestModel2{
					AdditionalProperties: map[string]any{
						"custom-attribute-1": "custom-value-1",
						"custom-attribute-2": "custom-value-2",
					},
					//Test1:  func() *string { s := "test1SubValue"; return &s }(),
					Test2: func() *string { s := "test2SubValue"; return &s }(),
					Test3: &davinci.EpochTime{time.UnixMilli(1707837221226)},
					//Test4:  func() *string { s := "test4SubValue"; return &s }(),
					Test7:  func() *float64 { s := 1e+50; return &s }(),
					Test8:  func() *string { s := "test8SubValue"; return &s }(),
					Test9:  func() *string { s := "test9SubValue"; return &s }(),
					Test10: func() *string { s := "test10SubValue"; return &s }(),
					Test11: func() *string { s := "test11SubValue"; return &s }(),
				}
				return &s
			}(),
			Test7:  func() *float64 { s := 1e-50; return &s }(),
			Test8:  func() *string { s := "test8Value"; return &s }(),
			Test9:  func() *string { s := "test9Value"; return &s }(),
			Test10: func() *string { s := "test10Value"; return &s }(),
			Test11: func() *string { s := "test11Value"; return &s }(),
		}

		if !cmp.Equal(newObj, expectedObj, cmpopts.EquateEmpty()) {
			t.Fatalf("Filter Equality failure (-want, +got):\n%s", cmp.Diff(expectedObj, newObj))
		}

	})

	t.Run("ignore unmapped properties", func(t *testing.T) {

		flowFile := "./test/data/test-encoding.json"

		jsonFile, err := os.Open(flowFile)
		if err != nil {
			t.Errorf("Failed to open file: %v", err)
		}

		jsonBytes, err := io.ReadAll(jsonFile)
		if err != nil {
			t.Errorf("Failed to read file: %v", err)
		}

		newObj := test.TestModel{}

		err = davinci.Unmarshal(jsonBytes, &newObj, davinci.ExportCmpOpts{
			IgnoreConfig:              false,
			IgnoreDesignerCues:        false,
			IgnoreEnvironmentMetadata: false,
			IgnoreUnmappedProperties:  true,
			IgnoreVersionMetadata:     false,
			IgnoreFlowMetadata:        false,
		})
		if err != nil {
			t.Fatal(err)
		}

		expectedObj := test.TestModel{
			//AdditionalProperties: map[string]any{
			//	"custom-attribute-1": "custom-value-1",
			//	"custom-attribute-2": "custom-value-2",
			//},
			Test1: func() *string { s := "test1Value"; return &s }(),
			Test2: func() *string { s := "test2Value"; return &s }(),
			Test3: &davinci.EpochTime{time.UnixMilli(1707837216607)},
			Test4: func() *string { s := "test4Value"; return &s }(),
			Test5: func() *test.TestModel2 {
				s := test.TestModel2{
					// AdditionalProperties: map[string]any{
					// 	"custom-attribute-1": "custom-value-1",
					// 	"custom-attribute-2": "custom-value-2",
					// },
					Test1:  func() *string { s := "test1SubValue"; return &s }(),
					Test2:  func() *string { s := "test2SubValue"; return &s }(),
					Test3:  &davinci.EpochTime{time.UnixMilli(1707837221226)},
					Test4:  func() *string { s := "test4SubValue"; return &s }(),
					Test7:  func() *float64 { s := 1e+50; return &s }(),
					Test8:  func() *string { s := "test8SubValue"; return &s }(),
					Test9:  func() *string { s := "test9SubValue"; return &s }(),
					Test10: func() *string { s := "test10SubValue"; return &s }(),
					Test11: func() *string { s := "test11SubValue"; return &s }(),
				}
				return &s
			}(),
			Test7:  func() *float64 { s := 1e-50; return &s }(),
			Test8:  func() *string { s := "test8Value"; return &s }(),
			Test9:  func() *string { s := "test9Value"; return &s }(),
			Test10: func() *string { s := "test10Value"; return &s }(),
			Test11: func() *string { s := "test11Value"; return &s }(),
		}

		if !cmp.Equal(newObj, expectedObj, cmpopts.EquateEmpty()) {
			t.Fatalf("Filter Equality failure (-want, +got):\n%s", cmp.Diff(expectedObj, newObj))
		}

	})

	t.Run("ignore version metadata", func(t *testing.T) {

		flowFile := "./test/data/test-encoding.json"

		jsonFile, err := os.Open(flowFile)
		if err != nil {
			t.Errorf("Failed to open file: %v", err)
		}

		jsonBytes, err := io.ReadAll(jsonFile)
		if err != nil {
			t.Errorf("Failed to read file: %v", err)
		}

		newObj := test.TestModel{}

		err = davinci.Unmarshal(jsonBytes, &newObj, davinci.ExportCmpOpts{
			IgnoreConfig:              false,
			IgnoreDesignerCues:        false,
			IgnoreEnvironmentMetadata: false,
			IgnoreUnmappedProperties:  false,
			IgnoreVersionMetadata:     true,
			IgnoreFlowMetadata:        false,
		})
		if err != nil {
			t.Fatal(err)
		}

		expectedObj := test.TestModel{
			AdditionalProperties: map[string]any{
				"custom-attribute-1": "custom-value-1",
				"custom-attribute-2": "custom-value-2",
			},
			Test1: func() *string { s := "test1Value"; return &s }(),
			Test2: func() *string { s := "test2Value"; return &s }(),
			//Test3: &davinci.EpochTime{time.UnixMilli(1707837216607)},
			Test4: func() *string { s := "test4Value"; return &s }(),
			Test5: func() *test.TestModel2 {
				s := test.TestModel2{
					AdditionalProperties: map[string]any{
						"custom-attribute-1": "custom-value-1",
						"custom-attribute-2": "custom-value-2",
					},
					Test1: func() *string { s := "test1SubValue"; return &s }(),
					Test2: func() *string { s := "test2SubValue"; return &s }(),
					//Test3:  &davinci.EpochTime{time.UnixMilli(1707837221226)},
					Test4:  func() *string { s := "test4SubValue"; return &s }(),
					Test7:  func() *float64 { s := 1e+50; return &s }(),
					Test8:  func() *string { s := "test8SubValue"; return &s }(),
					Test9:  func() *string { s := "test9SubValue"; return &s }(),
					Test10: func() *string { s := "test10SubValue"; return &s }(),
					Test11: func() *string { s := "test11SubValue"; return &s }(),
				}
				return &s
			}(),
			Test7:  func() *float64 { s := 1e-50; return &s }(),
			Test8:  func() *string { s := "test8Value"; return &s }(),
			Test9:  func() *string { s := "test9Value"; return &s }(),
			Test10: func() *string { s := "test10Value"; return &s }(),
			Test11: func() *string { s := "test11Value"; return &s }(),
		}

		if !cmp.Equal(newObj, expectedObj, cmpopts.EquateEmpty()) {
			t.Fatalf("Filter Equality failure (-want, +got):\n%s", cmp.Diff(expectedObj, newObj))
		}

	})

	t.Run("ignore flow metadata", func(t *testing.T) {

		flowFile := "./test/data/test-encoding.json"

		jsonFile, err := os.Open(flowFile)
		if err != nil {
			t.Errorf("Failed to open file: %v", err)
		}

		jsonBytes, err := io.ReadAll(jsonFile)
		if err != nil {
			t.Errorf("Failed to read file: %v", err)
		}

		newObj := test.TestModel{}

		err = davinci.Unmarshal(jsonBytes, &newObj, davinci.ExportCmpOpts{
			IgnoreConfig:              false,
			IgnoreDesignerCues:        false,
			IgnoreEnvironmentMetadata: false,
			IgnoreUnmappedProperties:  false,
			IgnoreVersionMetadata:     false,
			IgnoreFlowMetadata:        true,
		})
		if err != nil {
			t.Fatal(err)
		}

		expectedObj := test.TestModel{
			AdditionalProperties: map[string]any{
				"custom-attribute-1": "custom-value-1",
				"custom-attribute-2": "custom-value-2",
			},
			Test1: func() *string { s := "test1Value"; return &s }(),
			Test2: func() *string { s := "test2Value"; return &s }(),
			Test3: &davinci.EpochTime{time.UnixMilli(1707837216607)},
			Test4: func() *string { s := "test4Value"; return &s }(),
			Test5: func() *test.TestModel2 {
				s := test.TestModel2{
					AdditionalProperties: map[string]any{
						"custom-attribute-1": "custom-value-1",
						"custom-attribute-2": "custom-value-2",
					},
					Test1: func() *string { s := "test1SubValue"; return &s }(),
					Test2: func() *string { s := "test2SubValue"; return &s }(),
					Test3: &davinci.EpochTime{time.UnixMilli(1707837221226)},
					Test4: func() *string { s := "test4SubValue"; return &s }(),
					//Test7:  func() *float64 { s := 1e+50; return &s }(),
					Test8:  func() *string { s := "test8SubValue"; return &s }(),
					Test9:  func() *string { s := "test9SubValue"; return &s }(),
					Test10: func() *string { s := "test10SubValue"; return &s }(),
					Test11: func() *string { s := "test11SubValue"; return &s }(),
				}
				return &s
			}(),
			//Test7:  func() *float64 { s := 1e-50; return &s }(),
			Test8:  func() *string { s := "test8Value"; return &s }(),
			Test9:  func() *string { s := "test9Value"; return &s }(),
			Test10: func() *string { s := "test10Value"; return &s }(),
			Test11: func() *string { s := "test11Value"; return &s }(),
		}

		if !cmp.Equal(newObj, expectedObj, cmpopts.EquateEmpty()) {
			t.Fatalf("Filter Equality failure (-want, +got):\n%s", cmp.Diff(expectedObj, newObj))
		}

	})
}

func TestUnmarshal_DataTypes(t *testing.T) {

	t.Run("test data types implemented", func(t *testing.T) {

		flowFile := "./test/flows/full-basic-additionalproperties.json"

		jsonFile, err := os.Open(flowFile)
		if err != nil {
			t.Errorf("Failed to open file: %v", err)
		}

		jsonBytes, err := io.ReadAll(jsonFile)
		if err != nil {
			t.Errorf("Failed to read file: %v", err)
		}

		newObj := davinci.Flow{}

		err = davinci.Unmarshal(jsonBytes, &newObj, davinci.ExportCmpOpts{
			IgnoreConfig:              false,
			IgnoreDesignerCues:        false,
			IgnoreEnvironmentMetadata: false,
			IgnoreUnmappedProperties:  false,
			IgnoreVersionMetadata:     false,
			IgnoreFlowMetadata:        false,
		})
		if err != nil {
			t.Fatal(err)
		}

	})

}
