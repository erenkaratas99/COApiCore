package middleware

import (
	"github.com/erenkaratas99/COApiCore/pkg/customErrors"
	"github.com/erenkaratas99/COApiCore/pkg/helper"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const CorrelationID = "X-Correlation-Id"

func AddCorrelationID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Request().Header.Get(CorrelationID)
		var newID string
		if id == "" {
			newID = uuid.New().String()
		} else {
			newID = id
		}
		c.Request().Header.Set(CorrelationID, newID)
		c.Response().Header().Set(CorrelationID, newID)
		res := next(c)

		return res
	}
}

func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.WithFields(log.Fields{
			"correlationID": c.Request().Header.Get(CorrelationID),
			"method":        c.Request().Method,
			"fullPath":      c.Request().URL.Path,
			"rawPath":       c.Path(),
		}).Info("Request details")
		res := next(c)
		return res
	}
}

func Recovery(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		cId := helper.GetCorrelationID(c)
		defer func() {
			err := recover()
			if err != nil {
				log.WithFields(log.Fields{
					"fullPath":      c.Request().URL.Path,
					"rawPath":       c.Path(),
					"correlationID": cId,
				}).Error("PANIC : ", err)

				c.JSON(http.StatusInternalServerError, customErrors.NewHTTPError(http.StatusInternalServerError,
					"ServerErr", http.StatusText(500)))
			}
		}()
		err := next(c)
		var (
			code = http.StatusInternalServerError
			key  = "ServerError"
			msg  = "Internal Server Error"
		)
		if err != nil {
			if he, ok := err.(*customErrors.HttpError); ok {
				if he.Code >= 500 {
					customErrors.ReturnErrResponse(c, code, key, msg)
				}
				code = he.Code
				key = he.Key
				msg = he.Message
				customErrors.SeperateCustomErrLogs(code, c, cId, key, msg)
				customErrors.ReturnErrResponse(c, code, key, msg)
			} else {
				log.WithFields(log.Fields{
					"fullPath":      c.Request().URL.Path,
					"rawPath":       c.Path(),
					"correlationID": cId,
				}).Error(err)
				customErrors.ReturnErrResponse(c, code, key, msg)
			}
		}
		return err
	}
}
