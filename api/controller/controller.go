package controller

import (
	"github.com/AminN77/we-connect/api/dto"
	"github.com/AminN77/we-connect/internal"
	"github.com/gofiber/fiber/v2"
	"net/http"
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
	var response string
	//claims := c.Locals("user").(jwt.MapClaims)
	//customerID := uint(claims["customerId"].(float64))
	//
	//accID, err := strconv.Atoi(c.Params("accountId"))
	//if err != nil {
	//	response.Status = http.StatusBadRequest
	//	response.Message = ErrAccIDNil.Error()
	//	return c.Status(response.Status).JSON(response)
	//}
	//
	//res, err := cc.customerSrv.GetAccount(customerID, uint(accID))
	//if err != nil {
	//	response.Status = http.StatusInternalServerError
	//	response.Message = err.Error()
	//	return c.Status(response.Status).JSON(response)
	//}
	//
	//response.Result = &res
	//response.Status = http.StatusOK
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

	if err := con.srv.Add(&request); err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		return c.Status(response.Status).JSON(response)
	}

	response.Status = http.StatusCreated
	response.Message = "fd created"
	return c.JSON(response)
}
