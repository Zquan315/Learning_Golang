package main

import f "fmt"

type book struct {
	name   string
	author string
}

func (b book) PrintInfoBook() {
	f.Printf("Book name: %s\nAuthor: %s\n", b.name, b.author)
}
func RunBook() {
	var b book
	b.name = "The Go Programming Language"
	b.author = "Alan A. A. Donovan and Brian W. Kernighan"
	b.PrintInfoBook()
}

type square struct {
	side float64
}

func (s square) Area() float64 {
	return s.side * s.side
}

func RunSquare() {
	var s square
	f.Print("Enter the side length of the square:")
	f.Scan(&s.side)
	f.Printf("Area of square with side %.2f is %.2f\n", s.side, s.Area())
}
