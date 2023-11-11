package model

import (
	"fmt"
)

// magazine Структура для представлення загальної інформації про журнал
type magazine struct {
	Title       string
	Publisher   string
	ReleaseDate string
}

// NewMagazine Конструктор для створення нового журналу
func newMagazine(title, publisher, releaseDate string) magazine {
	return magazine{
		Title:       title,
		Publisher:   publisher,
		ReleaseDate: releaseDate,
	}
}

// PrintInfo Метод для виведення інформації про журнал
func (m magazine) PrintInfo() {
	fmt.Print("\n")
	fmt.Printf("Назва: %s\nВидавець: %s\nДата виходу: %s\n", m.Title, m.Publisher, m.ReleaseDate)

}

//// PrintSliceInfo Метод для виведення інформації с зрізу про журнал
//func (m magazine) PrintSliceInfo(magazines []store.Magazine) error {
//
//}
