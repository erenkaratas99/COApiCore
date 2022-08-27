package pkg

import (
	"github.com/erenkaratas99/COApiCore/pkg/customErrors"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func LimitOffsetValidation(l string, o string) (int64, int64) {
	offset, err := strconv.ParseInt(o, 10, 64)
	if err != nil {
		return 25, 0
	}
	limit, err := strconv.ParseInt(l, 10, 64)
	if err != nil {
		return 25, 0
	}
	if l != "" && o != "" {
		if offset < 0 {
			return 25, 0
		}
		if limit > 100 || limit <= 0 {
			return 25, 0
		}
	}
	return limit, offset
}

type CustomValidator struct {
	validator *validator.Validate
}

func NewValidation() echo.Validator {
	return &CustomValidator{validator: validator.New()}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return customErrors.NewHTTPError(http.StatusBadRequest,
			"ValidationErr",
			"Some parts are missing or wrong.")
	}
	return nil
}
