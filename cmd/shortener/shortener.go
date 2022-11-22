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
        <form action="/login" method="post">
            <label>Логин</label><input type="text" name="login">
            <label>Пароль<input type="password" name="password">
            <input type="submit" value="Login">
        </form>
    </body>
</html>`

func Login(w http.ResponseWriter, r *http.Request) {
	// проверяем, каким методом получили запрос
	switch r.Method {
	// если методом POST
	case "POST":
		login := r.FormValue("login")
		password := r.FormValue("password")
		// проверяем пароль вспомогательной функцией
		if !Auth(login, password) {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			// если пароль не верен, указываем код ошибки в заголовке
			w.WriteHeader(401)
			// пишем в тело ответа
			fmt.Fprintln(w, r.Method)
			return
		}
		// при успешной авторизации обрабатываем запрос
		// например, передаём другому обработчику
		//AuthorisedHandler(w, r)
		// в остальных случаях предлагаем форму авторизации
	default:
		fmt.Fprint(w, form)
	}
}

// Auth — вспомогательная функция авторизации
// за пределами урока реализация может выглядеть так
func Auth(l, p string) bool {
	pass, ok := Logins[l]
	return ok && p == pass
}

var Logins = make(map[string]string)

// Старт сервера
func main() {
	//маршрутизация запросов обработчику
	http.HandleFunc("/", Login)

	//запуск сервера с адресом localhost, порт 8080
	http.ListenAndServe(":8080", nil)

}
