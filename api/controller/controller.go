package controller

import "github.com/gofiber/fiber/v2"

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (cc *Controller) Get(c *fiber.Ctx) error {
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
