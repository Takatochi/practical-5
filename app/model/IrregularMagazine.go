package model

import "Pz5/pkg/store"

type IrregularMagazine struct {
	*magazine
	IssueFrequency string
}

// NewIrregularMagazine Конструктор для створення нового журналу, що виходить раз на кілька місяців
func NewIrregularMagazine(title, publisher, releaseDate string) IrregularMagazine {
	return IrregularMagazine{
		magazine:       newMagazine(title, publisher, releaseDate),
		IssueFrequency: "Раз на кілька місяців",
	}
}

// GetFrequency Метод для отримання частоти виходу журналу, що виходить раз на кілька місяців
func (i IrregularMagazine) GetFrequency() string {
	return i.IssueFrequency
}

// MagazinesEqual == Реалізація методу для порівняння журналів
func (i IrregularMagazine) MagazinesEqual(other store.Magazines) bool {
	if otherMag, ok := other.(IrregularMagazine); ok {
		return i.Title == otherMag.Title &&
			i.Publisher == otherMag.Publisher &&
			i.ReleaseDate == otherMag.ReleaseDate &&
			i.IssueFrequency == otherMag.IssueFrequency
	}
	return false
}
