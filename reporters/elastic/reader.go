package elastic

import (
	"time"
	"github.com/lilakurse/data-linter/models"
)


/*
type ReportWriter interface {
	Start() (*models.Report, error)
	Finish(r *models.Report) error
	Commit(r *models.Report, p []*models.Problem) error
}
*/

const (
	IndexNamePrefix string = `dl-report-`
	TypoReportName string = `report`
	TypoProblemName string = `problem`
)

func (e *ElasticWReport) Start () (*models.Report, error) {
	timePrefix := time.Now().Format(`2006-01-02`)
	e.indexName = IndexNamePrefix+timePrefix
	nr := models.NewReport()
	response,err := e.client.Index().Index(e.indexName).Type(TypoReportName).BodyJson(nr).Do()
	if err != nil {
		return nil,err
	}
	e.rid = response.Id
	nr.Name = TypoReportName+`-`+response.Id
	_,err = e.client.Index().Index(e.indexName).Type(TypoReportName).BodyJson(nr).Id(response.Id).Do()
	if err != nil {
		return nil,err
	}
	return nr,nil
}

func (e *ElasticWReport) Commit(r *models.Report, p []*models.Problem) error {
	 r.Update()
	_,err := e.client.Index().Index(e.indexName).Type(TypoReportName).BodyJson(r).Id(e.rid).Do()

	bulk := e.client.Bulk()
	e.PutToBulk(bulk,p)
	_,err = bulk.Do()
	if err != nil {
		r.Error = err.Error()
		_,err = e.client.Index().Index(e.indexName).Type(TypoReportName).BodyJson(r).Id(e.rid).Do()
	}
	return err
}


func (e *ElasticWReport) Finish(r *models.Report) error {
	r.Finish()
	_,err := e.client.Index().Index(e.indexName).Type(TypoReportName).BodyJson(r).Id(e.rid).Do()
	return err
}
