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
	longURL string `json:"URL"`
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
		long := strings.ReplaceAll(str, "1", "https://")
		long = strings.ReplaceAll(long, "2", "/")
		id := strconv.Itoa(rand.Int())
		//id := app.RandSeq(6) + "-" + app.RandSeq(6)
		strURL := "http://localhost:8080/ser?id=" + id

		w.WriteHeader(http.StatusCreated)
		//fmt.Println(string([]byte(b)), r.Method, strURL)

		var settings JSON

		newURL := Subj{
			ID:      id,
			longURL: long,
		}

		sett.Obj = append(settings.Obj, newURL)

		w.Write([]byte(strURL))

		fmt.Println(sett.Obj)
	//case "GET":
	//	q := r.URL.Query().Get("/")
	//	fmt.Println(string(q))
	//	for _, v := range sett.Obj {
	//		if q == v.ID {
	//
	//			url := strings.ReplaceAll(v.longURL, "url=", "")
	//			//url = "https://" + url + "/"
	//			fmt.Println(url)
	//			w.WriteHeader(307)
	//			w.Header().Add("Location", "text/plain; charset=utf-8")
	//			w.Header().Set("Location", url)
	//			//http.Redirect(w, r, url, http.StatusMovedPermanently)
	//			//w.WriteHeader(307)
	//
	//		}
	//	}

	default:
		w.WriteHeader(400)
	}
}
func GetHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("id")
	for _, v := range sett.Obj {
		if q == v.ID {

			url := strings.ReplaceAll(v.longURL, "url=", "")
			//url = "https://" + url + "/"
			fmt.Println(url)
			w.WriteHeader(307)
			w.Header().Add("Location", "text/plain; charset=utf-8")
			w.Header().Set("Location", url)
			http.Redirect(w, r, url, http.StatusMovedPermanently)
			//w.WriteHeader(307)

		}
	}
}

// Старт сервера
func main() {
	//маршрутизация запросов обработчику
	http.HandleFunc("/", BodyHandler)
	http.HandleFunc("/ser", GetHandler)

	//запуск сервера с адресом localhost, порт 8080
	http.ListenAndServe(":8080", nil)

}
