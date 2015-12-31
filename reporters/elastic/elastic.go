package elastic

import (
	"github.com/lilakurse/data-linter/models"
	el "gopkg.in/olivere/elastic.v3"
)

type ElasticWReport struct {
	client    el.Client
	indexName string
	rid       string
}

func New() (*ElasticWReport, error) {
	client, err := el.NewClient()
	if err != nil {
		return nil, err
	}
	_, _, err = client.Ping(el.DefaultURL).Do()
	if err != nil {
		return nil, err
	}
	return &ElasticWReport{client}, nil
}

func (e *ElasticWReport) PutToBulk(bulk el.BulkService, p []*models.Problem) {
	for _, problem := range p {
		problem.ReportId = e.rid
		bulk.Add(el.NewBulkIndexRequest().Index(e.indexName).Type(TypoProblemName).Doc(problem))
	}
}
