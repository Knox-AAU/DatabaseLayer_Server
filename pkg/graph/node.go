package graph

type Node struct {
	Child *Node   `json:"child"`
	Label *string `json:"label"`
	Value string  `json:"value"`
}
