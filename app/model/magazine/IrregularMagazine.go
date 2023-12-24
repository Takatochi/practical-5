package magazine

import (
	"Pz5/app/model"
)

type IrregularMagazine struct {
	*Magazine `json:"magazine"`
}

// NewIrregularMagazine Конструктор для створення нового журналу, що виходить раз на кілька місяців
func NewIrregularMagazine(title, publisher, releaseDate string) IrregularMagazine {
	return IrregularMagazine{
		Magazine: newMagazine(title, publisher, releaseDate, "Раз на кілька місяців"),
	}
}

// GetFrequency Метод для отримання частоти виходу журналу, що виходить раз на кілька місяців
func (i IrregularMagazine) GetFrequency() string {
	return i.Magazine.GetFrequency()
}

// MagazinesEqual == Реалізація методу для порівняння журналів
func (i IrregularMagazine) MagazinesEqual(other model.Magazines) bool {
	if otherMag, ok := other.(IrregularMagazine); ok {
		return i.Description.Title == otherMag.Description.Title &&
			i.Description.Publisher == otherMag.Description.Publisher &&
			i.Description.ReleaseDate == otherMag.Description.ReleaseDate &&
			i.GetFrequency() == otherMag.GetFrequency()
	}
	return false
}
