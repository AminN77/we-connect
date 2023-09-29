package internal

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewQuery() *Query {
	return &Query{
		IsSuppressed:     false,
		SuppressedFilter: false,
		MaxPeriod:        time.Now().UTC(),
		MinPeriod:        time.Time{},
		MaxDataValue:     1_000_000_000,
		MinDataValue:     0,
	}
}

// Query is the aggregation of the available query params of the repository
type Query struct {
	Limit, Skip                                                  int64
	MaxPeriod                                                    time.Time
	IsSuppressed                                                 bool
	MinPeriod                                                    time.Time
	SuppressedFilter                                             bool
	MaxDataValue, MinDataValue                                   float64
	SeriesReference, Status, Units, Subject, Group, SeriesTitle1 string
	SeriesTitle2, SeriesTitle3, SeriesTitle4, SeriesTitle5       string
}

func buildQuery(q *Query) (*options.FindOptions, interface{}, error) {
	filters := make(bson.M)
	if q.SeriesReference != "" {
		filters["seriesReference"] = q.SeriesReference
	}
	if q.SeriesTitle1 != "" {
		filters["seriesTitle1"] = q.SeriesTitle1
	}
	if q.SeriesTitle2 != "" {
		filters["seriesTitle2"] = q.SeriesTitle2
	}
	if q.SeriesTitle3 != "" {
		filters["seriesTitle3"] = q.SeriesTitle3
	}
	if q.SeriesTitle4 != "" {
		filters["seriesTitle4"] = q.SeriesTitle4
	}
	if q.SeriesTitle5 != "" {
		filters["seriesTitle5"] = q.SeriesTitle5
	}
	if q.Status != "" {
		filters["status"] = q.Status
	}
	if q.Units != "" {
		filters["units"] = q.Units
	}
	if q.Subject != "" {
		filters["subject"] = q.Subject
	}
	if q.Group != "" {
		filters["group"] = q.Group
	}
	if q.SuppressedFilter {
		filters["suppressed"] = q.IsSuppressed
	}
	filters["dataValue"] = bson.M{"$gte": q.MinDataValue, "$lte": q.MaxDataValue}
	filters["period"] = bson.M{"$gte": q.MinPeriod, "$lte": q.MaxPeriod}

	opts := options.Find()
	opts.SetLimit(q.Limit)
	opts.SetSkip(q.Skip)

	return opts, filters, nil
}
