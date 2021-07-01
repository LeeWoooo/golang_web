package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type fooHandler struct{}

// User user
// add json annotation
type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"create_at"`
}

// implement Handler ServeHTTP
func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	// json to struct
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request :(")
		return
	}

	user.CreatedAt = time.Now()
	// struct to json
	jsonByte, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonByte))
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// get query string parameter
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "world"
		}

		fmt.Fprintf(w, "Hello Golnag %s:)", name)
	})

	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello bar:)")
	})

	// add handler instance
	mux.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":3000", mux)
}
