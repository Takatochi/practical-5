package model

type IrregularMagazine struct {
	magazine
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
