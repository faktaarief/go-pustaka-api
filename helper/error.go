package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ResponIfErrors(errors error) (err interface{}) {
	if errors != nil {
		fmt.Println(errors)
		errorMessages := []string{}
		for _, error := range errors.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", error.Field(), error.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		return errorMessages
	}

	return nil
}
