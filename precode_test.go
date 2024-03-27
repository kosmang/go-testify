package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=3&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	responseStatusCode := responseRecorder.Result().StatusCode
	require.Equal(t, http.StatusOK, responseStatusCode)
	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=3&city=astana", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	responseStatusCode := responseRecorder.Result().StatusCode
	assert.Equal(t, http.StatusBadRequest, responseStatusCode)
	assert.Equal(t, []byte("wrong city value"), responseRecorder.Body.Bytes())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=5&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	expectedList := []byte("Мир кофе, Сладкоежка, Кофе и завтраки, Сытый студент")
	assert.Equal(t, expectedList, responseRecorder.Body.Bytes())
}
