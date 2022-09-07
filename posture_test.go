package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShow(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Posture Data", w.Body.String())
}

func TestCreate(t *testing.T) {
	router := setupRouter()

	rec := httptest.NewRecorder()
	data := url.Values{"posture": {"sit"}}
	buffer := bytes.NewBufferString(data.Encode())
	req, _ := http.NewRequest(http.MethodPost, "/", buffer)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, "Created posture=sit", rec.Body.String())
}
