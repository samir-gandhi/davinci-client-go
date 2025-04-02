package davinci_test

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/samir-gandhi/davinci-client-go/davinci"
)

func TestMarshal_Simple(t *testing.T) {

	t.Run("marshal simple flow object", func(t *testing.T) {

		jsonString := `{"x":1,"y":50}`

		x := float64(1)
		y := float64(50)

		newObj := davinci.Position{
			X: &x,
			Y: &y,
		}

		bytes, err := davinci.Marshal(newObj, davinci.ExportCmpOpts{})
		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(jsonString, string(bytes)) {
			t.Fatalf("Value Equality failure (-want, +got):\n-%s, +%s", jsonString, string(bytes))
		}
	})

}

func TestMarshal_Nested(t *testing.T) {

	t.Run("marshal flow object with nested properties", func(t *testing.T) {

		companyId := "1234"
		context := "test"
		typeName := "variable"
		displayName := "My Variable"

		jsonString := fmt.Sprintf(`{"companyId":"%s","context":"%s","fields":{"displayName":"%s","type":"%s"},"name":"test","type":"test2"}`, companyId, context, displayName, typeName)

		newObj := davinci.FlowVariable{
			Name:      "test",
			Type:      "test2",
			CompanyID: &companyId,
			Context:   &context,
			Fields: &davinci.FlowVariableFields{
				Type:        &typeName,
				DisplayName: &displayName,
			},
		}

		bytes, err := davinci.Marshal(newObj, davinci.ExportCmpOpts{})
		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(jsonString, string(bytes)) {
			t.Fatalf("Value Equality failure (-want, +got):\n-%s, +%s", jsonString, string(bytes))
		}
	})

}

func TestMarshal_AdditionalProperties(t *testing.T) {

	t.Run("additional properites present", func(t *testing.T) {

		jsonString := `{"custom-attribute-1":"custom-value-1","custom-attribute-2":"custom-value-2","x":1,"y":50}`

		x := float64(1)
		y := float64(50)

		newObj := davinci.Position{
			AdditionalProperties: map[string]any{
				"custom-attribute-1": "custom-value-1",
				"custom-attribute-2": "custom-value-2",
			},
			X: &x,
			Y: &y,
		}

		bytes, err := davinci.Marshal(newObj, davinci.ExportCmpOpts{})
		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(jsonString, string(bytes)) {
			t.Fatalf("Value Equality failure (-want, +got):\n-%s, +%s", jsonString, string(bytes))
		}

	})

	t.Run("filter only additional properites", func(t *testing.T) {

		x := float64(1)
		y := float64(50)

		newObj := davinci.Position{
			AdditionalProperties: map[string]any{
				"custom-attribute-1": "custom-value-1",
				"custom-attribute-2": "custom-value-2",
			},
			X: &x,
			Y: &y,
		}

		bytes, err := davinci.Marshal(newObj, davinci.ExportCmpOpts{
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

		jsonString := `{"custom-attribute-1":"custom-value-1","custom-attribute-2":"custom-value-2"}`

		if !cmp.Equal(jsonString, string(bytes)) {
			t.Fatalf("Value Equality failure (-want, +got):\n-%s, +%s", jsonString, string(bytes))
		}

	})
}

func TestMarshal_Nested_AdditionalProperties(t *testing.T) {

	t.Run("no additional properties", func(t *testing.T) {

		companyId := "1234"
		context := "test"
		typeName := "variable"
		displayName := "My Variable"

		jsonString := fmt.Sprintf(`{"companyId":"%s","context":"%s","fields":{"displayName":"%s","type":"%s"},"name":"test","type":"test2"}`, companyId, context, displayName, typeName)

		newObj := davinci.FlowVariable{
			Name:      "test",
			Type:      "test2",
			CompanyID: &companyId,
			Context:   &context,
			Fields: &davinci.FlowVariableFields{
				Type:        &typeName,
				DisplayName: &displayName,
			},
		}

		bytes, err := davinci.Marshal(newObj, davinci.ExportCmpOpts{})
		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(jsonString, string(bytes)) {
			t.Fatalf("Value Equality failure (-want, +got):\n-%s, +%s", jsonString, string(bytes))
		}
	})

	t.Run("additional properites present", func(t *testing.T) {

		companyId := "1234"
		context := "test"
		typeName := "variable"
		displayName := "My Variable"

		jsonString := fmt.Sprintf(`{"companyId":"%s","context":"%s","fields":{"custom-attribute-1":"custom-value-1","custom-attribute-2":"custom-value-2","displayName":"%s","type":"%s"},"name":"test","type":"test2"}`, companyId, context, displayName, typeName)

		newObj := davinci.FlowVariable{
			Name:      "test",
			Type:      "test2",
			CompanyID: &companyId,
			Context:   &context,
			Fields: &davinci.FlowVariableFields{
				AdditionalProperties: map[string]any{
					"custom-attribute-1": "custom-value-1",
					"custom-attribute-2": "custom-value-2",
				},
				Type:        &typeName,
				DisplayName: &displayName,
			},
		}

		bytes, err := davinci.Marshal(newObj, davinci.ExportCmpOpts{})
		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(jsonString, string(bytes)) {
			t.Fatalf("Value Equality failure (-want, +got):\n-%s, +%s", jsonString, string(bytes))
		}

	})

	t.Run("filter only additional properites", func(t *testing.T) {

		companyId := "1234"
		context := "test"
		typeName := "variable"
		displayName := "My Variable"

		jsonString := `{"custom-attribute-1":"custom-value-1","custom-attribute-2":"custom-value-2","fields":{"custom-attribute-1":"custom-sub-value-1","custom-attribute-2":"custom-sub-value-2"}}`

		newObj := davinci.FlowVariable{
			AdditionalProperties: map[string]any{
				"custom-attribute-1": "custom-value-1",
				"custom-attribute-2": "custom-value-2",
			},
			Name:      "test",
			Type:      "test2",
			CompanyID: &companyId,
			Context:   &context,
			Fields: &davinci.FlowVariableFields{
				AdditionalProperties: map[string]any{
					"custom-attribute-1": "custom-sub-value-1",
					"custom-attribute-2": "custom-sub-value-2",
				},
				Type:        &typeName,
				DisplayName: &displayName,
			},
		}

		bytes, err := davinci.Marshal(newObj, davinci.ExportCmpOpts{
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

		if !cmp.Equal(jsonString, string(bytes)) {
			t.Fatalf("Value Equality failure (-want, +got):\n-%s, +%s", jsonString, string(bytes))
		}

	})
}

// func TestMarshal_Filter(t *testing.T) {

// t.Run("ignore config", func(t *testing.T) {

// 	flowFile := "./test/data/test-encoding.json"

// 	jsonFile, err := os.Open(flowFile)
// 	if err != nil {
// 		t.Errorf("Failed to open file: %v", err)
// 	}

// 	jsonBytes, err := io.ReadAll(jsonFile)
// 	if err != nil {
// 		t.Errorf("Failed to read file: %v", err)
// 	}

// 	newObj := TestModel{}

// 	err = Marshal(jsonBytes, &newObj, davinci.ExportCmpOpts{
// 		IgnoreConfig:              true,
// 		IgnoreDesignerCues:        false,
// 		IgnoreEnvironmentMetadata: false,
// 		IgnoreUnmappedProperties:  false,
// 		IgnoreVersionMetadata:     false,
// 		IgnoreFlowMetadata:        false,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	expectedObj := TestModel{
// 		AdditionalProperties: map[string]any{
// 			"custom-attribute-1": "custom-value-1",
// 			"custom-attribute-2": "custom-value-2",
// 		},
// 		Test1: func() *string { s := "test1Value"; return &s }(),
// 		//Test2: func() *string { s := "test2Value"; return &s }(),
// 		Test3: &EpochTime{time.UnixMilli(1707837216607)},
// 		Test4: func() *string { s := "test4Value"; return &s }(),
// 		Test5: func() *TestModel2 {
// 			s := TestModel2{
// 				AdditionalProperties: map[string]any{
// 					"custom-attribute-1": "custom-value-1",
// 					"custom-attribute-2": "custom-value-2",
// 				},
// 				Test1: func() *string { s := "test1SubValue"; return &s }(),
// 				//Test2:  func() *string { s := "test2SubValue"; return &s }(),
// 				Test3: &EpochTime{time.UnixMilli(1707837221226)},
// 				Test4: func() *string { s := "test4SubValue"; return &s }(),
// 				Test7: func() *float64 { s := 1e+50; return &s }(),
// 				//Test8:  func() *string { s := "test8SubValue"; return &s }(),
// 				Test9: func() *string { s := "test9SubValue"; return &s }(),
// 				//Test10: func() *string { s := "test10SubValue"; return &s }(),
// 				Test11: func() *string { s := "test11SubValue"; return &s }(),
// 			}
// 			return &s
// 		}(),
// 		Test7: func() *float64 { s := 1e-50; return &s }(),
// 		//Test8:  func() *string { s := "test8Value"; return &s }(),
// 		Test9: func() *string { s := "test9Value"; return &s }(),
// 		//Test10: func() *string { s := "test10Value"; return &s }(),
// 		Test11: func() *string { s := "test11Value"; return &s }(),
// 	}

// 	if !cmp.Equal(newObj, expectedObj, cmpopts.EquateEmpty()) {
// 		t.Fatalf("Filter Equality failure (-want, +got):\n%s", cmp.Diff(expectedObj, newObj))
// 	}

// })

// 	t.Run("ignore designer cues", func(t *testing.T) {

// 		flowFile := "./test/data/test-encoding.json"

// 		jsonFile, err := os.Open(flowFile)
// 		if err != nil {
// 			t.Errorf("Failed to open file: %v", err)
// 		}

// 		jsonBytes, err := io.ReadAll(jsonFile)
// 		if err != nil {
// 			t.Errorf("Failed to read file: %v", err)
// 		}

// 		newObj := TestModel{}

// 		err = Marshal(jsonBytes, &newObj, davinci.ExportCmpOpts{
// 			IgnoreConfig:              false,
// 			IgnoreDesignerCues:        true,
// 			IgnoreEnvironmentMetadata: false,
// 			IgnoreUnmappedProperties:  false,
// 			IgnoreVersionMetadata:     false,
// 			IgnoreFlowMetadata:        false,
// 		})
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		expectedObj := TestModel{
// 			AdditionalProperties: map[string]any{
// 				"custom-attribute-1": "custom-value-1",
// 				"custom-attribute-2": "custom-value-2",
// 			},
// 			Test1: func() *string { s := "test1Value"; return &s }(),
// 			Test2: func() *string { s := "test2Value"; return &s }(),
// 			Test3: &EpochTime{time.UnixMilli(1707837216607)},
// 			Test4: func() *string { s := "test4Value"; return &s }(),
// 			Test5: func() *TestModel2 {
// 				s := TestModel2{
// 					AdditionalProperties: map[string]any{
// 						"custom-attribute-1": "custom-value-1",
// 						"custom-attribute-2": "custom-value-2",
// 					},
// 					Test1: func() *string { s := "test1SubValue"; return &s }(),
// 					Test2: func() *string { s := "test2SubValue"; return &s }(),
// 					Test3: &EpochTime{time.UnixMilli(1707837221226)},
// 					Test4: func() *string { s := "test4SubValue"; return &s }(),
// 					Test7: func() *float64 { s := 1e+50; return &s }(),
// 					Test8: func() *string { s := "test8SubValue"; return &s }(),
// 					//Test9: func() *string { s := "test9SubValue"; return &s }(),
// 					Test10: func() *string { s := "test10SubValue"; return &s }(),
// 					//Test11: func() *string { s := "test11SubValue"; return &s }(),
// 				}
// 				return &s
// 			}(),
// 			Test7: func() *float64 { s := 1e-50; return &s }(),
// 			Test8: func() *string { s := "test8Value"; return &s }(),
// 			//Test9: func() *string { s := "test9Value"; return &s }(),
// 			Test10: func() *string { s := "test10Value"; return &s }(),
// 			//Test11: func() *string { s := "test11Value"; return &s }(),
// 		}

// 		if !cmp.Equal(newObj, expectedObj, cmpopts.EquateEmpty()) {
// 			t.Fatalf("Filter Equality failure (-want, +got):\n%s", cmp.Diff(expectedObj, newObj))
// 		}

// 	})

// 	t.Run("ignore environment metadata", func(t *testing.T) {

// 		flowFile := "./test/data/test-encoding.json"

// 		jsonFile, err := os.Open(flowFile)
// 		if err != nil {
// 			t.Errorf("Failed to open file: %v", err)
// 		}

// 		jsonBytes, err := io.ReadAll(jsonFile)
// 		if err != nil {
// 			t.Errorf("Failed to read file: %v", err)
// 		}

// 		newObj := TestModel{}

// 		err = Marshal(jsonBytes, &newObj, davinci.ExportCmpOpts{
// 			IgnoreConfig:              false,
// 			IgnoreDesignerCues:        false,
// 			IgnoreEnvironmentMetadata: true,
// 			IgnoreUnmappedProperties:  false,
// 			IgnoreVersionMetadata:     false,
// 			IgnoreFlowMetadata:        false,
// 		})
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		expectedObj := TestModel{
// 			AdditionalProperties: map[string]any{
// 				"custom-attribute-1": "custom-value-1",
// 				"custom-attribute-2": "custom-value-2",
// 			},
// 			//Test1: func() *string { s := "test1Value"; return &s }(),
// 			Test2: func() *string { s := "test2Value"; return &s }(),
// 			Test3: &EpochTime{time.UnixMilli(1707837216607)},
// 			//Test4: func() *string { s := "test4Value"; return &s }(),
// 			Test5: func() *TestModel2 {
// 				s := TestModel2{
// 					AdditionalProperties: map[string]any{
// 						"custom-attribute-1": "custom-value-1",
// 						"custom-attribute-2": "custom-value-2",
// 					},
// 					//Test1:  func() *string { s := "test1SubValue"; return &s }(),
// 					Test2: func() *string { s := "test2SubValue"; return &s }(),
// 					Test3: &EpochTime{time.UnixMilli(1707837221226)},
// 					//Test4:  func() *string { s := "test4SubValue"; return &s }(),
// 					Test7:  func() *float64 { s := 1e+50; return &s }(),
// 					Test8:  func() *string { s := "test8SubValue"; return &s }(),
// 					Test9:  func() *string { s := "test9SubValue"; return &s }(),
// 					Test10: func() *string { s := "test10SubValue"; return &s }(),
// 					Test11: func() *string { s := "test11SubValue"; return &s }(),
// 				}
// 				return &s
// 			}(),
// 			Test7:  func() *float64 { s := 1e-50; return &s }(),
// 			Test8:  func() *string { s := "test8Value"; return &s }(),
// 			Test9:  func() *string { s := "test9Value"; return &s }(),
// 			Test10: func() *string { s := "test10Value"; return &s }(),
// 			Test11: func() *string { s := "test11Value"; return &s }(),
// 		}

// 		if !cmp.Equal(newObj, expectedObj, cmpopts.EquateEmpty()) {
// 			t.Fatalf("Filter Equality failure (-want, +got):\n%s", cmp.Diff(expectedObj, newObj))
// 		}

// 	})

// 	t.Run("ignore unmapped properties", func(t *testing.T) {

// 		flowFile := "./test/data/test-encoding.json"

// 		jsonFile, err := os.Open(flowFile)
// 		if err != nil {
// 			t.Errorf("Failed to open file: %v", err)
// 		}

// 		jsonBytes, err := io.ReadAll(jsonFile)
// 		if err != nil {
// 			t.Errorf("Failed to read file: %v", err)
// 		}

// 		newObj := TestModel{}

// 		err = Marshal(jsonBytes, &newObj, davinci.ExportCmpOpts{
// 			IgnoreConfig:              false,
// 			IgnoreDesignerCues:        false,
// 			IgnoreEnvironmentMetadata: false,
// 			IgnoreUnmappedProperties:  true,
// 			IgnoreVersionMetadata:     false,
// 			IgnoreFlowMetadata:        false,
// 		})
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		expectedObj := TestModel{
// 			//AdditionalProperties: map[string]any{
// 			//	"custom-attribute-1": "custom-value-1",
// 			//	"custom-attribute-2": "custom-value-2",
// 			//},
// 			Test1: func() *string { s := "test1Value"; return &s }(),
// 			Test2: func() *string { s := "test2Value"; return &s }(),
// 			Test3: &EpochTime{time.UnixMilli(1707837216607)},
// 			Test4: func() *string { s := "test4Value"; return &s }(),
// 			Test5: func() *TestModel2 {
// 				s := TestModel2{
// 					// AdditionalProperties: map[string]any{
// 					// 	"custom-attribute-1": "custom-value-1",
// 					// 	"custom-attribute-2": "custom-value-2",
// 					// },
// 					Test1:  func() *string { s := "test1SubValue"; return &s }(),
// 					Test2:  func() *string { s := "test2SubValue"; return &s }(),
// 					Test3:  &EpochTime{time.UnixMilli(1707837221226)},
// 					Test4:  func() *string { s := "test4SubValue"; return &s }(),
// 					Test7:  func() *float64 { s := 1e+50; return &s }(),
// 					Test8:  func() *string { s := "test8SubValue"; return &s }(),
// 					Test9:  func() *string { s := "test9SubValue"; return &s }(),
// 					Test10: func() *string { s := "test10SubValue"; return &s }(),
// 					Test11: func() *string { s := "test11SubValue"; return &s }(),
// 				}
// 				return &s
// 			}(),
// 			Test7:  func() *float64 { s := 1e-50; return &s }(),
// 			Test8:  func() *string { s := "test8Value"; return &s }(),
// 			Test9:  func() *string { s := "test9Value"; return &s }(),
// 			Test10: func() *string { s := "test10Value"; return &s }(),
// 			Test11: func() *string { s := "test11Value"; return &s }(),
// 		}

// 		if !cmp.Equal(newObj, expectedObj, cmpopts.EquateEmpty()) {
// 			t.Fatalf("Filter Equality failure (-want, +got):\n%s", cmp.Diff(expectedObj, newObj))
// 		}

// 	})

// 	t.Run("ignore version metadata", func(t *testing.T) {

// 		flowFile := "./test/data/test-encoding.json"

// 		jsonFile, err := os.Open(flowFile)
// 		if err != nil {
// 			t.Errorf("Failed to open file: %v", err)
// 		}

// 		jsonBytes, err := io.ReadAll(jsonFile)
// 		if err != nil {
// 			t.Errorf("Failed to read file: %v", err)
// 		}

// 		newObj := TestModel{}

// 		err = Marshal(jsonBytes, &newObj, davinci.ExportCmpOpts{
// 			IgnoreConfig:              false,
// 			IgnoreDesignerCues:        false,
// 			IgnoreEnvironmentMetadata: false,
// 			IgnoreUnmappedProperties:  false,
// 			IgnoreVersionMetadata:     true,
// 			IgnoreFlowMetadata:        false,
// 		})
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		expectedObj := TestModel{
// 			AdditionalProperties: map[string]any{
// 				"custom-attribute-1": "custom-value-1",
// 				"custom-attribute-2": "custom-value-2",
// 			},
// 			Test1: func() *string { s := "test1Value"; return &s }(),
// 			Test2: func() *string { s := "test2Value"; return &s }(),
// 			//Test3: &EpochTime{time.UnixMilli(1707837216607)},
// 			Test4: func() *string { s := "test4Value"; return &s }(),
// 			Test5: func() *TestModel2 {
// 				s := TestModel2{
// 					AdditionalProperties: map[string]any{
// 						"custom-attribute-1": "custom-value-1",
// 						"custom-attribute-2": "custom-value-2",
// 					},
// 					Test1: func() *string { s := "test1SubValue"; return &s }(),
// 					Test2: func() *string { s := "test2SubValue"; return &s }(),
// 					//Test3:  &EpochTime{time.UnixMilli(1707837221226)},
// 					Test4:  func() *string { s := "test4SubValue"; return &s }(),
// 					Test7:  func() *float64 { s := 1e+50; return &s }(),
// 					Test8:  func() *string { s := "test8SubValue"; return &s }(),
// 					Test9:  func() *string { s := "test9SubValue"; return &s }(),
// 					Test10: func() *string { s := "test10SubValue"; return &s }(),
// 					Test11: func() *string { s := "test11SubValue"; return &s }(),
// 				}
// 				return &s
// 			}(),
// 			Test7:  func() *float64 { s := 1e-50; return &s }(),
// 			Test8:  func() *string { s := "test8Value"; return &s }(),
// 			Test9:  func() *string { s := "test9Value"; return &s }(),
// 			Test10: func() *string { s := "test10Value"; return &s }(),
// 			Test11: func() *string { s := "test11Value"; return &s }(),
// 		}

// 		if !cmp.Equal(newObj, expectedObj, cmpopts.EquateEmpty()) {
// 			t.Fatalf("Filter Equality failure (-want, +got):\n%s", cmp.Diff(expectedObj, newObj))
// 		}

// 	})

// 	t.Run("ignore flow metadata", func(t *testing.T) {

// 		flowFile := "./test/data/test-encoding.json"

// 		jsonFile, err := os.Open(flowFile)
// 		if err != nil {
// 			t.Errorf("Failed to open file: %v", err)
// 		}

// 		jsonBytes, err := io.ReadAll(jsonFile)
// 		if err != nil {
// 			t.Errorf("Failed to read file: %v", err)
// 		}

// 		newObj := TestModel{}

// 		err = Marshal(jsonBytes, &newObj, davinci.ExportCmpOpts{
// 			IgnoreConfig:              false,
// 			IgnoreDesignerCues:        false,
// 			IgnoreEnvironmentMetadata: false,
// 			IgnoreUnmappedProperties:  false,
// 			IgnoreVersionMetadata:     false,
// 			IgnoreFlowMetadata:        true,
// 		})
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		expectedObj := TestModel{
// 			AdditionalProperties: map[string]any{
// 				"custom-attribute-1": "custom-value-1",
// 				"custom-attribute-2": "custom-value-2",
// 			},
// 			Test1: func() *string { s := "test1Value"; return &s }(),
// 			Test2: func() *string { s := "test2Value"; return &s }(),
// 			Test3: &EpochTime{time.UnixMilli(1707837216607)},
// 			Test4: func() *string { s := "test4Value"; return &s }(),
// 			Test5: func() *TestModel2 {
// 				s := TestModel2{
// 					AdditionalProperties: map[string]any{
// 						"custom-attribute-1": "custom-value-1",
// 						"custom-attribute-2": "custom-value-2",
// 					},
// 					Test1: func() *string { s := "test1SubValue"; return &s }(),
// 					Test2: func() *string { s := "test2SubValue"; return &s }(),
// 					Test3: &EpochTime{time.UnixMilli(1707837221226)},
// 					Test4: func() *string { s := "test4SubValue"; return &s }(),
// 					//Test7:  func() *float64 { s := 1e+50; return &s }(),
// 					Test8:  func() *string { s := "test8SubValue"; return &s }(),
// 					Test9:  func() *string { s := "test9SubValue"; return &s }(),
// 					Test10: func() *string { s := "test10SubValue"; return &s }(),
// 					Test11: func() *string { s := "test11SubValue"; return &s }(),
// 				}
// 				return &s
// 			}(),
// 			//Test7:  func() *float64 { s := 1e-50; return &s }(),
// 			Test8:  func() *string { s := "test8Value"; return &s }(),
// 			Test9:  func() *string { s := "test9Value"; return &s }(),
// 			Test10: func() *string { s := "test10Value"; return &s }(),
// 			Test11: func() *string { s := "test11Value"; return &s }(),
// 		}

// 		if !cmp.Equal(newObj, expectedObj, cmpopts.EquateEmpty()) {
// 			t.Fatalf("Filter Equality failure (-want, +got):\n%s", cmp.Diff(expectedObj, newObj))
// 		}

// 	})
// }

func TestMarshal_DataTypes(t *testing.T) {

	t.Run("test data types implemented", func(t *testing.T) {

		flowFiles := []string{
			"./test/flows/export-version.json",
			"./test/flows/full-basic-additionalproperties.json",
		}

		for _, flowFile := range flowFiles {
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

			bytes, err := davinci.Marshal(newObj, davinci.ExportCmpOpts{
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

			newFlowObj := davinci.Flow{}

			err = davinci.Unmarshal(bytes, &newFlowObj, davinci.ExportCmpOpts{
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

			if !cmp.Equal(newObj, newFlowObj) {
				t.Fatalf("Value Equality failure (-want, +got):\n%s", cmp.Diff(newObj, newFlowObj))
			}
		}

	})

}
