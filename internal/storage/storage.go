package storage

type Storage interface {
	Writer
	Reader
}

type Writer interface {
	SaveUrl(url, alice string) error
}

type Reader interface {
	GetUrl(alice string) (string, error)
}
