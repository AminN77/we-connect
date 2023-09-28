package controller

import (
	"github.com/AminN77/we-connect/api/dto"
	"github.com/AminN77/we-connect/internal"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"net/mail"
	"strconv"
)

type Controller struct {
	srv internal.Service
}

func NewController(srv internal.Service) *Controller {
	return &Controller{
		srv: srv,
	}
}

func (con *Controller) Get(c *fiber.Ctx) error {
	var response dto.Response[[]*internal.FinancialData]
	q := internal.NewQuery()
	if err := con.bindQuery(q, c); err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		return c.Status(response.Status).JSON(response)
	}

	res, err := con.srv.Get(q)
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

func (con *Controller) Add(c *fiber.Ctx) error {
	var request internal.FinancialData
	var response dto.BaseResponse

	if err := c.BodyParser(&request); err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		return c.Status(response.Status).JSON(response)
	}

	if err := con.srv.Insert(&request); err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		return c.Status(response.Status).JSON(response)
	}

	response.Status = http.StatusCreated
	response.Message = "fd created"
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
