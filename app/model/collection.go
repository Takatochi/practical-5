package model

import (
	mag "Pz5/pkg/magazine"
)

// Collection Структура для представлення колекції журналів
type Collection struct {
	Magazines
	IrregularMagazine []IrregularMagazine `json:"irregularMagazine"`
	MonthlyMagazine   []MonthlyMagazine   `json:"monthlyMagazine"`
}

func NewCollectionFromMagazines(magazines ...mag.Magazines) *Collection {
	var coll Collection
	coll.AddMagazines(magazines...)
	return &coll
}

func (c *Collection) AddMagazines(magazines ...mag.Magazines) []mag.Magazines {
	var result []mag.Magazines
	for _, magazine := range magazines {
		switch mag := magazine.(type) {
		case IrregularMagazine:
			c.IrregularMagazine = append(c.IrregularMagazine, mag)
			result = append(result, mag)
		case MonthlyMagazine:
			c.MonthlyMagazine = append(c.MonthlyMagazine, mag)
			result = append(result, mag)
			// Додайте інші типи журналів, які ви хочете підтримувати
		}
	}
	return result
}

//
