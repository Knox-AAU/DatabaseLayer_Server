package graph

type DataType int

const (
	URI = iota
	Literal
)

type Node struct {
	Value    string   `json:"value"`
	Label    *string  `json:"label"`
	Children *[]Node  `json:"child"`
	DataType DataType `json:"datatype"`
}

type VirtuosoObject struct {
	Results struct {
		Bindings []struct {
			Subject   BindingAttribute `json:"subject"`
			Predicate BindingAttribute `json:"predicate"`
			Object    BindingAttribute `json:"object"`
		} `json:"bindings"`
	} `json:"results"`
}

type BindingAttribute struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
