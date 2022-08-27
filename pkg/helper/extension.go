package helper

import (
	"github.com/labstack/echo/v4"
)

func GetCorrelationID(c echo.Context) string {
	corId := c.Request().Header.Get("X-Correlation-Id")
	return corId
}

func CustomerSetGETendptToLocationHeader(c echo.Context, id string) {
	url := "/customer/" + id
	c.Response().Header().Set("Location", url)
}

func OrderSetGETendptToLocationHeader(c echo.Context, id string) {
	url := "/order/" + id
	c.Response().Header().Set("Location", url)
}
