package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

type Subj struct {
	ID      string `json:"id"`
	LongURL string `json:"LongURL"`
}

type JSON struct {
	Obj []Subj
}

var sett JSON

func BodyHandler(w http.ResponseWriter, r *http.Request) {
	// читаем Body
	switch r.Method {
	case "POST":
		b, err := io.ReadAll(r.Body)
		// обрабатываем ошибку
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		str := string(b)
		long := strings.ReplaceAll(str, "11111", "://")
		long = strings.ReplaceAll(long, "22222", "/")
		id := strconv.Itoa(rand.Int())
		strURL := "http://localhost:8080/ser?id=" + id

		w.WriteHeader(http.StatusCreated)

		var settings JSON

		newURL := Subj{
			ID:      id,
			LongURL: long,
		}

		sett.Obj = append(settings.Obj, newURL)

		w.Write([]byte(strURL))

		fmt.Println(sett.Obj)

	default:
		w.WriteHeader(400)
	}
}
func GetHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("id")
	for _, v := range sett.Obj {
		if q == v.ID {

			url := strings.ReplaceAll(v.LongURL, "url=", "")
			//w.WriteHeader(307)
			//fmt.Println(r.Header)
			w.Header().Set("Location", url)
			fmt.Println(w.Header())
			http.Redirect(w, r, url, http.StatusTemporaryRedirect)

			//w.WriteHeader(307)
		}
	}
	w.WriteHeader(307)
}

// Старт сервера
func main() {
	//маршрутизация запросов обработчику
	http.HandleFunc("/", BodyHandler)
	http.HandleFunc("/ser", GetHandler)

	//запуск сервера с адресом localhost, порт 8080
	http.ListenAndServe(":8080", nil)

}
