package main

import (
	"fmt"
	"io"
	"net/http"
	"url_reduce/internal/app"
)

// BodyHandler Обработка запросов
func BodyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// Обработчик метода POST
		b, err := io.ReadAll(r.Body)
		// обрабатываем ошибку
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		var sht = string([]byte(b))
		w.WriteHeader(http.StatusCreated)
		// Вызываем функцию генерации короткой ссылки
		newStr := app.Ramdomize(sht)

		// записываем в тело ответа коротку ссылку
		w.Write([]byte(newStr))
	case "GET":
		// Обработка GET запроса с id сокращенной ссылки
		q := r.URL.Query().Get("query")
		if q == "" {
			http.Error(w, "The query parameter is missing", http.StatusBadRequest)
			return
		}
		//fmt.Println(r.Method)
		//fmt.Println(q)

		value := app.JsonDecoder(q)
		fmt.Println(value)
		// передаем в заголовок location изначальную ссылку
		w.Header().Set("Location", value)

	default:
		// Возвращаем статус код 400
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Старт сервера
func main() {
	//маршрутизация запросов обработчику
	http.HandleFunc("/", BodyHandler)

	//запуск сервера с адресом localhost, порт 8080
	http.ListenAndServe(":8080", nil)

}
