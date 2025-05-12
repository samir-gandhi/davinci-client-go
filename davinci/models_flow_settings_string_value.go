package davinci

import (
	"encoding/json"
	"fmt"
)

type FlowSettingsStringValue struct {
	ValueString *string
	ValueObject *interface{}
}

func (o FlowSettingsStringValue) MarshalJSON() ([]byte, error) {
	return o.marshal()
}

func (o *FlowSettingsStringValue) UnmarshalJSON(bytes []byte) (err error) {
	return o.unmarshal(bytes)
}

func (o FlowSettingsStringValue) MarshalDavinci(_ ExportCmpOpts) ([]byte, error) {
	return o.marshal()
}

func (o *FlowSettingsStringValue) UnmarshalDavinci(bytes []byte, _ ExportCmpOpts) (err error) {
	return o.unmarshal(bytes)
}

func (o FlowSettingsStringValue) marshal() ([]byte, error) {
	if o.ValueString != nil {
		return json.Marshal(&o.ValueString)
	}

	if o.ValueObject != nil {
		return json.Marshal(&o.ValueObject)
	}

	return nil, nil // no data in oneOf schemas
}

func (o *FlowSettingsStringValue) unmarshal(bytes []byte) (err error) {

	match := false
	// try to unmarshal data into ValueString
	err = newStrictDecoder(bytes).Decode(&o.ValueString)
	if err == nil {
		jsonValueString, _ := json.Marshal(o.ValueString)
		if string(jsonValueString) == "{}" { // empty struct
			o.ValueString = nil
		} else {
			match = true
		}
	} else {
		o.ValueString = nil
	}

	if !match {
		// try to unmarshal data into ValueObject
		err = newStrictDecoder(bytes).Decode(&o.ValueObject)
		if err == nil {
			match = true
		} else {
			o.ValueObject = nil
		}
	}

	if !match { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(FlowSettingsStringValue)")
	}

	return nil
}
