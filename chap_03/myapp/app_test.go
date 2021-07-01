package myapp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler_WithOutName(t *testing.T) {
	//given

	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	//when
	mux := NewHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code, "they should be equal", res.Code)

	data, err := ioutil.ReadAll(res.Body)
	assert.NoError(err, "ReadAll has err", err)
	assert.Equal("Hello Golang world:)", string(data))
}

func TestIndexPathHandler_WithName(t *testing.T) {
	//given

	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/?name=leewoooo", nil)

	//when
	mux := NewHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code, "they should be equal", res.Code)

	data, err := ioutil.ReadAll(res.Body)
	assert.NoError(err, "ReadAll has err", err)
	assert.Equal("Hello Golang leewoooo:)", string(data))
}

func TestFooHandler(t *testing.T) {
	t.Run("FooHandler Test(JSON)", func(t *testing.T) {
		//given
		assert := assert.New(t)
		mux := NewHandler()

		user := User{
			FirstName: "lee",
			LastName:  "woogil",
			Email:     "leecoding2285@gmail.com",
		}

		jsonData, err := json.Marshal(user)
		assert.NoError(err, "json marshal has error", err)
		res := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/foo", bytes.NewBuffer(jsonData))

		//when
		mux.ServeHTTP(res, req)

		//then
		resUser := new(User)
		err = json.NewDecoder(res.Body).Decode(resUser)
		assert.NoError(err, "json decode has error", err)

		if assert.Equal(http.StatusOK, res.Code, "they should be eqaul") {
			assert.Equal("lee", resUser.FirstName, "they should be eqaul")
		}
	})
}
