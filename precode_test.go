package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenOK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=3&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	responseStatusCode := responseRecorder.Result().StatusCode
	assert.Equal(t, http.StatusOK, responseStatusCode)
	assert.NotNil(t, responseRecorder.Body)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=3&city=astana", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	responseStatusCode := responseRecorder.Result().StatusCode
	assert.Equal(t, http.StatusBadRequest, responseStatusCode)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=5&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	parameterCount := req.URL.Query().Get("count")
	requestedCount, err := strconv.Atoi(parameterCount)
	if err != nil {
		t.Errorf("want type string, have %T: %v", parameterCount, err)
	}

	assert.GreaterOrEqual(t, requestedCount, totalCount)
}
