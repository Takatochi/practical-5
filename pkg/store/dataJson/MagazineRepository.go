package dataJson

import (
	magazzine "Pz5/app/model"
	"Pz5/app/model/magazine"
	"encoding/json"
	"fmt"
	"os"
)

type MagazineRepository struct {
	store *Store
}

// SaveJSON функція для збереження користувачів у JSON-файл
func (m *MagazineRepository) SaveJSON(df []magazzine.Magazines) error {
	collection := magazine.NewCollectionFromMagazines(df...)
	data, err := json.MarshalIndent(collection, " ", "	")
	if err != nil {
		return err
	}

	err = os.WriteFile(m.store.fileName, data, 0644)
	if err != nil {
		return err
	}
	fmt.Printf("magazines saved to %s\n", m.store.fileName)
	return nil
}

// LoadJSON функція для завантаження користувачів з JSON-файлу
func (m *MagazineRepository) LoadJSON() (*magazine.Collection, error) {

	data, err := os.ReadFile(m.store.fileName)
	if err != nil {
		return nil, err
	}

	var df magazine.Collection
	err = json.Unmarshal(data, &df)

	if err != nil {
		return nil, err
	}
	fmt.Printf("magazines load %s\n", m.store.fileName)
	return &df, nil
}
