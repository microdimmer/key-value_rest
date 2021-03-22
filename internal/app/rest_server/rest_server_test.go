package rest_server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestRESTServer_Upsert(t *testing.T) {
	s := New(NewConfig())
	s.configureRouter()
	rec := httptest.NewRecorder()

	data := &DataObject{}
	data.Key = "asd"
	data.Value = "asd"
	payload, _ := json.Marshal(data)
	req, _ := http.NewRequest(http.MethodPost, "/upsert", bytes.NewBuffer(payload))
	s.handleUpsert().ServeHTTP(rec, req)
	assert.Equal(t, payload, rec.Body.Bytes())
}

func TestRESTServer_Delete(t *testing.T) {
	s := New(NewConfig())
	s.configureRouter()
	// s.Start()
	rec := httptest.NewRecorder()

	data := &DataObject{}
	data.Key = "asd"
	data.Value = "asd"
	payload, _ := json.Marshal(data)
	req, _ := http.NewRequest(http.MethodPost, "/upsert", bytes.NewBuffer(payload))
	s.handleUpsert().ServeHTTP(rec, req)
	assert.Equal(t, payload, rec.Body.Bytes())

	rec = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodDelete, "/delete", nil)
	vars := map[string]string{
		"key": data.Key,
	}
	s.handleDelete().ServeHTTP(rec, mux.SetURLVars(req, vars))
	mapB, _ := json.Marshal(map[string]string{"message": "Record succesfully deleteted"})
	assert.Equal(t, string(mapB), rec.Body.String())
}

func TestRESTServer_List(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()

	data := &DataObject{}
	data.Key = "asd"
	data.Value = "asd"
	payload, _ := json.Marshal(data)
	req, _ := http.NewRequest(http.MethodPost, "/upsert", bytes.NewBuffer(payload))
	s.handleUpsert().ServeHTTP(rec, req)
	assert.Equal(t, payload, rec.Body.Bytes())

	rec = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPost, "/list", nil)
	s.handleList().ServeHTTP(rec, req)
	mapB, _ := json.Marshal(map[string]string{"asd": "asd"})

	assert.Equal(t, string(mapB), rec.Body.String())
}

func TestRESTServer_Get(t *testing.T) {
	s := New(NewConfig())
	s.configureRouter()
	rec := httptest.NewRecorder()

	data := &DataObject{}
	data.Key = "asd"
	data.Value = "testing_get"
	payload, _ := json.Marshal(data)
	req, _ := http.NewRequest(http.MethodPost, "/upsert", bytes.NewBuffer(payload))
	s.handleUpsert().ServeHTTP(rec, req)
	assert.Equal(t, payload, rec.Body.Bytes())

	rec = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/get", nil)
	vars := map[string]string{
		"key": data.Key,
	}
	s.handleGet().ServeHTTP(rec, mux.SetURLVars(req, vars))
	assert.Equal(t, string(payload), rec.Body.String())

}
