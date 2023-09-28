package dto

import "github.com/AminN77/we-connect/internal"

func MapDtoToFinancialData(dto *FinancialDataDto) *internal.FinancialData {
	return &internal.FinancialData{
		SeriesReference: "",
		Period:          "",
		DataValue:       0,
		Suppressed:      false,
		Status:          "",
		Units:           "",
		Magnitude:       0,
		Subject:         "",
		Group:           "",
		SeriesTitle1:    "",
		SeriesTitle2:    "",
		SeriesTitle3:    "",
		SeriesTitle4:    "",
		SeriesTitle5:    "",
	}
}
