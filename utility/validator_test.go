package utility

import (
	"fmt"
	"testing"
)

type testStruct struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
}

type testStructWithCustomError struct {
	Email    string `json:"email" validate:"required,email" validationError:"Email is missing or not valid"`
	Username string `json:"username" validate:"required" validationError:"Username is required"`
}

func TestValidate_Success(t *testing.T) {
	test := testStruct{
		Email:    "test@gmail.com",
		Username: "test",
	}

	vl := NewValidator()
	err := vl.Validate(test)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

}

func TestValidate_Fail(t *testing.T) {
	test := testStruct{
		Email:    "example",
		Username: "test",
	}

	vl := NewValidator()
	err := vl.Validate(&test)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	field := "Email"
	value := "example"
	rule := "email"
	param := ""

	if err.Error() != fmt.Sprintf(defaultErrorFormat, field, value, rule, param) {
		t.Error(fmt.Sprintf("Expected %s, got %s", fmt.Sprintf(defaultErrorFormat, field, value, rule, param), err.Error()))
	}
}

func TestValidate_Fail_CustomError(t *testing.T) {
	test := testStructWithCustomError{
		Email:    "example",
		Username: "",
	}

	vl := NewValidator()
	err := vl.Validate(&test)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expected := "Email is missing or not valid,Username is required"
	if err.Error() != expected {
		t.Errorf("Expected error message to be 'Email is missing or not valid', got %v", err.Error())
	}
}
