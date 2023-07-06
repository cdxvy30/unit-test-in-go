package main

import (
	"fmt"
	"jubo-unit-test/golden-test/books"
	"jubo-unit-test/golden-test/report"
)

func main() {
	report := report.Generate(books.Books)
	fmt.Println(report)
}
