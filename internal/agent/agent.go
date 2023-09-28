package agent

import (
	"errors"
	"github.com/AminN77/we-connect/internal"
	csvPkg "github.com/AminN77/we-connect/pkg/csv"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Agent struct {
	repo internal.Repository
	csv  *csvPkg.Csv
	data []internal.FinancialData
	mu   sync.Mutex
}

func New(repo internal.Repository, csv *csvPkg.Csv) *Agent {
	return &Agent{
		repo: repo,
		csv:  csv,
		data: make([]internal.FinancialData, 0),
	}
}

func (a *Agent) Run(file *os.File) {
	// parse csv
	a.csv.ParseFileConcurrent(file, a)

	// insert into mongo
	for i := range a.data {
		if err := a.repo.Add(&a.data[i]); err != nil {
			log.Fatal("could not insert row, err", err.Error())
		}
	}
}

func (a *Agent) Unmarshal(record []string) error {
	if len(record) != 14 {
		return errors.New("len is not 14")
	}

	tempFd := internal.FinancialData{
		SeriesReference: record[0],
		Status:          record[4],
		Units:           record[5],
		Subject:         record[7],
		Group:           record[8],
		SeriesTitle1:    record[9],
		SeriesTitle2:    record[10],
		SeriesTitle3:    record[11],
		SeriesTitle4:    record[12],
		SeriesTitle5:    record[13],
	}

	// Data value
	if record[2] != "" {
		rawDataValue, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatal("could not parse data value, err:", err.Error())
		}
		tempFd.DataValue = rawDataValue
	} else {
		tempFd.DataValue = 0
	}

	// Suppressed
	if record[3] == "" {
		tempFd.Suppressed = false
	} else {
		tempFd.Suppressed = true
	}

	// Magnitude
	rawMag, err := strconv.Atoi(record[6])
	if err != nil {
		log.Fatal("could not parse magnitude, err:", err.Error())
	}
	tempFd.Magnitude = rawMag

	// Period
	rawDate := strings.Split(record[1], ".")
	year, err := strconv.Atoi(rawDate[0])
	month, err := strconv.Atoi(rawDate[1])
	if err != nil {
		log.Fatal("could not parse period, err:", err.Error())
	}
	tempFd.Period = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

	a.mu.Lock()
	a.data = append(a.data, tempFd)
	a.mu.Unlock()

	return nil
}
