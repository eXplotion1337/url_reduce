package main

import (
	"net/http"
	"url_reduce/cmd/hendlers"
)

// Старт сервера
func main() {
	//маршрутизация запросов обработчику
	http.HandleFunc("/", hendlers.BodyHandler)
	http.HandleFunc("/ser", hendlers.GetHandler)

	//запуск сервера с адресом localhost, порт 8080
	http.ListenAndServe(":8080", nil)

}
