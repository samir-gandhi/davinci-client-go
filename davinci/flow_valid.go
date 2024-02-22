package davinci

import (
	"encoding/json"
	"log"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func ValidFlowsInfoJSON(data []byte, cmpOpts ExportCmpOpts) bool {
	if ok := json.Valid(data); !ok {
		return false
	}

	var flowTypeObject FlowsInfo

	if err := json.Unmarshal([]byte(data), &flowTypeObject); err != nil {
		return false
	}

	jsonBytes, err := json.Marshal(flowTypeObject)
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
		empty := FlowsInfo{}

		if ok := FlowsInfoEqual(empty, flowTypeObject, ExportCmpOpts{
			IgnoreConfig:              true,
			IgnoreDesignerCues:        true,
			IgnoreEnvironmentMetadata: true,
			IgnoreFlowMetadata:        true,
			IgnoreUnmappedProperties:  false,
			IgnoreVersionMetadata:     true,
		}); !ok {
			return false
		}
	}

	// TODO validate required struct attributes
	return true
}

func ValidFlowInfoJSON(data []byte, cmpOpts ExportCmpOpts) bool {
	if ok := json.Valid(data); !ok {
		return false
	}

	var flowTypeObject FlowInfo

	if err := json.Unmarshal([]byte(data), &flowTypeObject); err != nil {
		return false
	}

	jsonBytes, err := json.Marshal(flowTypeObject)
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
		empty := FlowInfo{}

		if ok := FlowInfoEqual(empty, flowTypeObject, ExportCmpOpts{
			IgnoreConfig:              true,
			IgnoreDesignerCues:        true,
			IgnoreEnvironmentMetadata: true,
			IgnoreUnmappedProperties:  false,
			IgnoreVersionMetadata:     true,
			IgnoreFlowMetadata:        true,
		}); !ok {
			return false
		}
	}

	// TODO validate required struct attributes
	return true
}

func ValidFlowJSON(data []byte, cmpOpts ExportCmpOpts) bool {
	if ok := json.Valid(data); !ok {
		log.Printf("HERE!!!!!1")
		return false
	}

	var flowTypeObject Flow

	if err := json.Unmarshal([]byte(data), &flowTypeObject); err != nil {
		log.Printf("HERE!!!!!2")
		return false
	}

	jsonBytes, err := json.Marshal(flowTypeObject)
	if err != nil {
		log.Printf("HERE!!!!!3")
		return false
	}

	if string(jsonBytes) == "{}" {
		log.Printf("HERE!!!!!4")
		return false
	}

	if cmp.Equal(flowTypeObject, Flow{}, cmpopts.EquateEmpty()) {
		log.Printf("HERE!!!!!5")
		return false
	}

	if !validateRequiredFlowAttributes(flowTypeObject, cmpOpts) {
		log.Printf("HERE!!!!!6")
		return false
	}

	if !cmpOpts.IgnoreUnmappedProperties {
		empty := Flow{}

		if ok := FlowEqual(empty, flowTypeObject, ExportCmpOpts{
			IgnoreConfig:              true,
			IgnoreDesignerCues:        true,
			IgnoreEnvironmentMetadata: true,
			IgnoreUnmappedProperties:  false,
			IgnoreVersionMetadata:     true,
			IgnoreFlowMetadata:        true,
		}); !ok {
			log.Printf("HERE!!!!!7")
			return false
		}
	}

	// TODO validate required struct attributes
	return true
}

func ValidJSON(data []byte, cmpOpts ExportCmpOpts) bool {
	return ValidFlowJSON(data, cmpOpts) || ValidFlowInfoJSON(data, cmpOpts) || ValidFlowsInfoJSON(data, cmpOpts)
}

func validateRequiredFlowAttributes(v Flow, opts ExportCmpOpts) bool {

	if !opts.IgnoreConfig && cmp.Equal(v.FlowConfiguration, FlowConfiguration{}, cmpopts.EquateEmpty()) {
		log.Printf("HERE!!!!!x.1")
		return false
	}

	if !opts.IgnoreFlowMetadata && cmp.Equal(v.FlowMetadata, FlowMetadata{}, cmpopts.EquateEmpty()) {
		log.Printf("HERE!!!!!x.2")
		return false
	}

	// TODO - anything else to validate?

	return true
}
