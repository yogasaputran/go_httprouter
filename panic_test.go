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

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, error interface{}) {
		fmt.Fprint(w, "Panic : ", error)
	}

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("Ups")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	reecorder := httptest.NewRecorder()

	router.ServeHTTP(reecorder, request)
	response := reecorder.Result()

	body, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Panic : Ups", string(body))
}
