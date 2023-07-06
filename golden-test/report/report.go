package report

import (
	"bytes"
	"log"
	"text/template"

	"jubo-unit-test/golden-test/books"
)

const (
	header string = `
| Title         | Author        | Pages  | ISBN  | Price  |
| ------------- | ------------- | ------ | ----- | ------ |
`
	rowTemplate string = "| {{ .Title }} | {{ .Author }} | {{ .Pages }} | {{ .ISBN }} | {{ .Price }} |\n"
)

func Generate(books []books.Book) string {
	buf := bytes.NewBufferString(header)

	t := template.Must(template.New("table").Parse(rowTemplate))

	for _, book := range books {
		err := t.Execute(buf, book)
		if err != nil {
			log.Println("Error executing template:", err)
		}
	}

	return buf.String()
}
