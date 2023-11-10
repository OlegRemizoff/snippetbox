package main

import "net/http"




func(app *Application) routes() *http.ServeMux {
	// Используется функция http.NewServeMux() для инициализации нового рутера, затем
    // регестрируем обработчики для URL-шаблона "/".
	r :=  http.NewServeMux()
	r.HandleFunc("/", app.home)
	r.HandleFunc("/snippet", app.showSnippet)
	r.HandleFunc("/snippet/create", app.createSnippet)


	// Регистрации обработчика для всех запросов, которые начинаются с "/static/"
	// Убираем префикс "/static" перед тем как запрос достигнет http.FileServer
	fileServer := http.FileServer(http.Dir("./ui/static"))
	r.Handle("/static/", http.StripPrefix("/static", fileServer))
	
	return r
}