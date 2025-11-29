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

	// TODO Надо к этому вернуться ! реализовать провервку времени даты сразу в validate
	// Регистрируем тег 'dateformat'
	// err = validate.RegisterValidation("dateformat", func(fl validator.FieldLevel) bool {
	// 	value := fl.Field().String()
	// 	layout := fl.Param() // формат передаём через параметр

	// 	// Пытаемся разобрать строку как дату
	// 	_, err := time.Parse(layout, value)
	// 	return err == nil // true, если parse прошёл без ошибок
	// })
	return nil
}
