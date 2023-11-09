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

// POSTBuilder creates the query needed to insert the triples into the virtuoso database
func (r virtuosoRepository) POSTBuilder(triples []graph.Triple) string {
	query := "INSERT DATA {"
	query += "GRAPH <" + string(r.GraphURI) + ">{"
	for _, triple := range triples { //This might not be needed since the http.Post()
		query += "<" + triple.S.Value + "> <" + triple.P.Value + "> <" + triple.O.Value + ">." //takes the triples in the request body
	}
	query += "}}"
	return query
}

func (r virtuosoRepository) GETBuilder(edges, subjects, objects []string, depth int) string {
	baseQuery := fmt.Sprintf(
		`SELECT ?%s ?%s ?%s WHERE { GRAPH <%s> { ?%s ?%s ?%s`,
		graph.Subject,
		graph.Predicate,
		graph.Object,
		r.GraphURI,
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
func buildSubQuery(elements []string, attribute graph.Attribute, _op operator) string {
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

func buildContains(attribute graph.Attribute, element string) string {
	return fmt.Sprintf(`contains(str(?%s), '%s')`, attribute, element)
}
