package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestMethodNotAllowed(t *testing.T) {
	router := httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Gak Boleh")
	})

	router.POST("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "POST")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	reecorder := httptest.NewRecorder()

	router.ServeHTTP(reecorder, request)
	response := reecorder.Result()

	body, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Gak Boleh", string(body))
}
