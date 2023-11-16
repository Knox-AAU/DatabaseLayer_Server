package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/config"
	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
)

type virtuosoRepository struct {
	VirtuosoServerURL config.VirtuosoURL
	GraphURI          config.GraphURI
}

// encode adds necessary parameters for virtuoso
func encode(query string) string {
	params := url.Values{}
	params.Add("query", query)
	params.Add("format", "json")
	return params.Encode()
}

func NewVirtuosoRepository(url config.VirtuosoURL, graphURI config.GraphURI) graph.Repository {
	return &virtuosoRepository{
		VirtuosoServerURL: url,
		GraphURI:          graphURI,
	}
}

func (r virtuosoRepository) ExecuteGET(query string) ([]graph.Triple, error) {
	res, err := http.Get(string(r.VirtuosoServerURL) + "?" + encode(query))
	if err != nil {
		for _, c := range r.VirtuosoServerURL {
			fmt.Print(string(c) + ", ")
		}
		fmt.Print("\n")
		log.Println("error when executing query:", err)
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)

	virtuosoRes := graph.VirtuosoResponse{}
	if err := json.Unmarshal(buf.Bytes(), &virtuosoRes); err != nil {
		return nil, err
	}

	return virtuosoRes.Results.Bindings, nil
}

func (r virtuosoRepository) ExeutePOST(query string) error {
	insertQuery := []byte(query)
	res, err := http.Post(string(r.VirtuosoServerURL), "application/sparql-update", bytes.NewBuffer(insertQuery))
	if err != nil {
		for _, c := range r.VirtuosoServerURL {
			fmt.Print(string(c) + ", ")
		}
		fmt.Print("\n")
		log.Println("error when executing query: ", err)
		return err
	}
	fmt.Println(res)

	return nil
}
