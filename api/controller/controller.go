package controller

import (
	"context"
	"github.com/AminN77/we-connect/api/dto"
	"github.com/AminN77/we-connect/internal"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"net/mail"
	"strconv"
	"time"
)

type Controller struct {
	srv internal.Service
}

func NewController(srv internal.Service) *Controller {
	return &Controller{
		srv: srv,
	}
}

// Get godoc
// @Summary      List financial data
// @Description  get financial data
// @Tags         financial data
// @Accept       json
// @Produce      json
// @Param        skip    query     int  false "skip" default(0) minimum(0)
// @Param        limit    query     int  false "limit" default(10) minimum(1) maximum(100)
// @Param        seriesReference    query     string  false  "search by series reference"
// @Param        seriesTitle1    query     string  false  "search by series title 1"
// @Param        seriesTitle2    query     string  false  "search by series title 2"
// @Param        seriesTitle3    query     string  false  "search by series title 3"
// @Param        seriesTitle4    query     string  false  "search by series title 4"
// @Param        seriesTitle5    query     string  false  "search by series title 5"
// @Param        status    query     string  false  "search by status" Enums(F,R,C)
// @Param        units    query     string  false  "search by units"
// @Param        subject    query     string  false  "search by subject"
// @Param        group    query     string  false  "search by group"
// @Param        suppressedFilter    query     bool  false  "enable suppressed filter"
// @Param        isSuppressed    query     bool  false  "search by suppressed value"
// @Param        maxDataValue    query     number  false  "upper bound for data value"
// @Param        minDataValue    query     number  false  "lower bound for data value"
// @Param        maxPeriod    query     string  false  "upper bound for period" example("Thu, 20 Dec 2020 00:00:00 MDT")
// @Param        minPeriod    query     string  false  "lower bound for period" example("Thu, 20 Dec 2020 00:00:00 MDT")
// @Router       /financialData [get]
func (con *Controller) Get(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var response dto.Response[[]*internal.FinancialData]
	q := internal.NewQuery()
	if err := con.bindQuery(q, c); err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		return c.Status(response.Status).JSON(response)
	}
	
	res, err := con.srv.Get(q, ctx)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		return c.Status(response.Status).JSON(response)
	}

	response.Result = &res
	response.Status = http.StatusOK
	response.ResulCount = len(res)
	return c.JSON(response)
}

func (con *Controller) bindQuery(q *internal.Query, c *fiber.Ctx) error {
	var err error
	q.Skip, err = strconv.ParseInt(c.Query("skip"), 10, 64)
	q.Limit, err = strconv.ParseInt(c.Query("limit"), 10, 64)
	q.SeriesReference = c.Query("seriesReference")
	q.Status = c.Query("status")
	q.Group = c.Query("group")
	q.Units = c.Query("units")
	q.SeriesTitle1 = c.Query("seriesTitle1")
	q.SeriesTitle2 = c.Query("seriesTitle2")
	q.SeriesTitle3 = c.Query("seriesTitle3")
	q.SeriesTitle4 = c.Query("seriesTitle4")
	q.SeriesTitle5 = c.Query("seriesTitle5")

	supF, errP := strconv.ParseBool(c.Query("suppressedFilter"))
	if errP == nil {
		if supF {
			IsSup, errP := strconv.ParseBool(c.Query("isSuppressed"))
			if errP == nil {
				q.SuppressedFilter = true
				q.IsSuppressed = IsSup
			}
		}
	}

	if c.Query("maxDataValue") != "" {
		q.MaxDataValue, err = strconv.ParseFloat(c.Query("maxDataValue"), 64)
	}

	if c.Query("minDataValue") != "" {
		q.MinDataValue, err = strconv.ParseFloat(c.Query("minDataValue"), 64)
	}

	if c.Query("maxPeriod") != "" {
		q.MaxPeriod, err = mail.ParseDate(c.Query("maxPeriod"))
	}

	if c.Query("minPeriod") != "" {
		q.MinPeriod, err = mail.ParseDate(c.Query("minPeriod"))
	}

	return err
}
