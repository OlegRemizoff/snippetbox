package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //  _ для исбежания ошибки, main.go ничего не использует из этого пакетп
	"github.com/OlegRemizoff/snippetbox/pkg/models/mysql"

)


func main() {
	addr := flag.String("addr", ":8000", "Сетевой адрес HTTP") // -addr "127.0.0.1:9999"
	dsn := flag.String("dsn", "oleg:1234@/snippetbox?parseTime=true", "Название MySQL источника данных")

	flag.Parse()
	
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile) 

	// Передаем в отдельную функцию openDB()
	// полученный  источник данных (DSN) из флага командной строки.
	db, err := openDB(*dsn)
	if err != nil {
		errLog.Fatal(err)
	}

	defer db.Close()

	app := &Application{
		errLog: errLog,
		infoLog:  infoLog,
		snippets: &mysql.SnippetModel{DB: db},
	}


	// Инициализируем новую структуру http.Server и передаем наши данные
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  app.routes(),
	}

	// Используется функция http.ListenAndServe() для запуска нового веб-сервера. 
    // Мы передаем два параметра: TCP-адрес сети для прослушивания и созданный роутер
	// err := http.ListenAndServe(*addr, r)  old err := http.ListenAndServe(":8000", r)
	infoLog.Printf("Запуск сервера на %s", *addr)
	err = srv.ListenAndServe()
	errLog.Fatal(err)

}


// Функция openDB() обертывает sql.Open() и возвращает пул соединений sql.DB 
// для заданной строки подключения (DSN)
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn) // возвращает объект sql.DB - это пул множества соединений
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil { // для создания соединения с MySQL и проверки на наличие ошибок
		return nil, err
	}
	return db, nil
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