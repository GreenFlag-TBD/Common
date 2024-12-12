package utility

import (
	"errors"
	"fmt"
	validator2 "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strings"
)

const (
	validationMessageFormat = "%s: '%v' | Needs to conform '%s': '%s'"
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
		for _, err := range err.(validator2.ValidationErrors) {
			errorsCollected = append(errorsCollected, errorResponse{message: fmt.Sprintf(
				validationMessageFormat,
				err.Field(),
				err.Value(),
				err.Tag(),
				err.Param(),
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
