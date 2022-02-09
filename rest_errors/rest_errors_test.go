package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerError(t *testing.T) {
	errorMessage := "this is the message"
	err := NewInternalServerError(errorMessage, errors.New("database error"))

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Code)
	assert.EqualValues(t, errorMessage, err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "database error", err.Causes[0])
}

func TestNewBadRequestError(t *testing.T) {
	errorMessage := "this is the message"
	err := NewBadRequestError(errorMessage)

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Code)
	assert.EqualValues(t, errorMessage, err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestNewNotFoundError(t *testing.T) {
	errorMessage := "this is the message"
	err := NewNotFoundError(errorMessage)

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Code)
	assert.EqualValues(t, errorMessage, err.Message)
	assert.EqualValues(t, "not_found", err.Error)
}

func TestNewUnauthorizedError(t *testing.T) {
	errorMessage := "this is the message"
	err := NewUnauthorizedError(errorMessage)

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Code)
	assert.EqualValues(t, errorMessage, err.Message)
	assert.EqualValues(t, "unauthorized", err.Error)
}

func TestNewError(t *testing.T) {
	errorMessage := "this is the message"
	err := NewError(errorMessage)

	assert.EqualValues(t, err.Error(), errorMessage)
}

func TestNewRestError(t *testing.T) {
	var errorsCauses []interface{}

	err := NewRestError("this is the message", 404, "error message", errorsCauses)

	assert.EqualValues(t, err.Message, "this is the message")
	assert.EqualValues(t, err.Code, 404)
	assert.EqualValues(t, err.Error, "error message")
	assert.EqualValues(t, err.Causes, errorsCauses)
}
