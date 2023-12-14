package model

import (
	"errors"
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Produce struct {
	Code      string  `json:"code" binding:"required,code,len=19"`
	Name      string  `json:"name" binding:"required,alphanum"`
	UnitPrice float64 `json:"unit_price" binding:"required,gte=0"`
}

func RegisterValidator() error {
	// register custom field validator for model.Produce.Code
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("code", validateCode)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("invalid validator")
}

func validateCode(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[A-Za-z0-9-]*$`)
	return re.MatchString(fl.Field().String())
}
