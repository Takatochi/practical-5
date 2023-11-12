package model

import (
	"Pz5/app/model/Data"
	"fmt"
)

// Magazines Структура для представлення загальної інформації про журнал
type Magazines struct {
	Description    Data.Description `json:"description"`
	IssueFrequency string           `json:"issue_frequency"`
}

// NewMagazine Конструктор для створення нового журналу
func NewMagazine(title, publisher, releaseDate, issueFrequency string) *Magazines {
	return &Magazines{
		Description: Data.Description{Title: title,
			Publisher:   publisher,
			ReleaseDate: releaseDate,
		}, IssueFrequency: issueFrequency,
	}
}

// PrintInfo Метод для виведення інформації про журнал
func (m Magazines) PrintInfo() {
	fmt.Printf("Назва: %s\nВидавець: %s\nДата виходу: %s\n", m.Description.Title, m.Description.Publisher, m.Description.ReleaseDate)
	fmt.Print("\n")
}

// GetFrequency Метод для отримання частоти виходу журналу, що виходить раз на кілька місяців
func (m Magazines) GetFrequency() string {
	return m.IssueFrequency
}

//// SaveUsers функція для збереження користувачів у JSON-файл
//func (m magazines) SaveUsers(filename string) error {
//
//}
