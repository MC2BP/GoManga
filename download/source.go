package download

type Title struct {
	Id          int
	Name        string
	Description string
	Author      string
	Artist      string
	Language    string
}

type Source interface {
	GetTitle(int)
}
