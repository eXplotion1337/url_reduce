package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	case "GET":
		q := r.URL.Query().Get("id")
		if q == "" {
			http.Error(w, "The query parameter is missing", http.StatusBadRequest)
			return
		}
		for _, v := range sett.Obj {
			if q == v.ID {
				w.Header().Set("Location", v.URL)
				w.WriteHeader(307)
				resp, err := json.Marshal(v)
				if err != nil {
					http.Error(w, err.Error(), 500)
					return
				}
				w.Write(resp)
				fmt.Println(r.Method, q)
			}
		}

	default:
		w.WriteHeader(400)
	}
}

// Старт сервера
func main() {
	//маршрутизация запросов обработчику
	http.HandleFunc("/", BodyHandler)

	//запуск сервера с адресом localhost, порт 8080
	http.ListenAndServe(":8080", nil)

}
