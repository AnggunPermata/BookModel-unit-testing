package model

type BookModel interface {
	Get() []Book
	Insert(Book) error
}
