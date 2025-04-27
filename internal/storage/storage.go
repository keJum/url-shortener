package storage

type Storage interface {
	SaveUrl
}

type SaveUrl interface {
	SaveUrl(url, alice string) error
}
