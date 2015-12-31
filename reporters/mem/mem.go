package mem

import (
	"github.com/lilakurse/data-linter/models"
	"strconv"
	"strings"
)

type MemWriter struct {
	reports []*models.Report
}

func New() *MemWriter {
	return &MemWriter{}
}

func (m *MemWriter) Start() (*models.Report, error) {
	nr := models.NewReport()
	nr.Name = "mem-" + strconv.Itoa(nr.Created.Nanosecond())
	m.reports = append(m.reports, nr)
	return nr, nil
}

func (m *MemWriter) Commit(r *models.Report, p []*models.Problem) error {
	r.Update()
	r.Problems = append(r.Problems, p...)
	return nil
}

func (m *MemWriter) Finish(r *models.Report) error {
	r.Finish()
	return nil
}

func (m *MemWriter) PrintReport() {
	println("REPORTS")
	println("______________________________________________________________")
	println("Index | Name \t| Created \t\t| Updated \t\t| Error | Finished \t")
	for index, report := range m.reports {
		text := strings.Join([]string{strconv.Itoa(index+1) + "\t", report.Name, report.Created.Format("2006-01-02 15:04:05"), report.Updated.Format("2006-01-02 15:04:05"), report.Error, report.Finished.Format("2006-01-02 15:04:05")}, ` | `)
		println(text)
		println("\n\tTotal/Inspected= ( " + strconv.Itoa(int(report.Statistics.Total)) + " / " + strconv.Itoa(int(report.Statistics.Inspected)) + " )")
		println("\tValid/Invalid= ( " + strconv.Itoa(int(report.Statistics.Valid)) + " / " + strconv.Itoa(int(report.Statistics.Invalid)) + " )")
		println("\tCount Problems= " + strconv.Itoa(len(report.Problems)))
		println("\n\t\tPROBLEMS")
		println("______________________________________________________________")
		for j, problem := range report.Problems {
			println("\n\t\tINDEX = " + strconv.Itoa(j+1))
			println("\t\tID = " + problem.Id)
			println("\t\tOriginal = " + problem.Original)
			println("\t\tFirst Description = " + problem.Details[0].Description + "... Count Of ditails " + strconv.Itoa(len(problem.Details)))
		}
	}
}
