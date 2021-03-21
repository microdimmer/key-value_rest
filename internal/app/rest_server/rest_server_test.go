package rest_server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRESTServer_Hello(t *testing.T) {
	s := New(NewConfig())

	rec := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, "Hello", rec.Body.String())
}
