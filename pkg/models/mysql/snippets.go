package postgres

import (
	"database/sql"
	"github.com/OlegRemizoff/snippetbox/pkg/models"

)


type SnippetModel struct {
	DB *sql.DB
}


func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}