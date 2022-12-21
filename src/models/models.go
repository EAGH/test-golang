package models

// Test struct example
type Test struct {
	Test    string `json:"test" bson:"test"`
	TestAux Test2  `json:"text_aux" bson:"text_aux"`
}

// Test2 struct example
type Test2 struct {
	Test2 string `json:"test2" bson:"test2"`
}
