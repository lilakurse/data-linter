package multi

import (
	"github.com/lilakurse/data-linter/models"
)

type ReportWriter interface {
	Start() (*models.Report, error)
	Finish(r *models.Report) error
	Commit(r *models.Report, p []*models.Problem) error
}

type MultiWReporter struct {
	reporters []ReportWriter
}

func New(writers ...ReportWriter) *MultiWReporter {
	mr := &MultiWReporter{}
	mr.reporters = append(mr.reporters, writers...)
	return mr
}

func (m *MultiWReporter) Start() (*models.Report, error) {
	var (
		report *models.Report
		err    error
	)
	//block mode
	for _, r := range m.reporters {
		report, err = r.Start()
		if err != nil {
			return nil, err
		}
	}
	return report, nil
}

func (m *MultiWReporter) Commit(r *models.Report, p []*models.Problem) error {
	var err error
	//block mode
	for _, rr := range m.reporters {
		err = rr.Commit(r, p)
		if err != nil {
			return err
		}
	}
	return err
}

func (m *MultiWReporter) Finish(r *models.Report) error {
	var err error
	//block mode
	for _, rr := range m.reporters {
		err = rr.Finish(r)
		if err != nil {
			return err
		}
	}
	return nil
}
