package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	
)



func main() {
	addr := flag.String("addr", ":8000", "Сетевой адрес HTTP") // -addr "127.0.0.1:9999"
	flag.Parse()
	
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile) 

	app := &Application{
		errLog:  errLog,
		infoLog: infoLog,
	}


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

	// Инициализируем новую структуру http.Server и передаем наши данные
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errLog,
		Handler: r,
	}

	// Используется функция http.ListenAndServe() для запуска нового веб-сервера. 
    // Мы передаем два параметра: TCP-адрес сети для прослушивания и созданный роутер
	// err := http.ListenAndServe(*addr, r)  old err := http.ListenAndServe(":8000", r)
	infoLog.Printf("Запуск сервера на %s", *addr)
	err := srv.ListenAndServe()
	errLog.Fatal(err)

}







// Ограничение просмотра файловой системы

// import  "path/filepath"
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