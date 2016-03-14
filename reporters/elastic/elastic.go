package elastic

import (
	"github.com/lilakurse/data-linter/models"
	el "gopkg.in/olivere/elastic.v3"
	"fmt"
)

type ElasticWReport struct {
	client    *el.Client
	indexName string
	rid       string
}

func New() (*ElasticWReport, error) {
	client, err := el.NewClient()
	if err != nil {
		return nil, fmt.Errorf("NewClient: %s",err.Error())
	}
	info, code, err := client.Ping(el.DefaultURL).Do()
	if err != nil {
		return nil, fmt.Errorf("Ping: %s",err.Error())
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s", code, info.Version.Number)
	return &ElasticWReport{client,"",""}, nil
}

func (e *ElasticWReport) PutToBulk(bulk *el.BulkService, p []*models.Problem) {
	for _, problem := range p {
		problem.ReportId = e.rid
		bulk.Add(el.NewBulkIndexRequest().Index(e.indexName).Type(TypoProblemName).Doc(problem))
	}
}
