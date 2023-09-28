package graph

type DataType int

const (
	URI = iota
	Literal
	BNode
	TypedLiteral
)

type Node struct {
	Value    string   `json:"value"`
	Label    *string  `json:"label"`
	Children *[]*Node `json:"child"`
	DataType DataType `json:"datatype"`
}

type VirtuosoObject struct {
	Results struct {
		Bindings []Binding `json:"bindings"`
	} `json:"results"`
}

type Binding struct {
	Subject   BindingAttribute `json:"subject"`
	Predicate BindingAttribute `json:"predicate"`
	Object    BindingAttribute `json:"object"`
}

type BindingAttribute struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
