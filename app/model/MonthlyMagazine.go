package model

import mag "Pz5/pkg/magazine"

type MonthlyMagazine struct {
	*Magazines
}

// NewMonthlyMagazine Конструктор для створення нового щомісячного журналу
func NewMonthlyMagazine(title, publisher, releaseDate string) MonthlyMagazine {
	return MonthlyMagazine{
		Magazines: NewMagazine(title, publisher, releaseDate, "Щомісячно"),
	}
}

// GetFrequency Метод для отримання частоти виходу щомісячного журналу
func (m MonthlyMagazine) GetFrequency() string {
	return m.Magazines.GetFrequency()
}

// MagazinesEqual Реалізація методу для порівняння журналів
func (m MonthlyMagazine) MagazinesEqual(other mag.Magazines) bool {
	if otherMag, ok := other.(MonthlyMagazine); ok {
		return m.Description.Title == otherMag.Description.Title &&
			m.Description.Publisher == otherMag.Description.Publisher &&
			m.Description.ReleaseDate == otherMag.Description.ReleaseDate &&
			m.IssueFrequency == otherMag.IssueFrequency
	}
	return false
}
