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

	// template에서 반복문을 사용할 때는 range를 이용
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

	//template을 가져올 때는 parseFeils에 파일경로를 넣어준다.
	temp3, err := template.New("prac template file").ParseFiles("templates/temp1.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	var buffer3 bytes.Buffer
	//가져온 template를 excute를 할 때는 Excute가 아닌 ExcuteTemplate
	temp3.ExecuteTemplate(&buffer3, "temp1.tmpl", newUser)
	log.Println(buffer3.String())
}
