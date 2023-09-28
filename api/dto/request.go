package dto

type FinancialDataDto struct {
	SeriesReference string `json:"seriesReference"`
	Period          string `json:"period"`
	DataValue       string `json:"dataValue"`
	Suppressed      string `json:"suppressed"`
	Status          string `json:"status"`
	Units           string `json:"units"`
	Magnitude       string `json:"magnitude"`
	Subject         string `json:"subject"`
	Group           string `json:"group"`
	SeriesTitle1    string `json:"seriesTitle1"`
	SeriesTitle2    string `json:"seriesTitle2"`
	SeriesTitle3    string `json:"seriesTitle3"`
	SeriesTitle4    string `json:"seriesTitle4"`
	SeriesTitle5    string `json:"seriesTitle5"`
}
