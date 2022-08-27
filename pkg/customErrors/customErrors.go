package customErrors

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type HttpError struct {
	Code    int    `json:"code"`
	Key     string `json:"error"`
	Message string `json:"message"`
}

func NewHTTPError(errcode int, key string, msg string) *HttpError {
	return &HttpError{
		Code:    errcode,
		Key:     key,
		Message: msg,
	}
}

var (
	BindErr     = NewHTTPError(http.StatusBadRequest, "BindErr", "Request could not have been binded.")
	DocNotFound = NewHTTPError(http.StatusBadRequest, "NotFoundErr", "There is no matching document with given parameter.")
)

func (e *HttpError) Error() string {
	return e.Key + ": " + e.Message
}

func SeperateCustomErrLogs(code int, c echo.Context, cId, key, msg string) {
	if code >= 500 {
		log.WithFields(log.Fields{
			"correlationID": cId,
			"status":        code,
			"fullPath":      c.Request().URL.Path,
			"rawPath":       c.Path(),
		}).Error(key, " : ", msg)
	} else {
		log.WithFields(log.Fields{
			"correlationID": cId,
			"status":        code,
			"fullPath":      c.Request().URL.Path,
			"rawPath":       c.Path(),
		}).Info(key, " : ", msg)
	}
}

func ReturnErrResponse(c echo.Context, code int, key, msg string) {
	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			err := c.NoContent(code)
			if err != nil {
				log.Error(err)
			}
		} else {
			err := c.JSON(code, NewHTTPError(code, key, msg))
			if err != nil {
				log.Error(err)
			}
		}
	}
}
