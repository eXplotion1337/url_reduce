package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"url_reduce/internal/app"
)

type Subj struct {
	ID  string `json:"id"`
	URL string `json:"URL"`
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
		str := string([]byte(b))

		id := app.RandSeq(6) + "-" + app.RandSeq(6)
		strURL := "url=" + app.RandSeq(6) + ".ru"

		w.WriteHeader(http.StatusCreated)
		fmt.Println(string([]byte(b)), r.Method, strURL)

		var settings JSON

		newURL := Subj{
			ID:  id,
			URL: str,
		}

		sett.Obj = append(settings.Obj, newURL)

		w.Write([]byte(str))

		fmt.Println(sett.Obj)

	default:
		w.WriteHeader(400)
	}
}
func GetHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("id")
	for _, v := range sett.Obj {
		if q == v.ID {

			url := strings.ReplaceAll(v.URL, "url=", "")
			url = "https://" + url + "/"
			fmt.Println(url)
			w.WriteHeader(307)
			//w.Header().Add("Location", "url")
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
	http.HandleFunc("/snip", GetHandler)

	//запуск сервера с адресом localhost, порт 8080
	http.ListenAndServe(":8080", nil)

}
