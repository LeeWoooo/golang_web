package app

import (
	"bytes"
	"encoding/json"
	"golang_web/chap_05/dto"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	// create mock server
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, err := ioutil.ReadAll(resp.Body)
	assert.NoError(err)
	assert.Equal("Hello Go:)", string(data))
}

func TestFindOneUserByID(t *testing.T) {
	//prepared
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	//mock
	//saveUser
	mockUser := &User{
		ID:   0,
		Name: "leewoooo",
		Age:  26,
	}
	UsersMap[0] = mockUser

	//when
	resp, err := http.Get(ts.URL + "/users/0")
	assert.NoError(err)
	defer resp.Body.Close()

	//then
	assert.Equal(http.StatusOK, resp.StatusCode)

	selected := new(User)
	err = json.NewDecoder(resp.Body).Decode(selected)
	assert.NoError(err)

	if assert.NotNil(selected) {
		assert.Equal(0, selected.ID)
		assert.Equal("leewoooo", selected.Name)
		assert.Equal(26, selected.Age)
	}
}

func TestFindListUser(t *testing.T) {
	//prepared
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	//mock
	//saveUser
	mockUser := []*User{
		{
			ID:   0,
			Name: "leewoooo",
			Age:  26,
		},
		{
			ID:   1,
			Name: "leewooo",
			Age:  26,
		},
		{
			ID:   2,
			Name: "leewoo",
			Age:  26,
		},
	}
	for i, u := range mockUser {
		UsersMap[i] = u
	}

	//when
	resp, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	defer resp.Body.Close()

	//then
	assert.Equal(http.StatusOK, resp.StatusCode)

	var userList []*User
	assert.NoError(err)
	json.NewDecoder(resp.Body).Decode(&userList)
	if assert.NotNil(userList) {
		assert.Equal(3, len(userList))

		jsonData, err := json.MarshalIndent(userList, "", " ")
		assert.NoError(err)

		t.Log(string(jsonData))
	}

}

func TestSaveUser(t *testing.T) {
	//prepared
	assert := assert.New(t)

	//given
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	user := &User{
		Name: "leewooo",
		Age:  26,
	}
	jsonData, err := json.Marshal(user)
	assert.NoError(err)

	//when
	resp, err := http.Post(ts.URL+"/users", "application/json", bytes.NewBuffer(jsonData))
	assert.NoError(err)
	defer resp.Body.Close()

	//then
	assert.Equal(http.StatusCreated, resp.StatusCode)

	saved := new(User)
	err = json.NewDecoder(resp.Body).Decode(saved)
	assert.NoError(err)

	assert.Equal("leewooo", saved.Name)
	assert.Equal(26, saved.Age)
}

func TestUpdateUser(t *testing.T) {
	//prepared
	assert := assert.New(t)
	mux := NewHandler()

	//saveUser
	mockUser := &User{
		ID:   0,
		Name: "leewoooo",
		Age:  26, // will be 27
	}
	UsersMap[0] = mockUser

	//updateUser
	updateRequest := dto.UpdateUserRequestDTO{
		Age: 27,
	}
	jsonByte, err := json.Marshal(updateRequest)
	assert.NoError(err)

	//givne
	resp := httptest.NewRecorder()
	req := httptest.NewRequest("PUT", "/users/0", bytes.NewBuffer(jsonByte))

	//when
	mux.ServeHTTP(resp, req)

	//then
	assert.Equal(http.StatusOK, resp.Code)
	updated := UsersMap[0]
	assert.Equal(27, updated.Age)
}

func TestDeleteUserTest(t *testing.T) {
	//prepared
	assert := assert.New(t)
	mux := NewHandler()

	//saveUser
	mockUser := &User{
		ID:   0,
		Name: "leewoooo",
		Age:  26, // will be 27
	}
	UsersMap[0] = mockUser

	//given
	resp := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/users/0", nil)

	//when
	mux.ServeHTTP(resp, req)

	//then
	assert.Equal(http.StatusOK, resp.Code)

	_, exist := UsersMap[0]
	assert.False(exist)
}

func TestUserList(t *testing.T) {
	//prepared
	assert := assert.New(t)

	//given
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	//when
	resp, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	defer resp.Body.Close()

	//then
	assert.Equal(http.StatusOK, resp.StatusCode)
	jsonByte, err := json.MarshalIndent(resp.Body, "", " ")
	assert.NoError(err)
	t.Log(string(jsonByte))
}
