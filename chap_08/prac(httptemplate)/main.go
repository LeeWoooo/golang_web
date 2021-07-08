package main

import (
	"bytes"
	"html/template"
	"log"
)

// User for httpTemplate
type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	temp, err := template.New("prac").
		Parse(`
		Name:{{.Name}}
		Email:{{.Email}}
		Age:{{.Age}}`)

	if err != nil {
		log.Fatal(err)
	}

	newUser := User{
		Name:  "1",
		Email: "1",
		Age:   1,
	}

	var buffer bytes.Buffer
	err = temp.Execute(&buffer, newUser)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(buffer.String())

	temp2, err := template.New("prac range").
		Parse(`
		{{range . -}}
		Name:{{.Name}}
		Email:{{.Email}}
		Age:{{.Age}}
		{{end}}`)

	users := []User{
		{
			Name:  "1",
			Email: "1",
			Age:   1,
		},
		{
			Name:  "2",
			Email: "2",
			Age:   2,
		},
	}

	var buffer2 bytes.Buffer
	temp2.Execute(&buffer2, users)
	log.Println(buffer2.String())
}
