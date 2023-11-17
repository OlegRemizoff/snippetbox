package mysql

import (
	"database/sql"
	"errors"
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
	stmt := `SELECT id, title, content, created, expires FROM snippets 
			 WHERE expires > UTC_TIMESTAMP() AND id = ?`
	row := m.DB.QueryRow(stmt, id)
	// Инициализируем указатель на новую структуру Snippet.
	s := &models.Snippet{}

	// row.Scan(), чтобы скопировать значения из каждого поля от sql.Row
	// аргументы для row.Scan - это указатели на место, куда требуется скопировать данные
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	// err := m.DB.QueryRow("SELECT ...", id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires) вместо stmt
	if err != nil {
		//Если ошибка обнаружена, то возвращаем нашу ошибку из модели models.ErrNoRecords.Is(err, sql.ErrNoRows)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	// Возвращается объект Snippet.
	return s, nil

}