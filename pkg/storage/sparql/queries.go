package sparql

import (
	"fmt"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
)

type operator int

const (
	OR operator = iota
	AND
)

var GetAll = fmt.Sprintf(`SELECT ?%s ?%s ?%s WHERE { ?%s ?%s ?%s }`,
	graph.Subject, graph.Predicate, graph.Object, graph.Subject, graph.Predicate, graph.Object)

func Builder(edges, subjects, objects []string, depth int) string {
	if len(subjects) == 0 && len(objects) == 0 && len(edges) == 0 {
		return GetAll
	}

	result := fmt.Sprintf(`SELECT ?%s ?%s ?%s WHERE { ?%s ?%s ?%s . FILTER (`,
		graph.Subject, graph.Predicate, graph.Object, graph.Subject, graph.Predicate, graph.Object)

	if len(subjects) > 0 {
		result += buildSubQuery(subjects, graph.Subject, OR)
	}

	if len(objects) > 0 {
		if len(subjects) > 0 {
			result += " && "
		}
		result += buildSubQuery(objects, graph.Object, OR)
	}

	if len(edges) > 0 {
		if len(subjects) > 0 || len(objects) > 0 {
			result += " && "
		}
		result += buildSubQuery(edges, graph.Predicate, OR)
	}

	result += ") . }"
	return result
}

// buildSubQuery builds a subquery, encapsulated by paranthesis
func buildSubQuery(elements []string, attribute string, _op operator) string {
	if len(elements) == 0 {
		return ""
	}

	result := "(" + buildContains(attribute, elements[0])
	op := ""

	switch _op {
	case OR:
		op = " || "
	case AND:
		op = " && "
	}

	elements = elements[1:]
	for _, element := range elements {
		result += op + buildContains(attribute, element)
	}
	return result + ")"
}

func buildContains(attribute, element string) string {
	return fmt.Sprintf(`contains(str(?%s), '%s')`, attribute, element)
}
