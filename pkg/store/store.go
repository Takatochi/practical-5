package store

type Magazines interface {
	Periodicals
	MagazineComparer
}

// Periodicals Інтерфейс для представлення журналу
type Periodicals interface {
	GetFrequency() string
	PrintInfo()
}

// MagazineComparer Інтерфейс для порівняння журналів
type MagazineComparer interface {
	MagazinesEqual(other Magazines) bool
	//MagazinesEqual(Magazines) bool
	//MagazineLess() bool
	//MagazineGreater() bool
}
