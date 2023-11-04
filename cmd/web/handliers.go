package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"strconv"
)


// Обработчик главной страницы
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string {
		"./ui/templates/home.html",
		"./ui/templates/base.html",
		"./ui/templates/inc/footer.html",
		"./ui/templates/inc/sidebar.html",
	}
	
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}


	// Метод записывает содержимое шаблона в тело http ответа
	// последний параметр позволяет отправлять динамические данные в шаблон
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

} 


// Обработчик для отображения содержимого заметки.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Отображение выбранной заметки с ID %d...", id)
}


// Обработчик для создания заметки.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод запрещен!", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Создания новой заметки..."))
}





