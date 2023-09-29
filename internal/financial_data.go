package internal

import "time"

// FinancialData is the data model representation of this application domain
type FinancialData struct {
	DataValue  float64   `json:"dataValue" bson:"dataValue"`   // 8B
	Magnitude  int       `json:"magnitude" bson:"magnitude"`   // 8B
	Period     time.Time `json:"period" bson:"period"`         // ~ 12B -> 16B
	Suppressed bool      `json:"suppressed" bson:"suppressed"` // 1B -> 4B

	SeriesReference string `json:"seriesReference" bson:"seriesReference"`
	Status          string `json:"status" bson:"status"`
	Units           string `json:"units" bson:"units"`
	Subject         string `json:"subject" bson:"subject"`
	Group           string `json:"group" bson:"group"`
	SeriesTitle1    string `json:"seriesTitle1" bson:"seriesTitle1"`
	SeriesTitle2    string `json:"seriesTitle2" bson:"seriesTitle2"`
	SeriesTitle3    string `json:"seriesTitle3" bson:"seriesTitle3"`
	SeriesTitle4    string `json:"seriesTitle4" bson:"seriesTitle4"`
	SeriesTitle5    string `json:"seriesTitle5" bson:"seriesTitle5"`
}

// NOTE: This struct aligned in the way which reduces the extra memory padding needed for the allocation.
// Arm64 arch is the target.
