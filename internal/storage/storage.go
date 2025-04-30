package storage

type Storage interface {
	Writer
}

type Writer interface {
	SaveUrl(url, alice string) error
}
