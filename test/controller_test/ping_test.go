package controller_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reservation/internal/router"
	"testing"
)

func TestPing(t *testing.T) {
	ginSever := router.Ping()
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	ginSever.ServeHTTP(res, req)
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "pong", res.Body.String())
}
