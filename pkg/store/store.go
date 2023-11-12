package store

import (
	"Pz5/app/model"
	magazzine "Pz5/pkg/magazine"
)

// MagazineRepository ...
type MagazineRepository interface {
	Save(df []magazzine.Magazines) error
	Load() ([]model.Magazines, error)
}

// Store ...
type Store interface {
	Collection() MagazineRepository
}
