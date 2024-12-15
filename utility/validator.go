package utility

import (
	"errors"
	"fmt"
	validator2 "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"reflect"
	"strings"
)

const (
	defaultErrorFormat = "Field %s with value %s failed on validation rule %s for param %s"
	customErrorTag     = "validationError"
)

type (
	CustomJsonDecoder struct {
		*fiber.Ctx
		*CustomValidator
	}

	errorResponse struct {
		message string
	}

	CustomValidator struct {
	}
)

func NewValidator() *CustomValidator {
	return &CustomValidator{}
}

func NewJsonDecoder(ctx *fiber.Ctx) *CustomJsonDecoder {
	return &CustomJsonDecoder{ctx, NewValidator()}
}

func (c *CustomValidator) Validate(v interface{}) error {
	validator := validator2.New()
	err := validator.Struct(v)

	if err != nil {
		errorsCollected := []errorResponse{}
		value := reflect.ValueOf(v)
		typ := value.Type()

		for _, validationErr := range err.(validator2.ValidationErrors) {
			field, _ := typ.Elem().FieldByName(validationErr.Field())

			// Check for a custom error message in the `errormsg` tag
			customMessage := field.Tag.Get(customErrorTag)

			if customMessage != "" {
				errorsCollected = append(errorsCollected, errorResponse{message: customMessage})
				continue
			}
			// Fall back to default validation message format
			errorsCollected = append(errorsCollected, errorResponse{message: fmt.Sprintf(
				defaultErrorFormat,
				validationErr.Field(),
				validationErr.Value(),
				validationErr.Tag(),
				validationErr.Param(),
			)})
		}

		return errors.New(flatten(errorsCollected))
	}
	return nil
}

func (c *CustomJsonDecoder) Decode(v interface{}) error {
	err := c.BodyParser(&v)
	if err != nil {
		return err
	}
	return c.Validate(v)
}

func flatten(errs []errorResponse) string {
	errStrings := []string{}
	for _, err := range errs {
		errStrings = append(errStrings, err.message)
	}
	return strings.Join(errStrings, ",")

}
