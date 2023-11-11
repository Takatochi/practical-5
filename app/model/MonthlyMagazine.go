package model

import "Pz5/pkg/store"

type MonthlyMagazine struct {
	*magazine
	IssueFrequency string
}

// NewMonthlyMagazine Конструктор для створення нового щомісячного журналу
func NewMonthlyMagazine(title, publisher, releaseDate string) MonthlyMagazine {
	return MonthlyMagazine{
		magazine:       newMagazine(title, publisher, releaseDate),
		IssueFrequency: "Щомісячно",
	}
}

// GetFrequency Метод для отримання частоти виходу щомісячного журналу
func (m MonthlyMagazine) GetFrequency() string {
	return m.IssueFrequency
}

// MagazinesEqual Реалізація методу для порівняння журналів
func (m MonthlyMagazine) MagazinesEqual(other store.Magazines) bool {
	if otherMag, ok := other.(MonthlyMagazine); ok {
		return m.Title == otherMag.Title &&
			m.Publisher == otherMag.Publisher &&
			m.ReleaseDate == otherMag.ReleaseDate &&
			m.IssueFrequency == otherMag.IssueFrequency
	}
	return false
}
