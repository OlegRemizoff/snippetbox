package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)



// Помощник serverError записывает сообщение об ошибке в errorLog и
// затем отправляет пользователю ответ 500 "Внутренняя ошибка сервера"
func (app *Application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errLog.Output(2, trace)
}


// Помощник clientError отправляет определенный код состояния и соответствующее описание пользователю
func (app *Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}


// Удобная оболочка вокруг clientError, которая отправляет пользователю ответ "404 Страница не найдена".
func (app *Application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}