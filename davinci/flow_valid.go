package davinci

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var (
	ErrInvalidJson               = errors.New("Invalid JSON")
	ErrEmptyFlow                 = errors.New("Flow JSON is empty")
	ErrNoFlowDefinition          = errors.New("No flow definition found in flow export array. Expecting exactly one flow definition")
	ErrMissingSaveVariableValues = errors.New("Save flow variable nodes present but missing variable values")
)

type DiffTypeError struct {
	Diff string
}

type EquatesEmptyTypeError DiffTypeError

func (e *EquatesEmptyTypeError) Error() string {
	return fmt.Sprintf("Flow has been evaluated to be empty. Invalid type? Diff: %s", e.Diff)
}

type MissingRequiredFlowFieldsTypeError DiffTypeError

func (e *MissingRequiredFlowFieldsTypeError) Error() string {
	return fmt.Sprintf("Flow has missing fields. Diff: %s", e.Diff)
}

type UnknownAdditionalFieldsTypeError DiffTypeError

func (e *UnknownAdditionalFieldsTypeError) Error() string {
	return fmt.Sprintf("Flow has unknown additional fields. Diff: %s", e.Diff)
}

type MinFlowDefinitionsExceededTypeError struct {
	Min int
}

func (e *MinFlowDefinitionsExceededTypeError) Error() string {
	return fmt.Sprintf("There are not enough flows exported in the flow group.  Expecting a minimum of %d", e.Min)
}

type MaxFlowDefinitionsExceededTypeError struct {
	Max int
}

func (e *MaxFlowDefinitionsExceededTypeError) Error() string {
	return fmt.Sprintf("There are too many flows exported in the flow group.  Expecting a maximum of %d", e.Max)
}

func ValidFlowsInfoExport(data []byte, cmpOpts ExportCmpOpts) (err error) {
	if ok := json.Valid(data); !ok {
		return ErrInvalidJson
	}

	var flowTypeObject FlowsInfo

	if err := Unmarshal([]byte(data), &flowTypeObject, cmpOpts); err != nil {
		return err
	}

	jsonBytes, err := Marshal(flowTypeObject, cmpOpts)
	if err != nil {
		return err
	}

	if string(jsonBytes) == "{}" {
		return ErrEmptyFlow
	}

	if cmp.Equal(flowTypeObject, FlowsInfo{}, cmpopts.EquateEmpty()) {
		diff := cmp.Diff(flowTypeObject, FlowsInfo{}, cmpopts.EquateEmpty())
		return &EquatesEmptyTypeError{
			Diff: diff,
		}
	}

	if flowTypeObject.Flow == nil {
		return ErrNoFlowDefinition
	}

	if cmpOpts.MaxFlows != nil && len(flowTypeObject.Flow) > *cmpOpts.MaxFlows {
		return &MaxFlowDefinitionsExceededTypeError{
			Max: *cmpOpts.MaxFlows,
		}
	}

	if cmpOpts.MinFlows != nil && len(flowTypeObject.Flow) < *cmpOpts.MinFlows {
		return &MinFlowDefinitionsExceededTypeError{
			Min: *cmpOpts.MinFlows,
		}
	}

	for _, flow := range flowTypeObject.Flow {
		if err := validateRequiredFlowAttributes(flow, cmpOpts); err != nil {
			return err
		}
		if err := validateVariableValues(flow, cmpOpts); err != nil {
			return err
		}
	}

	if !cmpOpts.IgnoreUnmappedProperties {
		empty := FlowsInfo{
			Flow: []Flow{
				{
					FlowConfiguration: FlowConfiguration{
						FlowUpdateConfiguration: FlowUpdateConfiguration{
							GraphData: &GraphData{
								Elements: &Elements{
									Nodes: []Node{
										{
											Data: &NodeData{
												AdditionalProperties: map[string]any{
													"test1": "test1", // to overcome odd behaviours with cmpopts
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}

		cmpOpts := ExportCmpOpts{
			IgnoreConfig:              true,
			IgnoreDesignerCues:        true,
			IgnoreEnvironmentMetadata: true,
			IgnoreFlowMetadata:        true,
			IgnoreUnmappedProperties:  false,
			IgnoreVersionMetadata:     true,
			IgnoreFlowVariables:       true,
		}

		opts := cmpopts.IgnoreFields(Elements{}, "Nodes")

		if ok := Equal(empty, flowTypeObject, cmpOpts, opts); !ok {
			diff := Diff(empty, flowTypeObject, cmpOpts, opts)
			return &UnknownAdditionalFieldsTypeError{
				Diff: diff,
			}
		}
	}

	// TODO validate required struct attributes
	return nil
}

func ValidFlowInfoExport(data []byte, cmpOpts ExportCmpOpts) (err error) {
	if ok := json.Valid(data); !ok {
		return ErrInvalidJson
	}

	var flowTypeObject FlowInfo

	if err := Unmarshal([]byte(data), &flowTypeObject, cmpOpts); err != nil {
		return err
	}

	jsonBytes, err := Marshal(flowTypeObject, cmpOpts)
	if err != nil {
		return err
	}

	if string(jsonBytes) == "{}" {
		return ErrEmptyFlow
	}

	if cmp.Equal(flowTypeObject, FlowInfo{}, cmpopts.EquateEmpty()) {
		diff := cmp.Diff(flowTypeObject, FlowInfo{}, cmpopts.EquateEmpty())
		return &EquatesEmptyTypeError{
			Diff: diff,
		}
	}

	if err := validateRequiredFlowAttributes(flowTypeObject.Flow, cmpOpts); err != nil {
		return err
	}

	if err := validateVariableValues(flowTypeObject.Flow, cmpOpts); err != nil {
		return err
	}

	if !cmpOpts.IgnoreUnmappedProperties {
		empty := FlowInfo{
			Flow: Flow{
				FlowConfiguration: FlowConfiguration{
					FlowUpdateConfiguration: FlowUpdateConfiguration{
						GraphData: &GraphData{
							Elements: &Elements{
								Nodes: []Node{
									{
										Data: &NodeData{
											AdditionalProperties: map[string]any{
												"test1": "test1", // to overcome odd behaviours with cmpopts
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}

		cmpOpts := ExportCmpOpts{
			IgnoreConfig:              true,
			IgnoreDesignerCues:        true,
			IgnoreEnvironmentMetadata: true,
			IgnoreUnmappedProperties:  false,
			IgnoreVersionMetadata:     true,
			IgnoreFlowMetadata:        true,
			IgnoreFlowVariables:       true,
		}

		opts := cmpopts.IgnoreFields(Elements{}, "Nodes")

		if ok := Equal(empty, flowTypeObject, cmpOpts, opts); !ok {
			diff := Diff(empty, flowTypeObject, cmpOpts, opts)
			return &UnknownAdditionalFieldsTypeError{
				Diff: diff,
			}
		}
	}

	// TODO validate required struct attributes
	return nil
}

func ValidFlowsExport(data []byte, cmpOpts ExportCmpOpts) (err error) {
	if ok := json.Valid(data); !ok {
		return ErrInvalidJson
	}

	var flowTypeObject Flows

	if err := Unmarshal([]byte(data), &flowTypeObject, cmpOpts); err != nil {
		return err
	}

	jsonBytes, err := Marshal(flowTypeObject, cmpOpts)
	if err != nil {
		return err
	}

	if string(jsonBytes) == "{}" {
		return ErrEmptyFlow
	}

	if cmp.Equal(flowTypeObject, Flows{}, cmpopts.EquateEmpty()) {
		diff := cmp.Diff(flowTypeObject, Flows{}, cmpopts.EquateEmpty())
		return &EquatesEmptyTypeError{
			Diff: diff,
		}
	}

	if flowTypeObject.Flow == nil {
		return ErrNoFlowDefinition
	}

	if cmpOpts.MaxFlows != nil && len(flowTypeObject.Flow) > *cmpOpts.MaxFlows {
		return &MaxFlowDefinitionsExceededTypeError{
			Max: *cmpOpts.MaxFlows,
		}
	}

	if cmpOpts.MinFlows != nil && len(flowTypeObject.Flow) < *cmpOpts.MinFlows {
		return &MinFlowDefinitionsExceededTypeError{
			Min: *cmpOpts.MinFlows,
		}
	}

	for _, flow := range flowTypeObject.Flow {
		if err := validateRequiredFlowAttributes(flow, cmpOpts); err != nil {
			return err
		}
		if err := validateVariableValues(flow, cmpOpts); err != nil {
			return err
		}
	}

	if !cmpOpts.IgnoreUnmappedProperties {
		empty := Flows{
			Flow: []Flow{
				{
					FlowConfiguration: FlowConfiguration{
						FlowUpdateConfiguration: FlowUpdateConfiguration{
							GraphData: &GraphData{
								Elements: &Elements{
									Nodes: []Node{
										{
											Data: &NodeData{
												AdditionalProperties: map[string]any{
													"test1": "test1", // to overcome odd behaviours with cmpopts
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}

		cmpOpts := ExportCmpOpts{
			IgnoreConfig:              true,
			IgnoreDesignerCues:        true,
			IgnoreEnvironmentMetadata: true,
			IgnoreFlowMetadata:        true,
			IgnoreUnmappedProperties:  false,
			IgnoreVersionMetadata:     true,
			IgnoreFlowVariables:       true,
		}

		opts := cmpopts.IgnoreFields(Elements{}, "Nodes")

		if ok := Equal(empty, flowTypeObject, cmpOpts, opts); !ok {
			diff := Diff(empty, flowTypeObject, cmpOpts, opts)
			return &UnknownAdditionalFieldsTypeError{
				Diff: diff,
			}
		}
	}

	// TODO validate required struct attributes
	return nil
}

func ValidFlowExport(data []byte, cmpOpts ExportCmpOpts) (err error) {
	if ok := json.Valid(data); !ok {
		return ErrInvalidJson
	}

	var flowTypeObject Flow

	if err := Unmarshal([]byte(data), &flowTypeObject, cmpOpts); err != nil {
		return err
	}

	jsonBytes, err := Marshal(flowTypeObject, cmpOpts)
	if err != nil {
		return err
	}

	if string(jsonBytes) == "{}" {
		return ErrEmptyFlow
	}

	if cmp.Equal(flowTypeObject, Flow{}, cmpopts.EquateEmpty()) {
		diff := cmp.Diff(flowTypeObject, Flow{}, cmpopts.EquateEmpty())
		return &EquatesEmptyTypeError{
			Diff: diff,
		}
	}

	if err := validateRequiredFlowAttributes(flowTypeObject, cmpOpts); err != nil {
		return err
	}

	if err := validateVariableValues(flowTypeObject, cmpOpts); err != nil {
		return err
	}

	if !cmpOpts.IgnoreUnmappedProperties {
		empty := Flow{
			FlowConfiguration: FlowConfiguration{
				FlowUpdateConfiguration: FlowUpdateConfiguration{
					GraphData: &GraphData{
						Elements: &Elements{
							Nodes: []Node{
								{
									Data: &NodeData{
										AdditionalProperties: map[string]any{
											"test1": "test1", // to overcome odd behaviours with cmpopts
										},
									},
								},
							},
						},
					},
				},
			},
		}

		cmpOpts := ExportCmpOpts{
			IgnoreConfig:              true,
			IgnoreDesignerCues:        true,
			IgnoreEnvironmentMetadata: true,
			IgnoreUnmappedProperties:  false,
			IgnoreVersionMetadata:     true,
			IgnoreFlowMetadata:        true,
			IgnoreFlowVariables:       true,
		}

		opts := cmpopts.IgnoreFields(Elements{}, "Nodes")

		if ok := Equal(empty, flowTypeObject, cmpOpts, opts); !ok {
			diff := Diff(empty, flowTypeObject, cmpOpts, opts)
			return &UnknownAdditionalFieldsTypeError{
				Diff: diff,
			}
		}
	}

	// TODO validate required struct attributes
	return nil
}

func ValidExport(data []byte, cmpOpts ExportCmpOpts) (err error) {

	if err := ValidFlowExport(data, cmpOpts); err != nil {
		return err
	}

	if err := ValidFlowInfoExport(data, cmpOpts); err != nil {
		return err
	}

	if err := ValidFlowsInfoExport(data, cmpOpts); err != nil {
		return err
	}

	return nil
}

func validateRequiredFlowAttributes(v Flow, opts ExportCmpOpts) error {

	if !opts.IgnoreConfig && cmp.Equal(v.FlowConfiguration, FlowConfiguration{}, cmpopts.EquateEmpty()) {
		diff := cmp.Diff(v.FlowConfiguration, FlowConfiguration{}, cmpopts.EquateEmpty())
		return &MissingRequiredFlowFieldsTypeError{
			Diff: diff,
		}
	}

	if !opts.IgnoreFlowMetadata && cmp.Equal(v.FlowMetadata, FlowMetadata{}, cmpopts.EquateEmpty()) {
		diff := cmp.Diff(v.FlowMetadata, FlowMetadata{}, cmpopts.EquateEmpty())
		return &MissingRequiredFlowFieldsTypeError{
			Diff: diff,
		}
	}

	// TODO - anything else to validate?

	return nil
}

func validateVariableValues(v Flow, opts ExportCmpOpts) error {

	if opts.NodeOpts != nil && opts.NodeOpts.VariablesConnector != nil && opts.NodeOpts.VariablesConnector.ExpectVariableValues {

		// If there are no variable nodes, return true
		if v.FlowConfiguration.GraphData == nil || v.FlowConfiguration.GraphData.Elements == nil || v.FlowConfiguration.GraphData.Elements.Nodes == nil || len(v.FlowConfiguration.GraphData.Elements.Nodes) == 0 {
			return nil
		}

		saveVariableNodePresent := false

		// Check variables in the variables connectors
		for _, node := range v.FlowConfiguration.GraphData.Elements.Nodes {
			if node.Data != nil && node.Data.Properties != nil {

				if node.Data.Properties.SaveFlowVariables != nil && node.Data.Properties.SaveFlowVariables.Value != nil && len(node.Data.Properties.SaveFlowVariables.Value) > 0 {
					saveVariableNodePresent = true
					for _, variable := range node.Data.Properties.SaveFlowVariables.Value {
						if variable.Value != nil {
							return nil
						}
					}
				}

				if node.Data.Properties.SaveVariables != nil && node.Data.Properties.SaveVariables.Value != nil && len(node.Data.Properties.SaveVariables.Value) > 0 {
					saveVariableNodePresent = true
					for _, variable := range node.Data.Properties.SaveVariables.Value {
						if variable.Value != nil {
							return nil
						}
					}
				}

			}
		}

		if saveVariableNodePresent {
			return ErrMissingSaveVariableValues
		}
	}

	return nil
}
