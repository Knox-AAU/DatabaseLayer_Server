package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/Knox-AAU/DatabaseLayer_Server/pkg/graph"
)

type virtuosoRepository struct {
	VirtuosoServerURL string
}

func NewVirtuosoRepository(url string) graph.Repository {
	return &virtuosoRepository{
		VirtuosoServerURL: url,
	}
}

func (r virtuosoRepository) Execute(query string) ([]graph.Triple, error) {
	res, err := http.Get(r.VirtuosoServerURL + "?" + encode(query))
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

func (r virtuosoRepository) ExecutePost(query string, tripleArray []graph.Triple) error {
	jsonPayload, err := json.Marshal(tripleArray)
	if err != nil {
		log.Println("error marshalling JSON: ", err)
		return err
	}

	buf := bytes.NewBuffer(jsonPayload)
	res, err := http.Post(r.VirtuosoServerURL+"?"+query, "application/json", buf)
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

// encode adds necessary parameters for virtuoso
func encode(query string) string {
	params := url.Values{}
	params.Add("query", query)
	params.Add("format", "json")
	return params.Encode()
}
