package davinci

import (
	"encoding/json"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func ValidFlowsInfoExport(data []byte, cmpOpts ExportCmpOpts) bool {
	if ok := json.Valid(data); !ok {
		return false
	}

	var flowTypeObject FlowsInfo

	if err := Unmarshal([]byte(data), &flowTypeObject, cmpOpts); err != nil {
		return false
	}

	jsonBytes, err := Marshal(flowTypeObject, cmpOpts)
	if err != nil {
		return false
	}

	if string(jsonBytes) == "{}" {
		return false
	}

	if cmp.Equal(flowTypeObject, FlowsInfo{}, cmpopts.EquateEmpty()) {
		return false
	}

	if flowTypeObject.Flow == nil || len(flowTypeObject.Flow) > 1 {
		return false
	}

	for _, flow := range flowTypeObject.Flow {
		if !validateRequiredFlowAttributes(flow, cmpOpts) {
			return false
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
												AdditionalProperties: map[string]interface{}{
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

		if ok := Equal(empty, flowTypeObject, ExportCmpOpts{
			IgnoreConfig:              true,
			IgnoreDesignerCues:        true,
			IgnoreEnvironmentMetadata: true,
			IgnoreFlowMetadata:        true,
			IgnoreUnmappedProperties:  false,
			IgnoreVersionMetadata:     true,
		},
			cmpopts.IgnoreFields(Elements{}, "Nodes"),
		); !ok {
			return false
		}
	}

	// TODO validate required struct attributes
	return true
}

func ValidFlowInfoExport(data []byte, cmpOpts ExportCmpOpts) bool {
	if ok := json.Valid(data); !ok {
		return false
	}

	var flowTypeObject FlowInfo

	if err := Unmarshal([]byte(data), &flowTypeObject, cmpOpts); err != nil {
		return false
	}

	jsonBytes, err := Marshal(flowTypeObject, cmpOpts)
	if err != nil {
		return false
	}

	if string(jsonBytes) == "{}" {
		return false
	}

	if cmp.Equal(flowTypeObject, FlowInfo{}, cmpopts.EquateEmpty()) {
		return false
	}

	if !validateRequiredFlowAttributes(flowTypeObject.Flow, cmpOpts) {
		return false
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
											AdditionalProperties: map[string]interface{}{
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

		if ok := Equal(empty, flowTypeObject, ExportCmpOpts{
			IgnoreConfig:              true,
			IgnoreDesignerCues:        true,
			IgnoreEnvironmentMetadata: true,
			IgnoreUnmappedProperties:  false,
			IgnoreVersionMetadata:     true,
			IgnoreFlowMetadata:        true,
		},
			cmpopts.IgnoreFields(Elements{}, "Nodes"),
		); !ok {
			return false
		}
	}

	// TODO validate required struct attributes
	return true
}

func ValidFlowExport(data []byte, cmpOpts ExportCmpOpts) bool {
	if ok := json.Valid(data); !ok {
		return false
	}

	var flowTypeObject Flow

	if err := Unmarshal([]byte(data), &flowTypeObject, cmpOpts); err != nil {
		return false
	}

	jsonBytes, err := Marshal(flowTypeObject, cmpOpts)
	if err != nil {
		return false
	}

	if string(jsonBytes) == "{}" {
		return false
	}

	if cmp.Equal(flowTypeObject, Flow{}, cmpopts.EquateEmpty()) {
		return false
	}

	if !validateRequiredFlowAttributes(flowTypeObject, cmpOpts) {
		return false
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
										AdditionalProperties: map[string]interface{}{
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

		if ok := Equal(empty, flowTypeObject, ExportCmpOpts{
			IgnoreConfig:              true,
			IgnoreDesignerCues:        true,
			IgnoreEnvironmentMetadata: true,
			IgnoreUnmappedProperties:  false,
			IgnoreVersionMetadata:     true,
			IgnoreFlowMetadata:        true,
		},
			cmpopts.IgnoreFields(Elements{}, "Nodes"),
		); !ok {
			return false
		}
	}

	// TODO validate required struct attributes
	return true
}

func ValidExport(data []byte, cmpOpts ExportCmpOpts) bool {
	return ValidFlowExport(data, cmpOpts) || ValidFlowInfoExport(data, cmpOpts) || ValidFlowsInfoExport(data, cmpOpts)
}

func validateRequiredFlowAttributes(v Flow, opts ExportCmpOpts) bool {

	if !opts.IgnoreConfig && cmp.Equal(v.FlowConfiguration, FlowConfiguration{}, cmpopts.EquateEmpty()) {
		return false
	}

	if !opts.IgnoreFlowMetadata && cmp.Equal(v.FlowMetadata, FlowMetadata{}, cmpopts.EquateEmpty()) {
		return false
	}

	// TODO - anything else to validate?

	return true
}
