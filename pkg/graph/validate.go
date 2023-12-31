package graph

import "fmt"

func (p PostBody) Validate() error {
	const length = 3
	for _, triple := range p.Triples {
		if len(triple) != length {
			return fmt.Errorf("got %d elements in %T, expected %d", len(triple), p, length)
		}
	}
	return nil
}

func (g TargetGraph) Validate() error {
	if g == "" {
		return fmt.Errorf("target graph cannot be empty")
	}
	return nil
}
