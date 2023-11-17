package mysql

import (
	"database/sql"
	"github.com/OlegRemizoff/snippetbox/pkg/models"

)

// SnippetModel - Определяем тип который обертывает пул подключения sql.DB
type SnippetModel struct {
	DB *sql.DB
}


// Insert - Метод для создания новой заметки в базе дынных.
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	// ? - плейсхолдер, для данных, которых требуется вставить в базу данных
	stmt := `INSERT INTO snippets (title, content, created, expires)
			 VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`	 
	// метод db.Eec() возвращает объект sql.Result, который содержит некоторые основные
	// данные о том, что произошло после выполнении запроса.
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	// LastInsertId() позволяет получить  последний ID созданной записи из таблицу snippets.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil

}


// Get - Метод для возвращения данных заметки по её идентификатору ID.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}