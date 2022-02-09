package middleware

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HunnTeRUS/filter-go/rest_errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFilterTransactionRequest_sending_nil_body_and_receiving_error(t *testing.T) {

	w := httptest.NewRecorder()
	r := gin.New()
	r.POST("/testFilter", FilterTransactionRequest)

	req, _ := http.NewRequest("POST", "/testFilter", nil)
	r.ServeHTTP(w, req)

	var err *rest_errors.RestErr

	if err := json.Unmarshal(w.Body.Bytes(), &err); err != nil {
		t.Fail()
		return
	}

	assert.NotNil(t, err)
	assert.EqualValues(t, err.Code, http.StatusBadRequest)
	assert.EqualValues(t, err.Message, "Request is invalid")
}

func TestFilterTransactionRequest_sending_invalid_header_and_receiving_error(t *testing.T) {

	w := httptest.NewRecorder()
	r := gin.New()
	r.POST("/testFilter", FilterTransactionRequest)

	requestByte, _ := json.Marshal(map[string]interface{}{
		"test": "test",
	})
	req, _ := http.NewRequest("POST", "/testFilter", bytes.NewReader(requestByte))
	r.ServeHTTP(w, req)

	var err *rest_errors.RestErr

	if err := json.Unmarshal(w.Body.Bytes(), &err); err != nil {
		t.Fail()
		return
	}

	assert.NotNil(t, err)
	assert.EqualValues(t, err.Code, http.StatusBadRequest)
	assert.EqualValues(t, err.Message, "Request is invalid")
}

func TestFilterTransactionRequest_sending_different_body_and_hash(t *testing.T) {

	w := httptest.NewRecorder()
	r := gin.New()
	r.POST("/testFilter", FilterTransactionRequest)

	requestByte, _ := json.Marshal(map[string]interface{}{
		"test": "test",
	})
	req, _ := http.NewRequest("POST", "/testFilter", bytes.NewReader(requestByte))
	req.Header.Add("body_hash", "TEST")
	r.ServeHTTP(w, req)

	var err *rest_errors.RestErr

	if err := json.Unmarshal(w.Body.Bytes(), &err); err != nil {
		t.Fail()
		return
	}

	assert.NotNil(t, err)
	assert.EqualValues(t, err.Code, http.StatusBadRequest)
	assert.EqualValues(t, err.Message, "Request body hash is different of body_hash header")
}

func TestFilterTransactionRequest_sending_equal_body_and_hash(t *testing.T) {

	w := httptest.NewRecorder()
	r := gin.New()
	r.POST("/testFilter", FilterTransactionRequest)

	requestByte, _ := json.Marshal(map[string]interface{}{
		"test": "test",
	})

	h := sha256.New()
	h.Write(requestByte)

	req, _ := http.NewRequest("POST", "/testFilter", bytes.NewReader(requestByte))
	req.Header.Add("body_hash", hex.EncodeToString(h.Sum(nil)))
	r.ServeHTTP(w, req)

	var err *rest_errors.RestErr
	json.Unmarshal(w.Body.Bytes(), &err)

	assert.Nil(t, err)
}
