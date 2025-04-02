package test

import "github.com/samir-gandhi/davinci-client-go/davinci"

func Data_FullBasic() davinci.Flow {

	return davinci.Flow{
		AdditionalProperties: map[string]any{
			"custom-attribute-1": "custom-attribute-1-value",
			"custom-attribute-2": "custom-attribute-2-value",
		},
		FlowConfiguration: davinci.FlowConfiguration{
			FlowUpdateConfiguration: davinci.FlowUpdateConfiguration{
				GraphData: &davinci.GraphData{
					AdditionalProperties: map[string]any{
						"custom-attribute-1": "custom-attribute-1-value",
						"custom-attribute-2": "custom-attribute-2-value",
					},
					Elements: &davinci.Elements{
						AdditionalProperties: map[string]any{
							"custom-attribute-1": "custom-attribute-1-value",
							"custom-attribute-2": "custom-attribute-2-value",
						},
						Nodes: []davinci.Node{
							{
								AdditionalProperties: map[string]any{
									"custom-attribute-1": "custom-attribute-1-value",
									"custom-attribute-2": "custom-attribute-2-value",
								},
								Data: &davinci.NodeData{
									AdditionalProperties: map[string]any{
										"custom-attribute-1": "custom-attribute-1-value",
										"custom-attribute-2": "custom-attribute-2-value",
									},
									ID: func() *string {
										v := "1u2m5vzr49"
										return &v
									}(),
									NodeType: func() *string {
										v := "CONNECTION"
										return &v
									}(),
									ConnectionID: func() *string {
										v := "867ed4363b2bc21c860085ad2baa817d"
										return &v
									}(),
									ConnectorID: func() *string {
										v := "httpConnector"
										return &v
									}(),
									Name: func() *string {
										v := "Http"
										return &v
									}(),
									Label: func() *string {
										v := "Http"
										return &v
									}(),
									Status: func() *string {
										v := "configured"
										return &v
									}(),
									CapabilityName: func() *string {
										v := "customHtmlMessage"
										return &v
									}(),
									Type: func() *string {
										v := "trigger"
										return &v
									}(),
									Properties: func() *davinci.Properties {
										return &davinci.Properties{
											AdditionalProperties: map[string]any{
												"custom-attribute-1": "custom-attribute-1-value",
												"custom-attribute-2": "custom-attribute-2-value",
												"message": map[string]any{
													"custom-attribute-1": "custom-attribute-1-value",
													"custom-attribute-2": "custom-attribute-2-value",
													"value":              "[\n  {\n    \"children\": [\n      {\n        \"text\": \"Hello, world?\"\n      }\n    ]\n  }\n]",
												},
											},
										}
									}(),
								},
								Position: &davinci.Position{
									AdditionalProperties: map[string]any{
										"custom-attribute-1": "custom-attribute-1-value",
										"custom-attribute-2": "custom-attribute-2-value",
									},
									X: func() *float64 {
										v := float64(277)
										return &v
									}(),
									Y: func() *float64 {
										v := float64(336)
										return &v
									}(),
								},
								Group:      "nodes",
								Removed:    false,
								Selected:   false,
								Selectable: true,
								Locked:     false,
								Grabbable:  true,
								Pannable:   false,
								Classes:    "",
							},
							{
								// AdditionalProperties: map[string]any{},
								Data: &davinci.NodeData{
									// AdditionalProperties: map[string]any{},
									ID: func() *string {
										v := "8fvg7tfr8j"
										return &v
									}(),
									NodeType: func() *string {
										v := "EVAL"
										return &v
									}(),
									Label: func() *string {
										v := "Evaluator"
										return &v
									}(),
								},
								Position: &davinci.Position{
									// AdditionalProperties: map[string]any{},
									X: func() *float64 {
										v := float64(426.5)
										return &v
									}(),
									Y: func() *float64 {
										v := float64(337.25)
										return &v
									}(),
								},
								Group:      "nodes",
								Removed:    false,
								Selected:   false,
								Selectable: true,
								Locked:     false,
								Grabbable:  true,
								Pannable:   false,
								Classes:    "",
							},
							{
								// AdditionalProperties: map[string]any{},
								Data: &davinci.NodeData{
									// AdditionalProperties: map[string]any{},
									ID: func() *string {
										v := "nx0o1b2cmw"
										return &v
									}(),
									NodeType: func() *string {
										v := "CONNECTION"
										return &v
									}(),
									ConnectionID: func() *string {
										v := "de650ca45593b82c49064ead10b9fe17"
										return &v
									}(),
									ConnectorID: func() *string {
										v := "functionsConnector"
										return &v
									}(),
									Name: func() *string {
										v := "Functions"
										return &v
									}(),
									Label: func() *string {
										v := "Functions"
										return &v
									}(),
									Status: func() *string {
										v := "configured"
										return &v
									}(),
									CapabilityName: func() *string {
										v := "AEqualsB"
										return &v
									}(),
									Type: func() *string {
										v := "trigger"
										return &v
									}(),
									Properties: func() *davinci.Properties {
										return &davinci.Properties{
											AdditionalProperties: map[string]any{
												"leftValueA": map[string]any{
													"value": "[\n  {\n    \"children\": [\n      {\n        \"text\": \"1\"\n      }\n    ]\n  }\n]",
												},
												"rightValueB": map[string]any{
													"value": "[\n  {\n    \"children\": [\n      {\n        \"text\": \"1\"\n      }\n    ]\n  }\n]",
												},
											},
										}
									}(),
								},
								Position: &davinci.Position{
									// AdditionalProperties: map[string]any{},
									X: func() *float64 {
										v := float64(576)
										return &v
									}(),
									Y: func() *float64 {
										v := float64(338.5)
										return &v
									}(),
								},
								Group:      "nodes",
								Removed:    false,
								Selected:   false,
								Selectable: true,
								Locked:     false,
								Grabbable:  true,
								Pannable:   false,
								Classes:    "",
							},
							{
								// AdditionalProperties: map[string]any{},
								Data: &davinci.NodeData{
									// AdditionalProperties: map[string]any{},
									ID: func() *string {
										v := "cdcw8k7dnx"
										return &v
									}(),
									NodeType: func() *string {
										v := "EVAL"
										return &v
									}(),
									Label: func() *string {
										v := "Evaluator"
										return &v
									}(),
									Properties: func() *davinci.Properties {
										return &davinci.Properties{
											AdditionalProperties: map[string]any{
												"vsp1ewtr9m": map[string]any{
													"value": "allTriggersFalse",
												},
												"xb74p6rkd8": map[string]any{
													"value": "anyTriggersFalse",
												},
											},
										}
									}(),
								},
								Position: &davinci.Position{
									// AdditionalProperties: map[string]any{},
									X: func() *float64 {
										v := float64(717)
										return &v
									}(),
									Y: func() *float64 {
										v := float64(326)
										return &v
									}(),
								},
								Group:      "nodes",
								Removed:    false,
								Selected:   false,
								Selectable: true,
								Locked:     false,
								Grabbable:  true,
								Pannable:   false,
								Classes:    "",
							},
							{
								// AdditionalProperties: map[string]any{},
								Data: &davinci.NodeData{
									// AdditionalProperties: map[string]any{},
									ID: func() *string {
										v := "ikt13crnhy"
										return &v
									}(),
									NodeType: func() *string {
										v := "CONNECTION"
										return &v
									}(),
									ConnectionID: func() *string {
										v := "867ed4363b2bc21c860085ad2baa817d"
										return &v
									}(),
									ConnectorID: func() *string {
										v := "httpConnector"
										return &v
									}(),
									Name: func() *string {
										v := "Http"
										return &v
									}(),
									Label: func() *string {
										v := "Http"
										return &v
									}(),
									Status: func() *string {
										v := "configured"
										return &v
									}(),
									CapabilityName: func() *string {
										v := "createSuccessResponse"
										return &v
									}(),
									Type: func() *string {
										v := "action"
										return &v
									}(),
									Properties: func() *davinci.Properties {
										return &davinci.Properties{
											// AdditionalProperties: map[string]any{},
										}
									}(),
								},
								Position: &davinci.Position{
									// AdditionalProperties: map[string]any{},
									X: func() *float64 {
										v := float64(1197)
										return &v
									}(),
									Y: func() *float64 {
										v := float64(266)
										return &v
									}(),
								},
								Group:      "nodes",
								Removed:    false,
								Selected:   false,
								Selectable: true,
								Locked:     false,
								Grabbable:  true,
								Pannable:   false,
								Classes:    "",
							},
							{
								// AdditionalProperties: map[string]any{},
								Data: &davinci.NodeData{
									// AdditionalProperties: map[string]any{},
									ID: func() *string {
										v := "vsp1ewtr9m"
										return &v
									}(),
									NodeType: func() *string {
										v := "CONNECTION"
										return &v
									}(),
									ConnectionID: func() *string {
										v := "53ab83a4a4ab919d9f2cb02d9e111ac8"
										return &v
									}(),
									ConnectorID: func() *string {
										v := "errorConnector"
										return &v
									}(),
									Name: func() *string {
										v := "Error Message"
										return &v
									}(),
									Label: func() *string {
										v := "Error Message"
										return &v
									}(),
									Status: func() *string {
										v := "configured"
										return &v
									}(),
									CapabilityName: func() *string {
										v := "customErrorMessage"
										return &v
									}(),
									Type: func() *string {
										v := "action"
										return &v
									}(),
									Properties: func() *davinci.Properties {
										return &davinci.Properties{
											AdditionalProperties: map[string]any{
												"errorMessage": map[string]any{
													"value": "[\n  {\n    \"children\": [\n      {\n        \"text\": \"Error\"\n      }\n    ]\n  }\n]",
												},
											},
										}
									}(),
								},
								Position: &davinci.Position{
									// AdditionalProperties: map[string]any{},
									X: func() *float64 {
										v := float64(1197)
										return &v
									}(),
									Y: func() *float64 {
										v := float64(416)
										return &v
									}(),
								},
								Group:      "nodes",
								Removed:    false,
								Selected:   false,
								Selectable: true,
								Locked:     false,
								Grabbable:  true,
								Pannable:   false,
								Classes:    "",
							},
							{
								// AdditionalProperties: map[string]any{},
								Data: &davinci.NodeData{
									// AdditionalProperties: map[string]any{},
									ID: func() *string {
										v := "xb74p6rkd8"
										return &v
									}(),
									NodeType: func() *string {
										v := "CONNECTION"
										return &v
									}(),
									ConnectionID: func() *string {
										v := "33329a264e268ab31fb19637debf1ea3"
										return &v
									}(),
									ConnectorID: func() *string {
										v := "flowConnector"
										return &v
									}(),
									Name: func() *string {
										v := "Flow Conductor"
										return &v
									}(),
									Label: func() *string {
										v := "Flow Conductor"
										return &v
									}(),
									Status: func() *string {
										v := "configured"
										return &v
									}(),
									CapabilityName: func() *string {
										v := "startUiSubFlow"
										return &v
									}(),
									Type: func() *string {
										v := "trigger"
										return &v
									}(),
									Properties: func() *davinci.Properties {
										return &davinci.Properties{
											// AdditionalProperties: map[string]any{},
											SubFlowID: func() *davinci.SubFlowID {
												return &davinci.SubFlowID{
													Value: &davinci.SubFlowValue{
														Label: func() *string {
															v := "subflow 2"
															return &v
														}(),
														Value: func() *string {
															v := "07503fed5c02849dbbd5ee932da654b2"
															return &v
														}(),
													},
												}
											}(),
											SubFlowVersionID: func() *davinci.SubFlowVersionID {
												return &davinci.SubFlowVersionID{
													Value: &davinci.SubFlowVersionIDValue{
														ValueInt: func() *int32 {
															v := int32(-1)
															return &v
														}(),
													},
												}
											}(),
										}
									}(),
								},
								Position: &davinci.Position{
									// AdditionalProperties: map[string]any{},
									X: func() *float64 {
										v := float64(867)
										return &v
									}(),
									Y: func() *float64 {
										v := float64(446)
										return &v
									}(),
								},
								Group:      "nodes",
								Removed:    false,
								Selected:   false,
								Selectable: true,
								Locked:     false,
								Grabbable:  true,
								Pannable:   false,
								Classes:    "",
							},
							{
								// AdditionalProperties: map[string]any{},
								Data: &davinci.NodeData{
									// AdditionalProperties: map[string]any{},
									ID: func() *string {
										v := "kq5ybvwvro"
										return &v
									}(),
									NodeType: func() *string {
										v := "CONNECTION"
										return &v
									}(),
									ConnectionID: func() *string {
										v := "33329a264e268ab31fb19637debf1ea3"
										return &v
									}(),
									ConnectorID: func() *string {
										v := "flowConnector"
										return &v
									}(),
									Name: func() *string {
										v := "Flow Conductor"
										return &v
									}(),
									Label: func() *string {
										v := "Flow Conductor"
										return &v
									}(),
									Status: func() *string {
										v := "configured"
										return &v
									}(),
									CapabilityName: func() *string {
										v := "startUiSubFlow"
										return &v
									}(),
									Type: func() *string {
										v := "trigger"
										return &v
									}(),
									Properties: func() *davinci.Properties {
										return &davinci.Properties{
											// AdditionalProperties: map[string]any{},
											SubFlowID: func() *davinci.SubFlowID {
												return &davinci.SubFlowID{
													Value: &davinci.SubFlowValue{
														Label: func() *string {
															v := "subflow 1"
															return &v
														}(),
														Value: func() *string {
															v := "00f66e8926ced6ef5b83619fde4a314a"
															return &v
														}(),
													},
												}
											}(),
											SubFlowVersionID: func() *davinci.SubFlowVersionID {
												return &davinci.SubFlowVersionID{
													Value: &davinci.SubFlowVersionIDValue{
														ValueInt: func() *int32 {
															v := int32(-1)
															return &v
														}(),
													},
												}
											}(),
										}
									}(),
								},
								Position: &davinci.Position{
									// AdditionalProperties: map[string]any{},
									X: func() *float64 {
										v := float64(867)
										return &v
									}(),
									Y: func() *float64 {
										v := float64(236)
										return &v
									}(),
								},
								Group:      "nodes",
								Removed:    false,
								Selected:   false,
								Selectable: true,
								Locked:     false,
								Grabbable:  true,
								Pannable:   false,
								Classes:    "",
							},
							{
								// AdditionalProperties: map[string]any{},
								Data: &davinci.NodeData{
									// AdditionalProperties: map[string]any{},
									ID: func() *string {
										v := "j74pmg6577"
										return &v
									}(),
									NodeType: func() *string {
										v := "EVAL"
										return &v
									}(),
								},
								Position: &davinci.Position{
									// AdditionalProperties: map[string]any{},
									X: func() *float64 {
										v := float64(1017)
										return &v
									}(),
									Y: func() *float64 {
										v := float64(236)
										return &v
									}(),
								},
								Group:      "nodes",
								Removed:    false,
								Selected:   false,
								Selectable: true,
								Locked:     false,
								Grabbable:  true,
								Pannable:   false,
								Classes:    "",
							},
							{
								// AdditionalProperties: map[string]any{},
								Data: &davinci.NodeData{
									// AdditionalProperties: map[string]any{},
									ID: func() *string {
										v := "pensvkew7y"
										return &v
									}(),
									NodeType: func() *string {
										v := "EVAL"
										return &v
									}(),
									Properties: nil,
								},
								Position: &davinci.Position{
									// AdditionalProperties: map[string]any{},
									X: func() *float64 {
										v := float64(1032)
										return &v
									}(),
									Y: func() *float64 {
										v := float64(431)
										return &v
									}(),
								},
								Group:      "nodes",
								Removed:    false,
								Selected:   false,
								Selectable: true,
								Locked:     false,
								Grabbable:  true,
								Pannable:   false,
								Classes:    "",
							},
							{
								// AdditionalProperties: map[string]any{},
								Data: &davinci.NodeData{
									// AdditionalProperties: map[string]any{},
									ID: func() *string {
										v := "3zvjdgdljx"
										return &v
									}(),
									NodeType: func() *string {
										v := "CONNECTION"
										return &v
									}(),
									ConnectionID: func() *string {
										v := "06922a684039827499bdbdd97f49827b"
										return &v
									}(),
									ConnectorID: func() *string {
										v := "variablesConnector"
										return &v
									}(),
									Name: func() *string {
										v := "Variables"
										return &v
									}(),
									Label: func() *string {
										v := "Variables"
										return &v
									}(),
									Status: func() *string {
										v := "configured"
										return &v
									}(),
									CapabilityName: func() *string {
										v := "saveFlowValue"
										return &v
									}(),
									Type: func() *string {
										v := "trigger"
										return &v
									}(),
									Properties: func() *davinci.Properties {
										return &davinci.Properties{
											// AdditionalProperties: map[string]any{},
											SaveFlowVariables: func() *davinci.SaveFlowVariables {
												return &davinci.SaveFlowVariables{
													Value: func() []davinci.SaveFlowVariable {
														return []davinci.SaveFlowVariable{
															{
																// AdditionalProperties: map[string]any{},
																Name: "fdgdfgfdg",
																Value: func() *string {
																	v := "[\n  {\n    \"children\": [\n      {\n        \"text\": \"test124\"\n      }\n    ]\n  }\n]"
																	return &v
																}(),
																Key: func() *float64 {
																	v := float64(0.8936786494474329)
																	return &v
																}(),
																Label: func() *string {
																	v := "fdgdfgfdg (string - flow)"
																	return &v
																}(),
																Type: "string",
															},
															{
																// AdditionalProperties: map[string]any{},
																Name: "test123",
																Value: func() *string {
																	v := "[\n  {\n    \"children\": [\n      {\n        \"text\": \"test456\"\n      }\n    ]\n  }\n]"
																	return &v
																}(),
																Key: func() *float64 {
																	v := float64(0.379286774724122)
																	return &v
																}(),
																Label: func() *string {
																	v := "test123 (number - flow)"
																	return &v
																}(),
																Type: "number",
															},
														}
													}(),
												}
											}(),
										}
									}(),
								},
								Position: &davinci.Position{
									// AdditionalProperties: map[string]any{},
									X: func() *float64 {
										v := float64(270)
										return &v
									}(),
									Y: func() *float64 {
										v := float64(180)
										return &v
									}(),
								},
								Group:      "nodes",
								Removed:    false,
								Selected:   false,
								Selectable: true,
								Locked:     false,
								Grabbable:  true,
								Pannable:   false,
								Classes:    "",
							},
							{
								// AdditionalProperties: map[string]any{},
								Data: &davinci.NodeData{
									// AdditionalProperties: map[string]any{},
									ID: func() *string {
										v := "bbemfztdyk"
										return &v
									}(),
									NodeType: func() *string {
										v := "EVAL"
										return &v
									}(),
								},
								Position: &davinci.Position{
									// AdditionalProperties: map[string]any{},
									X: func() *float64 {
										v := float64(273.5)
										return &v
									}(),
									Y: func() *float64 {
										v := float64(258)
										return &v
									}(),
								},
								Group:      "nodes",
								Removed:    false,
								Selected:   false,
								Selectable: true,
								Locked:     false,
								Grabbable:  true,
								Pannable:   false,
								Classes:    "",
							},
						},
						Edges: []davinci.Edge{
							{
								AdditionalProperties: map[string]any{
									"custom-attribute-1": "custom-attribute-1-value",
									"custom-attribute-2": "custom-attribute-2-value",
								},
								Data: &davinci.Data{
									AdditionalProperties: map[string]any{
										"custom-attribute-1": "custom-attribute-1-value",
										"custom-attribute-2": "custom-attribute-2-value",
									},
									ID: func() *string {
										v := "hseww5vtf0"
										return &v
									}(),
									Source: func() *string {
										v := "1u2m5vzr49"
										return &v
									}(),
									Target: func() *string {
										v := "8fvg7tfr8j"
										return &v
									}(),
								},
								Position: &davinci.Position{
									AdditionalProperties: map[string]any{
										"custom-attribute-1": "custom-attribute-1-value",
										"custom-attribute-2": "custom-attribute-2-value",
									},
									X: func() *float64 {
										v := float64(0)
										return &v
									}(),
									Y: func() *float64 {
										v := float64(0)
										return &v
									}(),
								},
								Group: func() *string {
									v := "edges"
									return &v
								}(),
								Removed: func() *bool {
									v := false
									return &v
								}(),
								Selected: func() *bool {
									v := false
									return &v
								}(),
								Selectable: func() *bool {
									v := true
									return &v
								}(),
								Locked: func() *bool {
									v := false
									return &v
								}(),
								Grabbable: func() *bool {
									v := true
									return &v
								}(),
								Pannable: func() *bool {
									v := true
									return &v
								}(),
								Classes: func() *string {
									v := ""
									return &v
								}(),
							},
							{
								// AdditionalProperties: map[string]any{},
								Data: &davinci.Data{
									// AdditionalProperties: map[string]any{},
									ID: func() *string {
										v := "ljavni2nky"
										return &v
									}(),
									Source: func() *string {
										v := "8fvg7tfr8j"
										return &v
									}(),
									Target: func() *string {
										v := "nx0o1b2cmw"
										return &v
									}(),
								},
								Position: &davinci.Position{
									// AdditionalProperties: map[string]any{},
									X: func() *float64 {
										v := float64(0)
										return &v
									}(),
									Y: func() *float64 {
										v := float64(0)
										return &v
									}(),
								},
								Group: func() *string {
									v := "edges"
									return &v
								}(),
								Removed: func() *bool {
									v := false
									return &v
								}(),
								Selected: func() *bool {
									v := false
									return &v
								}(),
								Selectable: func() *bool {
									v := true
									return &v
								}(),
								Locked: func() *bool {
									v := false
									return &v
								}(),
								Grabbable: func() *bool {
									v := true
									return &v
								}(),
								Pannable: func() *bool {
									v := true
									return &v
								}(),
								Classes: func() *string {
									v := ""
									return &v
								}(),
							},
							{
								// AdditionalProperties: map[string]any{},
								Data: &davinci.Data{
									// AdditionalProperties: map[string]any{},
									ID: func() *string {
										v := "0o2fqy3mf3"
										return &v
									}(),
									Source: func() *string {
										v := "nx0o1b2cmw"
										return &v
									}(),
									Target: func() *string {
										v := "cdcw8k7dnx"
										return &v
									}(),
								},
								Position: &davinci.Position{
									// AdditionalProperties: map[string]any{},
									X: func() *float64 {
										v := float64(0)
										return &v
									}(),
									Y: func() *float64 {
										v := float64(0)
										return &v
									}(),
								},
								Group: func() *string {
									v := "edges"
									return &v
								}(),
								Removed: func() *bool {
									v := false
									return &v
								}(),
								Selected: func() *bool {
									v := false
									return &v
								}(),
								Selectable: func() *bool {
									v := true
									return &v
								}(),
								Locked: func() *bool {
									v := false
									return &v
								}(),
								Grabbable: func() *bool {
									v := true
									return &v
								}(),
								Pannable: func() *bool {
									v := true
									return &v
								}(),
								Classes: func() *string {
									v := ""
									return &v
								}(),
							},
							{
								// AdditionalProperties: map[string]any{},
								Data: &davinci.Data{
									// AdditionalProperties: map[string]any{},
									ID: func() *string {
										v := "493yd0jbi6"
										return &v
									}(),
									Source: func() *string {
										v := "cdcw8k7dnx"
										return &v
									}(),
									Target: func() *string {
										v := "kq5ybvwvro"
										return &v
									}(),
								},
								Position: &davinci.Position{
									// AdditionalProperties: map[string]any{},
									X: func() *float64 {
										v := float64(0)
										return &v
									}(),
									Y: func() *float64 {
										v := float64(0)
										return &v
									}(),
								},
								Group: func() *string {
									v := "edges"
									return &v
								}(),
								Removed: func() *bool {
									v := false
									return &v
								}(),
								Selected: func() *bool {
									v := false
									return &v
								}(),
								Selectable: func() *bool {
									v := true
									return &v
								}(),
								Locked: func() *bool {
									v := false
									return &v
								}(),
								Grabbable: func() *bool {
									v := true
									return &v
								}(),
								Pannable: func() *bool {
									v := true
									return &v
								}(),
								Classes: func() *string {
									v := ""
									return &v
								}(),
							},
							{
								// AdditionalProperties: map[string]any{},
								Data: &davinci.Data{
									// AdditionalProperties: map[string]any{},
									ID: func() *string {
										v := "pn2kixnzms"
										return &v
									}(),
									Source: func() *string {
										v := "j74pmg6577"
										return &v
									}(),
									Target: func() *string {
										v := "ikt13crnhy"
										return &v
									}(),
								},
								Position: &davinci.Position{
									// AdditionalProperties: map[string]any{},
									X: func() *float64 {
										v := float64(0)
										return &v
									}(),
									Y: func() *float64 {
										v := float64(0)
										return &v
									}(),
								},
								Group: func() *string {
									v := "edges"
									return &v
								}(),
								Removed: func() *bool {
									v := false
									return &v
								}(),
								Selected: func() *bool {
									v := false
									return &v
								}(),
								Selectable: func() *bool {
									v := true
									return &v
								}(),
								Locked: func() *bool {
									v := false
									return &v
								}(),
								Grabbable: func() *bool {
									v := true
									return &v
								}(),
								Pannable: func() *bool {
									v := true
									return &v
								}(),
								Classes: func() *string {
									v := ""
									return &v
								}(),
							},
							{
								// AdditionalProperties: map[string]any{},
								Data: &davinci.Data{
									// AdditionalProperties: map[string]any{},
									ID: func() *string {
										v := "0sb4quzlgx"
										return &v
									}(),
									Source: func() *string {
										v := "kq5ybvwvro"
										return &v
									}(),
									Target: func() *string {
										v := "j74pmg6577"
										return &v
									}(),
								},
								Position: &davinci.Position{
									// AdditionalProperties: map[string]any{},
									X: func() *float64 {
										v := float64(0)
										return &v
									}(),
									Y: func() *float64 {
										v := float64(0)
										return &v
									}(),
								},
								Group: func() *string {
									v := "edges"
									return &v
								}(),
								Removed: func() *bool {
									v := false
									return &v
								}(),
								Selected: func() *bool {
									v := false
									return &v
								}(),
								Selectable: func() *bool {
									v := true
									return &v
								}(),
								Locked: func() *bool {
									v := false
									return &v
								}(),
								Grabbable: func() *bool {
									v := true
									return &v
								}(),
								Pannable: func() *bool {
									v := true
									return &v
								}(),
								Classes: func() *string {
									v := ""
									return &v
								}(),
							},
							{
								// AdditionalProperties: map[string]any{},
								Data: &davinci.Data{
									// AdditionalProperties: map[string]any{},
									ID: func() *string {
										v := "v5p4i55lt9"
										return &v
									}(),
									Source: func() *string {
										v := "cdcw8k7dnx"
										return &v
									}(),
									Target: func() *string {
										v := "xb74p6rkd8"
										return &v
									}(),
								},
								Position: &davinci.Position{
									// AdditionalProperties: map[string]any{},
									X: func() *float64 {
										v := float64(0)
										return &v
									}(),
									Y: func() *float64 {
										v := float64(0)
										return &v
									}(),
								},
								Group: func() *string {
									v := "edges"
									return &v
								}(),
								Removed: func() *bool {
									v := false
									return &v
								}(),
								Selected: func() *bool {
									v := false
									return &v
								}(),
								Selectable: func() *bool {
									v := true
									return &v
								}(),
								Locked: func() *bool {
									v := false
									return &v
								}(),
								Grabbable: func() *bool {
									v := true
									return &v
								}(),
								Pannable: func() *bool {
									v := true
									return &v
								}(),
								Classes: func() *string {
									v := ""
									return &v
								}(),
							},
							{
								// AdditionalProperties: map[string]any{},
								Data: &davinci.Data{
									// AdditionalProperties: map[string]any{},
									ID: func() *string {
										v := "k0trrhjqt6"
										return &v
									}(),
									Source: func() *string {
										v := "xb74p6rkd8"
										return &v
									}(),
									Target: func() *string {
										v := "pensvkew7y"
										return &v
									}(),
								},
								Position: &davinci.Position{
									// AdditionalProperties: map[string]any{},
									X: func() *float64 {
										v := float64(0)
										return &v
									}(),
									Y: func() *float64 {
										v := float64(0)
										return &v
									}(),
								},
								Group: func() *string {
									v := "edges"
									return &v
								}(),
								Removed: func() *bool {
									v := false
									return &v
								}(),
								Selected: func() *bool {
									v := false
									return &v
								}(),
								Selectable: func() *bool {
									v := true
									return &v
								}(),
								Locked: func() *bool {
									v := false
									return &v
								}(),
								Grabbable: func() *bool {
									v := true
									return &v
								}(),
								Pannable: func() *bool {
									v := true
									return &v
								}(),
								Classes: func() *string {
									v := ""
									return &v
								}(),
							},
							{
								// AdditionalProperties: map[string]any{},
								Data: &davinci.Data{
									// AdditionalProperties: map[string]any{},
									ID: func() *string {
										v := "2g0chago4l"
										return &v
									}(),
									Source: func() *string {
										v := "pensvkew7y"
										return &v
									}(),
									Target: func() *string {
										v := "vsp1ewtr9m"
										return &v
									}(),
								},
								Position: &davinci.Position{
									// AdditionalProperties: map[string]any{},
									X: func() *float64 {
										v := float64(0)
										return &v
									}(),
									Y: func() *float64 {
										v := float64(0)
										return &v
									}(),
								},
								Group: func() *string {
									v := "edges"
									return &v
								}(),
								Removed: func() *bool {
									v := false
									return &v
								}(),
								Selected: func() *bool {
									v := false
									return &v
								}(),
								Selectable: func() *bool {
									v := true
									return &v
								}(),
								Locked: func() *bool {
									v := false
									return &v
								}(),
								Grabbable: func() *bool {
									v := true
									return &v
								}(),
								Pannable: func() *bool {
									v := true
									return &v
								}(),
								Classes: func() *string {
									v := ""
									return &v
								}(),
							},
							{
								// AdditionalProperties: map[string]any{},
								Data: &davinci.Data{
									// AdditionalProperties: map[string]any{},
									ID: func() *string {
										v := "gs1fx4x303"
										return &v
									}(),
									Source: func() *string {
										v := "3zvjdgdljx"
										return &v
									}(),
									Target: func() *string {
										v := "bbemfztdyk"
										return &v
									}(),
								},
								Position: &davinci.Position{
									// AdditionalProperties: map[string]any{},
									X: func() *float64 {
										v := float64(0)
										return &v
									}(),
									Y: func() *float64 {
										v := float64(0)
										return &v
									}(),
								},
								Group: func() *string {
									v := "edges"
									return &v
								}(),
								Removed: func() *bool {
									v := false
									return &v
								}(),
								Selected: func() *bool {
									v := false
									return &v
								}(),
								Selectable: func() *bool {
									v := true
									return &v
								}(),
								Locked: func() *bool {
									v := false
									return &v
								}(),
								Grabbable: func() *bool {
									v := true
									return &v
								}(),
								Pannable: func() *bool {
									v := true
									return &v
								}(),
								Classes: func() *string {
									v := ""
									return &v
								}(),
							},
							{
								// AdditionalProperties: map[string]any{},
								Data: &davinci.Data{
									// AdditionalProperties: map[string]any{},
									ID: func() *string {
										v := "cum544luro"
										return &v
									}(),
									Source: func() *string {
										v := "bbemfztdyk"
										return &v
									}(),
									Target: func() *string {
										v := "1u2m5vzr49"
										return &v
									}(),
								},
								Position: &davinci.Position{
									// AdditionalProperties: map[string]any{},
									X: func() *float64 {
										v := float64(0)
										return &v
									}(),
									Y: func() *float64 {
										v := float64(0)
										return &v
									}(),
								},
								Group: func() *string {
									v := "edges"
									return &v
								}(),
								Removed: func() *bool {
									v := false
									return &v
								}(),
								Selected: func() *bool {
									v := false
									return &v
								}(),
								Selectable: func() *bool {
									v := true
									return &v
								}(),
								Locked: func() *bool {
									v := false
									return &v
								}(),
								Grabbable: func() *bool {
									v := true
									return &v
								}(),
								Pannable: func() *bool {
									v := true
									return &v
								}(),
								Classes: func() *string {
									v := ""
									return &v
								}(),
							},
						},
					},
					Data: &davinci.Data{
						AdditionalProperties: map[string]any{
							"custom-attribute-1": "custom-attribute-1-value",
							"custom-attribute-2": "custom-attribute-2-value",
						},
					},
					ZoomingEnabled: func() *bool {
						v := true
						return &v
					}(),
					UserZoomingEnabled: func() *bool {
						v := true
						return &v
					}(),
					Zoom: func() *int32 {
						v := int32(1)
						return &v
					}(),
					MinZoom: func() *float64 {
						v := float64(1e-50)
						return &v
					}(),
					MaxZoom: func() *float64 {
						v := float64(1e+50)
						return &v
					}(),
					PanningEnabled: func() *bool {
						v := true
						return &v
					}(),
					UserPanningEnabled: func() *bool {
						v := true
						return &v
					}(),
					Pan: &davinci.Pan{
						AdditionalProperties: map[string]any{
							"custom-attribute-1": "custom-attribute-1-value",
							"custom-attribute-2": "custom-attribute-2-value",
						},
						X: func() *float64 {
							v := float64(0)
							return &v
						}(),
						Y: func() *float64 {
							v := float64(0)
							return &v
						}(),
					},
					BoxSelectionEnabled: func() *bool {
						v := true
						return &v
					}(),
					Renderer: &davinci.Renderer{
						AdditionalProperties: map[string]any{
							"custom-attribute-1": "custom-attribute-1-value",
							"custom-attribute-2": "custom-attribute-2-value",
						},
						Name: func() *string {
							v := "null"
							return &v
						}(),
					},
				},
				InputSchema:  nil,
				OutputSchema: nil,
				Settings:     nil,
				Trigger:      nil,
			},
			FlowColor: func() *string {
				v := "#E3F0FF"
				return &v
			}(),
		},
		FlowEnvironmentMetadata: davinci.FlowEnvironmentMetadata{
			CompanyID:   "2c6123ae-108f-4d11-bcc2-6c8f4dfa9fdb",
			CreatedDate: *davinci.NewEpochTime(1706708769850),
			CustomerID:  "db5f4450b2bd8a56ce076dec0c358a9a",
			FlowID:      "c7062a8857740ee2185694bb855f8f21",
		},
		FlowMetadata: davinci.FlowMetadata{
			AuthTokenExpireIds: []interface{}{},
			Connections:        []interface{}{},
			ConnectorIds: []string{
				"httpConnector",
				"functionsConnector",
				"errorConnector",
				"flowConnector",
				"variablesConnector",
			},
			Description: func() *string {
				v := ""
				return &v
			}(),
			EnabledGraphData:     nil,
			FunctionConnectionID: nil,
			Name:                 "full-basic",
			Orx:                  nil,
			Timeouts:             "null",
			InputSchemaCompiled:  nil,
			IsInputSchemaSaved:   nil,
			IsOutputSchemaSaved:  false,
			OutputSchemaCompiled: nil,
			Variables: []davinci.FlowVariable{
				{
					AdditionalProperties: map[string]any{
						"custom-attribute-1": "custom-attribute-1-value",
						"custom-attribute-2": "custom-attribute-2-value",
					},
					Context: func() *string {
						v := "flow"
						return &v
					}(),
					CreatedDate: davinci.NewEpochTime(1706708735989),
					CustomerID: func() *string {
						v := "db5f4450b2bd8a56ce076dec0c358a9a"
						return &v
					}(),
					Fields: &davinci.FlowVariableFields{
						AdditionalProperties: map[string]any{
							"custom-attribute-1": "custom-attribute-1-value",
							"custom-attribute-2": "custom-attribute-2-value",
						},
						Type: func() *string {
							v := "string"
							return &v
						}(),
						DisplayName: func() *string {
							v := ""
							return &v
						}(),
						Value: nil,
						Mutable: func() *bool {
							v := true
							return &v
						}(),
						Min: func() *int32 {
							v := int32(0)
							return &v
						}(),
						Max: func() *int32 {
							v := int32(2000)
							return &v
						}(),
					},
					FlowID: func() *string {
						v := "c7062a8857740ee2185694bb855f8f21"
						return &v
					}(),
					Type:        "property",
					UpdatedDate: nil,
					Visibility: func() *string {
						v := "private"
						return &v
					}(),
					Name: "fdgdfgfdg##SK##flow##SK##c7062a8857740ee2185694bb855f8f21",
					CompanyID: func() *string {
						v := "2c6123ae-108f-4d11-bcc2-6c8f4dfa9fdb"
						return &v
					}(),
				},
				{
					// AdditionalProperties: map[string]any{},
					Context: func() *string {
						v := "flow"
						return &v
					}(),
					CreatedDate: davinci.NewEpochTime(1706708761083),
					CustomerID: func() *string {
						v := "db5f4450b2bd8a56ce076dec0c358a9a"
						return &v
					}(),
					Fields: &davinci.FlowVariableFields{
						// AdditionalProperties: map[string]any{},
						Type: func() *string {
							v := "number"
							return &v
						}(),
						DisplayName: func() *string {
							v := "test123"
							return &v
						}(),
						Value: func() *string {
							v := "10"
							return &v
						}(),
						Mutable: func() *bool {
							v := true
							return &v
						}(),
						Min: func() *int32 {
							v := int32(4)
							return &v
						}(),
						Max: func() *int32 {
							v := int32(20)
							return &v
						}(),
					},
					FlowID: func() *string {
						v := "c7062a8857740ee2185694bb855f8f21"
						return &v
					}(),
					Type:        "property",
					UpdatedDate: nil,
					Visibility: func() *string {
						v := "private"
						return &v
					}(),
					Name: "test123##SK##flow##SK##c7062a8857740ee2185694bb855f8f21",
					CompanyID: func() *string {
						v := "2c6123ae-108f-4d11-bcc2-6c8f4dfa9fdb"
						return &v
					}(),
				},
			},
		},
		FlowVersionMetadata: davinci.FlowVersionMetadata{
			CurrentVersion: func() *int32 {
				v := int32(8)
				return &v
			}(),
			DeployedDate: davinci.NewEpochTime(1706709739837),
			PublishedVersion: func() *int32 {
				v := int32(8)
				return &v
			}(),
			SavedDate:   *davinci.NewEpochTime(1706708769645),
			UpdatedDate: davinci.NewEpochTime(1706709739837),
			VersionID:   8,
			FlowStatus:  "enabled",
		},
	}
}
