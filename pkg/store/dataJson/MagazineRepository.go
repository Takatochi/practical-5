package dataJson

import (
	"Pz5/app/model"
	magazzine "Pz5/pkg/magazine"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type MagazineRepository struct {
	store *Store
}

// Save функція для збереження користувачів у JSON-файл
func (m *MagazineRepository) Save(df []magazzine.Magazines) error {
	data, err := json.MarshalIndent(df, " ", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(m.store.fileName, data, 0644)
	if err != nil {
		return err
	}
	fmt.Printf("magazines saved to %s\n", m.store.fileName)
	return nil
}

// Load функція для завантаження користувачів з JSON-файлу
func (m *MagazineRepository) Load() ([]model.Magazines, error) {
	data, err := ioutil.ReadFile(m.store.fileName)
	if err != nil {
		return nil, err
	}

	var df []model.Magazines
	err = json.Unmarshal(data, &df)
	if err != nil {
		return nil, err
	}

	return df, nil
}
