package collector

import "GoManga/db"

type Collector interface {
	GetName() string
	GetTitle(id int) (db.Manga, error)
	GetChapters(db.Manga) ([]Chapter, error)
}

type Chapter struct {
	TLGroup  string
	Language string
	Volume   string
	Chapter  string
	Pages    []string
}
