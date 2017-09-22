package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books
	var Book2 Books

	Book1.title = "Go语言"
	Book1.author = "golang"
	Book1.subject = "高性能Go语言"
	Book1.book_id = 1

	Book2.title = "2Go语言"
	Book2.author = "2golang"
	Book2.subject = "2高性能Go语言"
	Book2.book_id = 2

	printBook(&Book1)
	printBook(&Book2)

}

func printBook(book *Books) {
	fmt.Printf("Book title: %s\n", book.title)
	fmt.Printf("Book author: %s\n", book.author)
	fmt.Printf("Book subject: %s\n", book.subject)
	fmt.Printf("Book book_id: %d\n", book.book_id)
}
