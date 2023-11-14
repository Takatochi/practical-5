package model

// Magazines Інтерфейспредставляє комбінацію інтерфейсів Periodicals та MagazineComparer.
// Magazines interface represents the combination of Periodicals and MagazineComparer interfaces.
type Magazines interface {
	Periodicals
	MagazineComparer
}

// Periodicals Інтерфейс для представлення журналу
// Periodicals interface represents the basic functionalities of a magazine.
type Periodicals interface {
	GetFrequency() string
	PrintInfo()

	//SaveUsers(string) error
}

// MagazineComparer Інтерфейс для порівняння журналів
// MagazineComparer interface provides methods for comparing magazines.
type MagazineComparer interface {
	MagazinesEqual(Magazines) bool
	//MagazinesEqual(Magazines) bool
	//MagazineLess() bool
	//MagazineGreater() bool
}
