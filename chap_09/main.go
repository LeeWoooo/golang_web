package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render

var users []User

// User for http study
type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"crated_at"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:  "leewoooo",
		Email: "leecoding2285@gmail.com",
		Age:   26,
	}

	rd.JSON(w, http.StatusOK, user)
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		rd.Text(w, http.StatusBadRequest, "Bad Request")
	}
	user.CreatedAt = time.Now()

	users = append(users, *user)

	rd.JSON(w, http.StatusOK, user)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	//html을 등록할 때 확장자를 빼고 등록한다.
	//render package에서 default 확장자는 tmpl
	//template을 찾는 경로는 templates이다.
	rd.HTML(w, http.StatusOK, "hello", nil)
}

func main() {
	// option을 추가할 수 있다.
	rd = render.New(render.Options{
		Directory: "./templates/template", //default값은 ./templates
		Extensions: []string{
			".html", ".tmpl",
		},
	})
	mux := pat.New()

	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)
	mux.Get("/hello", helloHandler)

	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
