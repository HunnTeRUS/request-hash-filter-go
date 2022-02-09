package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/HunnTeRUS/filter-go/rest_errors"
	"github.com/gin-gonic/gin"
)

func FilterTransactionRequest(c *gin.Context) {

	body_hash := c.Request.Header.Get("body_hash")
	request_body_bytes := c.Request.Body

	if request_body_bytes == nil || strings.TrimSpace(body_hash) == "" {
		c.JSON(
			http.StatusBadRequest,
			rest_errors.NewBadRequestError("Request is invalid"),
		)
		return
	}

	request_body, _ := ioutil.ReadAll(request_body_bytes)
	h := sha256.New()
	h.Write(request_body)

	if hex.EncodeToString(h.Sum(nil)) != body_hash {
		c.JSON(
			http.StatusBadRequest,
			rest_errors.NewBadRequestError("Request body hash is different of body_hash header"),
		)
		return
	}

	return
}
