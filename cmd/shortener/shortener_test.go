package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"url_reduce/cmd/hendlers"
	//"url_reduce/cmd/shortener"
)

func TestGetHandler(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}
	// создаём массив тестов: имя и желаемый результат
	tests := []struct {
		name string
		want want
	}{
		// определяем все тесты
		{
			name: "positive test #1",
			want: want{
				code: 307,
			},
		},
	}
	for _, tt := range tests {
		// запускаем каждый тест
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/ser", nil)

			// создаём новый Recorder
			w := httptest.NewRecorder()
			// определяем хендлер
			h := http.HandlerFunc(hendlers.GetHandler)
			// запускаем сервер
			h.ServeHTTP(w, request)
			res := w.Result()
			defer res.Body.Close()
			// проверяем код ответа
			if res.StatusCode != tt.want.code {
				t.Errorf("Expected status code %d, got %d", tt.want.code, w.Code)
			}
		})
	}

}
