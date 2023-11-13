package store

import (
	"Pz5/app/model"
	magazzine "Pz5/pkg/magazine"
)

// MagazineRepository ...
type MagazineRepository interface {
	SaveJSON(df []magazzine.Magazines) error
	LoadJSON() (*model.Collection, error)
}

// Store ...
type Store interface {
	Collection() MagazineRepository
}
