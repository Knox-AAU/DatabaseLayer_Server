package graph

type DataType int

const (
	URI = iota
	Literal
	//hello
)

const (
	Subject   = "s"
	Predicate = "p"
	Object    = "o"
)

type Node struct {
	Value    string   `json:"value"`
	Label    *string  `json:"label"`
	Children *[]Node  `json:"child"`
	DataType DataType `json:"datatype"`
}

type Triple struct {
	Subject   BindingAttribute `json:"s"`
	Predicate BindingAttribute `json:"p"`
	Object    BindingAttribute `json:"o"`
}

// VirtuosoResponse requires the json tags to match with the queries that are used to retrieve it.
type VirtuosoResponse struct {
	Results struct {
		Bindings []Triple
	}
}

type BindingAttribute struct {
	Type  string
	Value string
}
