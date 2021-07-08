package main

import (
	"html/template"
	"log"
	"os"
)

type User struct {
	Name  string
	Email string
	Age   int
}

// template에서 method를 사용할 때 pointer receiver 사용이 안되는걸까?
func (u User) IsOld() bool {
	return u.Age > 20
}

func main() {
	tmpl, err := template.New("Tmpl1").ParseFiles("templates/temp1.tmpl", "templates/temp2.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	// user := User{
	// 	Name:  "leewooo",
	// 	Email: "leecoding2285@gmail.com",
	// 	Age:   26,
	// }

	users := []User{
		{
			Name:  "leewooo",
			Email: "leecoding2285@gmail.com",
			Age:   26,
		},
		{
			Name:  "leewooo2",
			Email: "leecoding2285@gmail.com",
			Age:   26,
		},
	}

	tmpl.ExecuteTemplate(os.Stdout, "temp2.tmpl", users)
}
