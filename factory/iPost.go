package factory

type IPost interface {
	GetTitle() string
	GetAuthor() string
	SetTitle(title string)
	SetAuthor(author string)
}
