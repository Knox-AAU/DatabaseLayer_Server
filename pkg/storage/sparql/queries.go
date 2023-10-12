package sparql

import (
	"fmt"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
)

var GetAll = fmt.Sprintf(`SELECT ?%s ?%s ?%s WHERE { ?%s ?%s ?%s }`,
	graph.Subject, graph.Predicate, graph.Object, graph.Subject, graph.Predicate, graph.Object)
