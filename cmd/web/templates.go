package main


import (
	"github.com/OlegRemizoff/snippetbox/pkg/models"

)


// Создаем тип templateData, который будет действовать как хранилище для
// любых динамических данных, которые нужно передать HTML-шаблонам.
type templateData struct {
    Snippet *models.Snippet
	Snippets []*models.Snippet
}