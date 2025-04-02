package test

import "github.com/samir-gandhi/davinci-client-go/davinci"

type TestModel struct {
	AdditionalProperties map[string]any     `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Test1                *string            `davinci:"test1Field,environmentmetadata,omitempty"`
	Test2                *string            `davinci:"test2Field,config,omitempty"`
	Test3                *davinci.EpochTime `davinci:"test3Field,versionmetadata,omitempty"`
	Test4                *string            `davinci:"test4Field,environmentmetadata,omitempty"`
	Test5                *TestModel2        `davinci:"test5Field,*,omitempty"`
	Test6                *string            `davinci:"test6Field,environmentmetadata,omitempty"`
	Test7                *float64           `davinci:"test7Field,flowmetadata,omitempty"`
	Test8                *string            `davinci:"test8Field,config,omitempty"`
	Test9                *string            `davinci:"test9Field,designercue,omitempty"`
	Test10               *string            `davinci:"test10Field,config,omitempty"`
	Test11               *string            `davinci:"test11Field,designercue,omitempty"`
	Test12               *string            `davinci:"test12Field,versionmetadata,omitempty"`
}

type TestModel2 struct {
	AdditionalProperties map[string]any     `json:"-" davinci:"-,unmappedproperties"` // used to capture all other properties that are not explicitly defined in the model
	Test1                *string            `davinci:"test1Field,environmentmetadata,omitempty"`
	Test2                *string            `davinci:"test2Field,config,omitempty"`
	Test3                *davinci.EpochTime `davinci:"test3Field,versionmetadata,omitempty"`
	Test4                *string            `davinci:"test4Field,environmentmetadata,omitempty"`
	Test6                *string            `davinci:"test6Field,-,omitempty"`
	Test7                *float64           `davinci:"test7Field,flowmetadata,omitempty"`
	Test8                *string            `davinci:"test8Field,config,omitempty"`
	Test9                *string            `davinci:"test9Field,designercue,omitempty"`
	Test10               *string            `davinci:"test10Field,config,omitempty"`
	Test11               *string            `davinci:"test11Field,designercue,omitempty"`
	Test12               *string            `davinci:"test12Field,versionmetadata,omitempty"`
}
