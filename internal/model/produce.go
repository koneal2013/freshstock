package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Money float64

func (m Money) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"$%.2f\"", m)), nil
}

func (m *Money) UnmarshalJSON(data []byte) error {
	var num float64
	if err := json.Unmarshal(data, &num); err != nil {
		return err
	}

	// Set the value of m to the float64 value
	*m = Money(num)

	return nil
}

type Produce struct {
	Code      string `json:"code" binding:"required,code,len=19"`
	Name      string `json:"name" binding:"required,alphanum"`
	UnitPrice Money  `json:"unit_price" binding:"required,gte=0"`
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
