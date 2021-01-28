package main

import "fmt"

type Aggregate interface {
	Iterator() Iterator
}

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type BookShelf struct {
	Books []*Book
}

func (self *BookShelf) Add(book *Book) {
	self.Books = append(self.Books, book)
}

func (self *BookShelf) Iterator() Iterator {
	return &BookShelfIterator{BookShelf: self}
}

type BookShelfIterator struct {
	BookShelf *BookShelf
	last      int
}

func (self *BookShelfIterator) HasNext() bool {
	if self.last >= len(self.BookShelf.Books) {
		return false
	}
	return true
}

func (self *BookShelfIterator) Next() interface{} {
	book := self.BookShelf.Books[self.last]
	self.last++
	return book
}

type Book struct {
	name string
}

func main() {
	bookShelf := &BookShelf{}

	asserts := []string{"A", "B", "C"}
	for _, assert := range asserts {
		bookShelf.Add(&Book{assert})
	}

	i := 0
	for iterator := bookShelf.Iterator(); iterator.HasNext(); {
		if book := iterator.Next(); book.(*Book).name != asserts[i] {
			fmt.Println(asserts[i])
			fmt.Println(book.(*Book).name)
		}
		i++
	}
}
