package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Collection Структура для представлення колекції журналів
type Collection struct {
	IrregularMagazine *IrregularMagazine `json:"irregularMagazine"`
	MonthlyMagazine   *MonthlyMagazine   `json:"monthlyMagazine"`
}

// SaveToFile Метод для збереження колекції в файл у форматі JSON
func (c Collection) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Collection saved to %s\n", filename)
	return nil
}

// LoadFromFile Метод для завантаження колекції з файлу у форматі JSON
func LoadFromFile(filename string) (*Collection, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var c Collection
	err = json.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
