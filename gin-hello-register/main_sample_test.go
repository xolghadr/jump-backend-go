package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	e := getEngine()
	req := httptest.NewRequest(http.MethodPost, "/register", nil)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusBadRequest)
	assert.Equal(t, w.Body.String(), `{"message":"firtname is required"}`)
	h := w.Header()["Content-Type"]
	assert.Equal(t, len(h), 1)
	if len(h) == 0 {
		return
	}
	assert.Equal(t, h[0], "application/json; charset=utf-8")
}

func TestRegisterWithInvalidAge(t *testing.T) {
	e := getEngine()
	reader := strings.NewReader(`firstname=John&lastname=Doe&age=ten`)
	req := httptest.NewRequest(http.MethodPost, "/register", reader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusBadRequest)
	assert.Equal(t, w.Body.String(), `{"message":"age should be integer"}`)
}

func TestRegisterWithDuplicatePerson(t *testing.T) {
	e := getEngine()
	reader := strings.NewReader(`firstname=John&lastname=Doe&age=10&job=tailor`)
	req := httptest.NewRequest(http.MethodPost, "/register", reader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, w.Body.String(), `{"message":"John Doe registered successfully"}`)

	dupReader := strings.NewReader(`firstname=John&lastname=Doe&age=33`)
	dupReq := httptest.NewRequest(http.MethodPost, "/register", dupReader)
	dupReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w = httptest.NewRecorder()
	e.ServeHTTP(w, dupReq)
	assert.Equal(t, w.Code, http.StatusConflict)
	assert.Equal(t, w.Body.String(), `{"message":"John Doe registered before"}`)
}

func TestRegisterWithJob(t *testing.T) {
	e := getEngine()
	reader := strings.NewReader(`firstname=John&lastname=Do&age=10&job=tailor`)
	req := httptest.NewRequest(http.MethodPost, "/register", reader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, w.Body.String(), `{"message":"John Do registered successfully"}`)

	getReq := httptest.NewRequest(http.MethodGet, "/hello/john/do", nil)
	w = httptest.NewRecorder()
	e.ServeHTTP(w, getReq)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, w.Body.String(), "Hello John Do; Job: tailor; Age: 10")
}

func TestHello(t *testing.T) {
	e := getEngine()
	req := httptest.NewRequest(http.MethodGet, "/hello/ali/alavi", nil)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusNotFound)
	assert.Equal(t, w.Body.String(), `ali alavi is not registered`)
	h := w.Header()["Content-Type"]
	if len(h) == 0 {
		return
	}
	assert.Equal(t, len(h), 1)
	assert.Equal(t, h[0], "text/plain; charset=utf-8")
}
