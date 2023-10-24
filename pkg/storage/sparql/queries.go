package sparql

import (
	"fmt"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
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
		result += buildSubQuery(subjects, graph.Subject)
	}

	if len(objects) > 0 {
		if len(subjects) > 0 {
			result += " && "
		}
		result += buildSubQuery(objects, graph.Object)
	}

	if len(edges) > 0 {
		if len(subjects) > 0 || len(objects) > 0 {
			result += " && "
		}
		result += buildSubQuery(edges, graph.Predicate)
	}

	result += ") . }"
	return result
}

// buildSubQuery starts building the subquery with a ||
func buildSubQuery(elements []string, attribute string) string {
	result := "(" + buildContains(attribute, elements[0])

	elements = elements[1:]
	for _, element := range elements {
		result += " || " + buildContains(attribute, element)
	}
	return result + ")"
}

func buildContains(attribute, element string) string {
	return fmt.Sprintf(`contains(str(?%s), '%s')`, attribute, element)
}
