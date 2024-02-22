package davinci

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestUnmarshal_Simple(t *testing.T) {

	t.Run("unmarshal simple flow object", func(t *testing.T) {

		x := float64(1)
		y := float64(50)

		jsonString := fmt.Sprintf(`{"x":%f,"y":%f}`, x, y)

		newObj := Position{}

		err := Unmarshal([]byte(jsonString), &newObj, ExportCmpOpts{})
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

		newObj := FlowVariable{}

		err := Unmarshal([]byte(jsonString), &newObj, ExportCmpOpts{})
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

		newObj := Position{}

		err := Unmarshal([]byte(jsonString), &newObj, ExportCmpOpts{})
		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(newObj.AdditionalProperties, map[string]interface{}{}, cmpopts.EquateEmpty()) {
			t.Fatalf("Additional Properties Equality failure (-want, +got):\n%s", cmp.Diff(map[string]interface{}{}, newObj.AdditionalProperties))
		}
	})

	t.Run("additional properites present", func(t *testing.T) {

		x := float64(1)
		y := float64(50)

		jsonString := fmt.Sprintf(`{"x":%f,"y":%f,"custom-attribute-1":"custom-value-1","custom-attribute-2":"custom-value-2"}`, x, y)

		newObj := Position{}

		err := Unmarshal([]byte(jsonString), &newObj, ExportCmpOpts{})
		if err != nil {
			t.Fatal(err)
		}

		additionalProperties := map[string]interface{}{
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

		newObj := Position{}

		err := Unmarshal([]byte(jsonString), &newObj, ExportCmpOpts{
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

		expectedObj := Position{
			AdditionalProperties: map[string]interface{}{
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

		newObj := FlowVariable{}

		err := Unmarshal([]byte(jsonString), &newObj, ExportCmpOpts{})
		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(newObj.Fields.AdditionalProperties, map[string]interface{}{}, cmpopts.EquateEmpty()) {
			t.Fatalf("Additional Properties Equality failure (-want, +got):\n%s", cmp.Diff(map[string]interface{}{}, newObj.AdditionalProperties))
		}
	})

	t.Run("additional properites present", func(t *testing.T) {

		companyId := "1234"
		context := "test"
		typeName := "variable"
		displayName := "My Variable"

		jsonString := fmt.Sprintf(`{"companyId":"%s","context":"%s","fields":{"type":"%s","displayName":"%s","custom-attribute-1":"custom-value-1","custom-attribute-2":"custom-value-2"}}`, companyId, context, typeName, displayName)

		newObj := FlowVariable{}

		err := Unmarshal([]byte(jsonString), &newObj, ExportCmpOpts{})
		if err != nil {
			t.Fatal(err)
		}

		additionalProperties := map[string]interface{}{
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

		newObj := FlowVariable{}

		err := Unmarshal([]byte(jsonString), &newObj, ExportCmpOpts{
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

		expectedObj := FlowVariable{
			AdditionalProperties: map[string]interface{}{
				"custom-attribute-1": "custom-value-1",
				"custom-attribute-2": "custom-value-2",
			},
			Fields: &FlowVariableFields{
				AdditionalProperties: map[string]interface{}{
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
