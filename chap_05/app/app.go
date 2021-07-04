package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang_web/chap_05/dto"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// User information user
type User struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// UsersMap UserList
var UsersMap map[int]*User

// ID user ID
var ID int

func indexHander(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Go:)")
}

func findUserListHandler(w http.ResponseWriter, r *http.Request) {
	var userList []*User
	for _, value := range UsersMap {
		userList = append(userList, value)
	}

	jsonData, err := json.Marshal(userList)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonData))
}

func findOneUserByID(w http.ResponseWriter, r *http.Request) {
	//get id at pathValiable
	taregetID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
	}

	//validation
	_, exist := UsersMap[taregetID]
	if !exist {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "user isn't exist")
		return
	}

	// Find
	selected := UsersMap[taregetID]

	//Marshal
	jsonData, err := json.Marshal(selected)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	//response
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, bytes.NewBuffer(jsonData))
}

func saveUserHandler(w http.ResponseWriter, r *http.Request) {
	// user struct decode
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	// save
	user.ID = ID
	UsersMap[ID] = user
	ID++

	// user json marshal
	jsonData, err := json.Marshal(user)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	// response
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(jsonData))
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	// get id at pathvaliable
	vars := mux.Vars(r)
	targetID, err := strconv.Atoi(vars["id"])
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	// validation?
	target, isExist := UsersMap[targetID]
	if !isExist {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "User isn't exist")
		return
	}

	// updatereqeust decode
	ur := new(dto.UpdateUserRequestDTO)
	err = json.NewDecoder(r.Body).Decode(ur)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	// update
	updateUser := User{
		ID:   targetID,
		Name: target.Name,
		Age:  ur.Age,
	}

	//response
	UsersMap[targetID] = &updateUser
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Update success")
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	//get id at pathValiable
	taregetID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
	}

	//validation
	_, exist := UsersMap[taregetID]
	if !exist {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "user isn't exist")
		return
	}

	//delete user
	delete(UsersMap, taregetID)

	//response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Delete Success")
}

// NewHandler Make http Handler Instance (REST)
func NewHandler() http.Handler {
	UsersMap = make(map[int]*User)
	ID = 0
	mux := mux.NewRouter()
	mux.HandleFunc("/", indexHander)
	mux.HandleFunc("/users", findUserListHandler).Methods("GET")
	mux.HandleFunc("/users/{id:[0-9]+}", findOneUserByID).Methods("GET")
	mux.HandleFunc("/users", saveUserHandler).Methods("POST")
	mux.HandleFunc("/users/{id:[0-9]+}", updateUserHandler).Methods("PUT")
	mux.HandleFunc("/users/{id:[0-9]+}", deleteUserHandler).Methods("DELETE")
	return mux
}
