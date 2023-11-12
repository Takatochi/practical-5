package dataJson

import (
	"Pz5/pkg/store"
	"fmt"
)

type Store struct {
	fileName           string
	magazineRepository *MagazineRepository
}

func New(db string) *Store {
	return &Store{
		fileName: db,
	}
}

// Collection interface for add parameters Store.Collection()....
func (s *Store) Collection() store.MagazineRepository {
	fmt.Println("save")
	if s.magazineRepository != nil {
		return s.magazineRepository
	}
	//s.db.Begin()
	s.magazineRepository = &MagazineRepository{
		store: s,
	}

	return s.magazineRepository
}
