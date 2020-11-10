package db

type Manga struct {
	Id       int    `xorm:"int auto_increment 'id'"`
	MangaId  int    `xorm:"int 'mangaid'"`
	Source   string `xorm:"varchar not null 'source'"`
	Author   string `xorm:"varchar not null 'author'"`
	Artist   string `xorm:"varchar not null 'artist'"`
	Language string `xorm:"varchar not null 'language'"`
}

type Chapter struct {
	Id    int    `xorm:"int auto_increment 'id'"`
	Name  string `xorm:"varchar 'name'"`
	Read  bool   `xorm:"bit 'read'"`
	Manga `xorm:"extends"`
}

type TLGroup struct {
}
