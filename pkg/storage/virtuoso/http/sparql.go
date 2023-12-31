package http

import (
	"fmt"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
)

type operator int

const (
	OR operator = iota
	AND
)

// POSTBuilder creates a query for inserting triples into the graph
func (r virtuosoRepository) POSTBuilder(postBody graph.PostBody, targetGraph graph.TargetGraph) string {
	query := "INSERT DATA { GRAPH <" + string(targetGraph) + "> {\n"
	for _, triple := range postBody.Triples {
		query += "<" + triple[0] + "> <" + triple[1] + "> <" + triple[2] + "> .\n"
	}
	query += "}}"
	return query
}

func (r virtuosoRepository) GETBuilder(edges, subjects, objects []string, depth int, targetGraph graph.TargetGraph) string {
	baseQuery := fmt.Sprintf(
		`SELECT ?%s ?%s ?%s WHERE { GRAPH <%s> { ?%s ?%s ?%s`,
		graph.Subject,
		graph.Predicate,
		graph.Object,
		targetGraph,
		graph.Subject,
		graph.Predicate,
		graph.Object,
	)
	if len(subjects) == 0 && len(objects) == 0 && len(edges) == 0 {
		return baseQuery + ` }}`
	}

	result := baseQuery + ` . FILTER (`

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

	result += ") . }}"
	return result
}

// buildSubQuery builds a subquery, encapsulated by paranthesis
func buildSubQuery(elements []string, attribute graph.Parameter, _op operator) string {
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

func buildContains(attribute graph.Parameter, element string) string {
	return fmt.Sprintf(`contains(str(?%s), '%s')`, attribute, element)
}
