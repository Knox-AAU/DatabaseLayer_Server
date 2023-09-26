package graph

type Node struct {
	Value string  `json:"value"`
	Label *string `json:"label"`
	Child *[]Node `json:"child"`
}
