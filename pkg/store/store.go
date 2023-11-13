package store

import (
	mag "Pz5/app/model"
	"Pz5/app/model/magazine"
)

// MagazineRepository ...
type MagazineRepository interface {
	SaveJSON(df []mag.Magazines) error
	LoadJSON() (*magazine.Collection, error)
}

// Store ...
type Store interface {
	Collection() MagazineRepository
}
