package model

import (
	"Pz5/app/model/Data"
	"fmt"
)

// Magazine Структура для представлення загальної інформації про журнал
type Magazine struct {
	Description    Data.Description `json:"description"`
	IssueFrequency string           `json:"issue_frequency"`
}

// newMagazine Конструктор для створення нового журналу
func newMagazine(title, publisher, releaseDate, issueFrequency string) *Magazine {
	return &Magazine{
		Description: Data.Description{Title: title,
			Publisher:   publisher,
			ReleaseDate: releaseDate,
		}, IssueFrequency: issueFrequency,
	}
}

// PrintInfo Метод для виведення інформації про журнал
func (m Magazine) PrintInfo() {
	fmt.Printf("Назва: %s\nВидавець: %s\nДата виходу: %s\n", m.Description.Title, m.Description.Publisher, m.Description.ReleaseDate)
	fmt.Print("\n")
}

// GetFrequency Метод для отримання частоти виходу журналу, що виходить раз на кілька місяців
func (m Magazine) GetFrequency() string {
	return m.IssueFrequency
}
