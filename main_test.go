package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/add?x=5&y=4", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"Action":"add","X":5,"Y":4,"Answer":9,"Cached":false}`, w.Body.String())
}

func TestSubtractRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/subtract?x=5&y=4", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"Action":"subtract","X":5,"Y":4,"Answer":1,"Cached":false}`, w.Body.String())
}

func TestMultiplyRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/multiply?x=5&y=4", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"Action":"multiply","X":5,"Y":4,"Answer":20,"Cached":false}`, w.Body.String())
}

func TestDivideRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/divide?x=20&y=4", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"Action":"divide","X":20,"Y":4,"Answer":5,"Cached":false}`, w.Body.String())
}

func TestCache(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/divide?x=20&y=4", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, `{"Action":"divide","X":20,"Y":4,"Answer":5,"Cached":true}`, w.Body.String())

}
