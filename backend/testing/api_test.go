package testing

import (
	"backend/api"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func ApiTest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := api.BasicApi()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status code 200")
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String(), "Expected body to be '{\"message\":\"pong\"}'")
}
