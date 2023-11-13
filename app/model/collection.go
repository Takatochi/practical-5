package model

import (
	mag "Pz5/pkg/magazine"
	"fmt"
)

// Collection Структура для представлення колекції журналів
type Collection struct {
	IrregularMagazine []IrregularMagazine `json:"irregularMagazine"`
	MonthlyMagazine   []MonthlyMagazine   `json:"monthlyMagazine"`
	magazines         []mag.Magazines
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
func (c *Collection) AddDefaultMagazines() []mag.Magazines {
	var result []mag.Magazines
	for _, magazine := range c.IrregularMagazine {
		result = append(result, magazine)
	}
	for _, magazine := range c.MonthlyMagazine {
		result = append(result, magazine)
	}
	return result
}

// PrintInfo Метод для виведення інформації про журнал
func (m Collection) PrintInfo() {
	fmt.Println("MonthlyMagazine:")
	for _, mag := range m.MonthlyMagazine {
		fmt.Printf("Назва: %s\nВидавець: %s\nДата виходу: %s\n", mag.Description.Title, mag.Description.Publisher, mag.Description.ReleaseDate)
		fmt.Printf("Частота виходу: %s\n", mag.GetFrequency())
		fmt.Println()
	}
	fmt.Println("IrregularMagazine:")
	for _, mag := range m.IrregularMagazine {
		fmt.Printf("Назва: %s\nВидавець: %s\nДата виходу: %s\n", mag.Description.Title, mag.Description.Publisher, mag.Description.ReleaseDate)
		fmt.Printf("Частота виходу: %s\n", mag.GetFrequency())
		fmt.Println()
	}
}

// GetFrequency Метод для отримання частоти виходу журналу, що виходить раз на кілька місяців

//
