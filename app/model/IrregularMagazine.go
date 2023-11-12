package model

import mag "Pz5/pkg/magazine"

type IrregularMagazine struct {
	*Magazines
}

// NewIrregularMagazine Конструктор для створення нового журналу, що виходить раз на кілька місяців
func NewIrregularMagazine(title, publisher, releaseDate string) IrregularMagazine {
	return IrregularMagazine{
		Magazines: NewMagazine(title, publisher, releaseDate, "Раз на кілька місяців"),
	}
}

// GetFrequency Метод для отримання частоти виходу журналу, що виходить раз на кілька місяців
func (i IrregularMagazine) GetFrequency() string {
	return i.Magazines.GetFrequency()
}

// MagazinesEqual == Реалізація методу для порівняння журналів
func (i IrregularMagazine) MagazinesEqual(other mag.Magazines) bool {
	if otherMag, ok := other.(IrregularMagazine); ok {
		return i.Description.Title == otherMag.Description.Title &&
			i.Description.Publisher == otherMag.Description.Publisher &&
			i.Description.ReleaseDate == otherMag.Description.ReleaseDate &&
			i.IssueFrequency == otherMag.IssueFrequency
	}
	return false
}
