package main

import (
	"log"
	"net/http"
)




func main() {
	// Используется функция http.NewServeMux() для инициализации нового рутера, затем
    // регестрируем обработчики для URL-шаблона "/".
	r :=  http.NewServeMux()
	r.HandleFunc("/", home)
	r.HandleFunc("/snippet", showSnippet)
	r.HandleFunc("/snippet/create", createSnippet)


	// Используется функция http.ListenAndServe() для запуска нового веб-сервера. 
    // Мы передаем два параметра: TCP-адрес сети для прослушивания и созданный роутер
	log.Println("Запуск сервера на http://127.0.0.1:8000")
	err := http.ListenAndServe(":8000", r)
	log.Fatal(err)

}