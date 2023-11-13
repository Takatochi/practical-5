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

// SaveJSON функція для збереження користувачів у JSON-файл
func (m *MagazineRepository) SaveJSON(df []magazzine.Magazines) error {
	collection := model.NewCollectionFromMagazines(df...)
	data, err := json.MarshalIndent(collection, " ", "	")
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

// LoadJSON функція для завантаження користувачів з JSON-файлу
func (m *MagazineRepository) LoadJSON() (*model.Collection, error) {

	data, err := ioutil.ReadFile(m.store.fileName)
	if err != nil {
		return nil, err
	}

	var df model.Collection
	err = json.Unmarshal(data, &df)

	if err != nil {
		return nil, err
	}
	fmt.Printf("magazines load %s\n", m.store.fileName)
	return &df, nil
}
