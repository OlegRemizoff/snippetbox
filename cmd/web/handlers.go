package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/OlegRemizoff/snippetbox/pkg/models"
	"github.com/OlegRemizoff/snippetbox/pkg/models/mysql"
)

// Структура  для хранения зависимостей всего веб-приложения.
type Application struct {
	errLog   *log.Logger
	infoLog  *log.Logger
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

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Выводит структуры
	// for _, snippet := range s {
	// 	fmt.Fprintf(w, "%v", snippet)
	// }

	// Создаем экземпляр структуры templateData,
    // содержащий срез с заметками.
	data := &templateData{Snippets: s}
	files := []string{
		"./ui/templates/detail.html",
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
	// Передаем структуру templateData в шаблонизатор.
	// Теперь она будет доступна внутри файлов шаблона через точку.
	err = tmpl.Execute(w, data)
    if err != nil {
        app.serverError(w, err)
    }

	// // Метод записывает содержимое шаблона в тело http ответа
	// // последний параметр позволяет отправлять динамические данные в шаблон
	// err = tmpl.Execute(w, nil)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	app.serverError(w, err)
	// }

}

// Обработчик для отображения содержимого заметки.
// func showSnippet(w http.ResponseWriter, r *http.Request) {
func (app *Application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	// Создаем экземпляр структуры templateData, содержащей данные заметки.
	data := &templateData{Snippet: s}

	files := []string{
		"./ui/templates/detail.html",
		"./ui/templates/home.html",
		"./ui/templates/base.html",
		"./ui/templates/inc/footer.html",
		"./ui/templates/inc/sidebar.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Передаем структуру templateData в качестве данных для шаблона.
	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}
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
