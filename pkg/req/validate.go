package req

import (
	"github.com/go-playground/validator"
)

func IsValid[T any](payload T) error {
	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {

		return err
	}
	return nil
}
