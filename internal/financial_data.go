package internal

import "time"

type FinancialData struct {
	SeriesReference string    `json:"seriesReference" bson:"seriesReference"`
	Period          time.Time `json:"period" bson:"period"`
	DataValue       float64   `json:"dataValue" bson:"dataValue"`
	Suppressed      bool      `json:"suppressed" bson:"suppressed"`
	Status          string    `json:"status" bson:"status"`
	Units           string    `json:"units" bson:"units"`
	Magnitude       int       `json:"magnitude" bson:"magnitude"`
	Subject         string    `json:"subject" bson:"subject"`
	Group           string    `json:"group" bson:"group"`
	SeriesTitle1    string    `json:"seriesTitle1" bson:"seriesTitle1"`
	SeriesTitle2    string    `json:"seriesTitle2" bson:"seriesTitle2"`
	SeriesTitle3    string    `json:"seriesTitle3" bson:"seriesTitle3"`
	SeriesTitle4    string    `json:"seriesTitle4" bson:"seriesTitle4"`
	SeriesTitle5    string    `json:"seriesTitle5" bson:"seriesTitle5"`
}
