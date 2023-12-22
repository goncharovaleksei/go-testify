package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=7&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	arrResponce := strings.Split(responseRecorder.Body.String(), ",")
	assert.Equal(t, totalCount, len(arrResponce))
}

func TestMainHandlerAnswerCodeAndBodyForEmpty(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=7&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())
}

func TestMainHandlerWhenCityInRequestIsWrong(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=7&city=spb", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

/*
v Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
v Город, который передаётся в параметре city, не поддерживается. Сервис возвращает код ответа 400 и ошибку wrong count value в теле ответа.
v Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
*/
