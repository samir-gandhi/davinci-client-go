package davinci

import (
	"encoding/json"
	"maps"
)

type _GraphData GraphData
type GraphData struct {
	AdditionalProperties map[string]any `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	BoxSelectionEnabled  *bool          `json:"boxSelectionEnabled,omitempty" davinci:"boxSelectionEnabled,designercue,omitempty"`
	Data                 *Data          `json:"data,omitempty" davinci:"data,*,omitempty"`
	Elements             *Elements      `json:"elements,omitempty" davinci:"elements,*,omitempty"`
	MaxZoom              *float64       `json:"maxZoom,omitempty" davinci:"maxZoom,designercue,omitempty"`
	MinZoom              *float64       `json:"minZoom,omitempty" davinci:"minZoom,designercue,omitempty"`
	Pan                  *Pan           `json:"pan,omitempty" davinci:"pan,*,omitempty"`
	PanningEnabled       *bool          `json:"panningEnabled,omitempty" davinci:"panningEnabled,designercue,omitempty"`
	Renderer             *Renderer      `json:"renderer,omitempty" davinci:"renderer,*,omitempty"`
	UserPanningEnabled   *bool          `json:"userPanningEnabled,omitempty" davinci:"userPanningEnabled,designercue,omitempty"`
	UserZoomingEnabled   *bool          `json:"userZoomingEnabled,omitempty" davinci:"userZoomingEnabled,designercue,omitempty"`
	Zoom                 *int32         `json:"zoom,omitempty" davinci:"zoom,designercue,omitempty"`
	ZoomingEnabled       *bool          `json:"zoomingEnabled,omitempty" davinci:"zoomingEnabled,designercue,omitempty"`
}

func (o GraphData) MarshalJSON() ([]byte, error) {
	result, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(result)
}

func (o GraphData) ToMap() (map[string]any, error) {

	result := map[string]any{}

	if o.BoxSelectionEnabled != nil {
		result["boxSelectionEnabled"] = o.BoxSelectionEnabled
	}

	if o.Data != nil {
		result["data"] = o.Data
	}

	if o.Elements != nil {
		result["elements"] = o.Elements
	}

	if o.MaxZoom != nil {
		result["maxZoom"] = o.MaxZoom
	}

	if o.MinZoom != nil {
		result["minZoom"] = o.MinZoom
	}

	if o.Pan != nil {
		result["pan"] = o.Pan
	}

	if o.PanningEnabled != nil {
		result["panningEnabled"] = o.PanningEnabled
	}

	if o.Renderer != nil {
		result["renderer"] = o.Renderer
	}

	if o.UserPanningEnabled != nil {
		result["userPanningEnabled"] = o.UserPanningEnabled
	}

	if o.UserZoomingEnabled != nil {
		result["userZoomingEnabled"] = o.UserZoomingEnabled
	}

	if o.Zoom != nil {
		result["zoom"] = o.Zoom
	}

	if o.ZoomingEnabled != nil {
		result["zoomingEnabled"] = o.ZoomingEnabled
	}

	maps.Copy(result, o.AdditionalProperties)

	return result, nil
}

func (o *GraphData) UnmarshalJSON(bytes []byte) (err error) {
	varGraphData := _GraphData{}

	if err = json.Unmarshal(bytes, &varGraphData); err == nil {
		*o = GraphData(varGraphData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "boxSelectionEnabled")
		delete(additionalProperties, "data")
		delete(additionalProperties, "elements")
		delete(additionalProperties, "maxZoom")
		delete(additionalProperties, "minZoom")
		delete(additionalProperties, "pan")
		delete(additionalProperties, "panningEnabled")
		delete(additionalProperties, "renderer")
		delete(additionalProperties, "userPanningEnabled")
		delete(additionalProperties, "userZoomingEnabled")
		delete(additionalProperties, "zoom")
		delete(additionalProperties, "zoomingEnabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}
