package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"strconv"
	"github.com/OlegRemizoff/snippetbox/pkg/models/mysql"
)



type Application struct {
	errLog 	*log.Logger
	infoLog *log.Logger
	snippets *mysql.SnippetModel

}


// old Обработчик главной страницы
// func home(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		http.NotFound(w, r)
// 		return
// 	}

// Обработчик главной страницы
func (app *Application) home(w http.ResponseWriter, r *http.Request) {
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
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		app.serverError(w, err)
		return
	}


	// Метод записывает содержимое шаблона в тело http ответа
	// последний параметр позволяет отправлять динамические данные в шаблон
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		app.serverError(w, err)
	}

} 


// Обработчик для отображения содержимого заметки.
// func showSnippet(w http.ResponseWriter, r *http.Request) {
func (app *Application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Отображение выбранной заметки с ID %d...", id)
}


// Обработчик для создания заметки.
func (app *Application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	title := "История про улитку"
	content := "Улитка выползла из раковины,\nвытянула рожки,\nи опять подобрала их."
	expires := "7"
	// Передаем данные в метод SnippetModel.Insert(), получая обратно
	// ID только что созданной записи в базу данных.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)

}





