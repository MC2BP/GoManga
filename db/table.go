package db

type Manga struct {
	Id          int64  `xorm:"serial 'id'"`
	MangaId     int    `xorm:"int 'mangaid'"`
	Name        string `xorm:"varchar not null 'name'"`
	Description string `xorm:"varchar not null 'description'"`
	Source      string `xorm:"varchar not null 'source'"`
	Language    string `xorm:"varchar not null 'language'"`
}

type Chapter struct {
	Id      int64  `xorm:"serial 'id'"`
	Name    string `xorm:"varchar 'name'"`
	Read    bool   `xorm:"bit 'read'"`
	MangaId int64  `xorm:"index"`
	TlGroup string `xorm:"varchar 'tlgroup'"`
}
