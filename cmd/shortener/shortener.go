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

		value := app.JSONDecoder(q)
		fmt.Println(value)
		// передаем в заголовок location изначальную ссылку
		w.Header().Set("Location", value)

	default:
		// Возвращаем статус код 400
		w.WriteHeader(http.StatusBadRequest)
	}
}

//var form = `<form name="form1" method="post" action="post.php">
//Введите текст:<br />
//<textarea name="text" cols="80" rows="10"></textarea>
//<input name="" type="submit" value="Отправить"/>
//</form>`
//
//func HelloWorld(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte(form))
//}

var form = `<html>
    <head>
    <title></title>
    </head>
    <body>
        <form action="/" method="post">
            <label>Полный URL </label><input type="text" name="FullUrl">
			<input type="submit" value="Login">
            <label>Сокращенный URL</label> <output type="text" name="password">
        </form>
    </body>
</html>`

func Login(w http.ResponseWriter, r *http.Request) {
	// проверяем, каким методом получили запрос
	switch r.Method {
	// если методом POST
	case "POST":
		fullstr := r.FormValue("FullUrl")
		w.WriteHeader(http.StatusCreated)
		//fmt.Fprint(w, fullstr)
		b, err := io.ReadAll(r.Body)
		// обрабатываем ошибку
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Println(b)
		w.Write([]byte(fullstr))

		// проверяем пароль вспомогательной функцией
		// при успешной авторизации обрабатываем запрос
		// например, передаём другому обработчику
		//AuthorisedHandler(w, r)
		// в остальных случаях предлагаем форму авторизации
	case "GET":
		w.WriteHeader(307)
		//w.Write([]byte(r.Method))

	}
}

// Auth — вспомогательная функция авторизации
// за пределами урока реализация может выглядеть так

//var Logins = make(map[string]string)
func BodyHandlerr(w http.ResponseWriter, r *http.Request) {
	// читаем Body
	switch r.Method {
	case "POST":
		b, err := io.ReadAll(r.Body)
		// обрабатываем ошибку
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Println(string([]byte(b)), r.Method)

	case "GET":
		q := r.URL.Query().Get("id")
		if q == "" {
			http.Error(w, "The query parameter is missing", http.StatusBadRequest)
			return
		}
		w.WriteHeader(307)
		fmt.Println(r.Method, q)
	default:
		w.WriteHeader(400)
	}

	// в нашем случае q примет значение "something"
	// продолжаем обработку запроса
	// ...
	// продолжаем обработку
	// ...
}

// Старт сервера
func main() {
	//маршрутизация запросов обработчику
	http.HandleFunc("/", BodyHandlerr)

	//запуск сервера с адресом localhost, порт 8080
	http.ListenAndServe(":8080", nil)

}
