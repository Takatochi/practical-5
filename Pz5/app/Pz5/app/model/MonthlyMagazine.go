package model

type MonthlyMagazine struct {
	magazine
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
