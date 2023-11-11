package store

type Magazine interface {
	Periodicals
}

// Periodicals Інтерфейс для представлення журналу
type Periodicals interface {
	GetFrequency() string
	PrintInfo()
}
