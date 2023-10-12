package graph

type DataType int

const (
	URI = iota
	Literal
)

const (
	Subject   = "s"
	Predicate = "p"
	Object    = "o"
)

// Triple requires the json tags to match with the queries that are used to retrieve it.
// swagger:model
type Triple struct {
	Subject   BindingAttribute `json:"s"`
	Predicate BindingAttribute `json:"p"`
	Object    BindingAttribute `json:"o"`
}

// swagger:model
type BindingAttribute struct {
	Type  string
	Value string
}

type VirtuosoResponse struct {
	Results struct {
		Bindings []Triple
	}
}
