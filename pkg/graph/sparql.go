package graph

import (
	"fmt"
)

type operator int

const (
	OR operator = iota
	AND
)

func Builder(edges, subjects, objects []string, depth int, uri string) string {
	baseQuery := fmt.Sprintf(`SELECT ?%s ?%s ?%s WHERE { GRAPH <%s> { ?%s ?%s ?%s`, Subject, Predicate, Object, uri, Subject, Predicate, Object)
	if len(subjects) == 0 && len(objects) == 0 && len(edges) == 0 {
		return baseQuery + ` }}`
	}

	result := baseQuery + ` . FILTER (`

	if len(subjects) > 0 {
		result += buildSubQuery(subjects, Subject, OR)
	}

	if len(objects) > 0 {
		if len(subjects) > 0 {
			result += " && "
		}
		result += buildSubQuery(objects, Object, OR)
	}

	if len(edges) > 0 {
		if len(subjects) > 0 || len(objects) > 0 {
			result += " && "
		}
		result += buildSubQuery(edges, Predicate, OR)
	}

	result += ") . }}"
	return result
}

// buildSubQuery builds a subquery, encapsulated by paranthesis
func buildSubQuery(elements []string, attribute Attribute, _op operator) string {
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

func buildContains(attribute Attribute, element string) string {
	return fmt.Sprintf(`contains(str(?%s), '%s')`, attribute, element)
}

// PostQueryBuilder creates the query needed to insert the triples into the virtuoso database
func PostQueryBuilder(triples []Triple) string {
	query := "INSERT DATA {\n"

	for _, triple := range triples {
		query += fmt.Sprintf("  %s %s %s .\n", triple.S.Value, triple.P.Value, triple.O.Value)
	}

	query += "}"

	return query
}
