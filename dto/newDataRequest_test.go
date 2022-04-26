package dto

import (
	"net/http"
	"testing"
)

func Test_should_return_error_when_new_data_temperature_is_too_low(t *testing.T) {
	//Arrange
	request := NewDataRequest{
		Temp: -101,
	}

	//Act
	appError := request.Validate()

	//Assert
	if appError.Message != "Check temperature value" {
		t.Error("Invalid message while testing New Data Request")
	}
	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing New Data Request")
	}
}

func Test_should_return_error_when_new_data_temperature_is_too_high(t *testing.T) {
	//Arrange
	request := NewDataRequest{
		Temp: 101,
	}

	//Act
	appError := request.Validate()

	//Assert
	if appError.Message != "Check temperature value" {
		t.Error("Invalid message while testing New Data Request")
	}
	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing New Data Request")
	}
}

func Test_should_return_no_error_when_new_data_temperature_is_ok(t *testing.T) {
	//Arrange
	request := NewDataRequest{
		Temp: 23,
	}

	//Act
	appError := request.Validate()

	//Assert
	if appError != nil {
		t.Error("Validation function is not working")
	}
}
