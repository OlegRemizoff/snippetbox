package main

import (
	"log"
	"net/http"
	// "path/filepath"
)




func main() {
	// Используется функция http.NewServeMux() для инициализации нового рутера, затем
    // регестрируем обработчики для URL-шаблона "/".
	r :=  http.NewServeMux()
	r.HandleFunc("/", home)
	r.HandleFunc("/snippet", showSnippet)
	r.HandleFunc("/snippet/create", createSnippet)


	// Регистрации обработчика для всех запросов, которые начинаются с "/static/"
	// Убираем префикс "/static" перед тем как запрос достигнет http.FileServer
	fileServer := http.FileServer(http.Dir("./ui/static"))
	r.Handle("/static/", http.StripPrefix("/static", fileServer))


	// Используется функция http.ListenAndServe() для запуска нового веб-сервера. 
    // Мы передаем два параметра: TCP-адрес сети для прослушивания и созданный роутер
	log.Println("Запуск сервера на http://127.0.0.1:8000")
	err := http.ListenAndServe(":8000", r)
	log.Fatal(err)

}







// Ограничение просмотра файловой системы
// type neuteredFileSystem struct {
// 	fs http.FileSystem
// }

// func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
// 	f, err := nfs.fs.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}

// 	s, err := f.Stat()
// 	if s.IsDir() {
// 		index := filepath.Join(path, "index.html")
// 		if _, err := nfs.fs.Open(index); err != nil {
// 			closeErr := f.Close()
// 			if closeErr != nil {
// 				return nil, closeErr
// 			}

// 			return nil, err
// 		}
// 	}

// 	return f, nil
// }


// main () {
// 	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./static")})
//     mux.Handle("/static", http.NotFoundHandler())
//     mux.Handle("/static/", http.StripPrefix("/static", fileServer))
// }