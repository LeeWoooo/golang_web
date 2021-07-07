package main

import (
	"bufio"
	"bytes"
	"golang_web/chap_07/app"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPage(t *testing.T) {
	//prepared
	assert := assert.New(t)
	mux := app.NewHandler()

	//given
	ts := httptest.NewServer(mux)
	defer ts.Close()

	//when
	resp, err := http.Get(ts.URL)

	//then
	assert.NoError(err)
	resp.Body.Close()
	assert.Equal(http.StatusOK, resp.StatusCode)

	bodyData, err := ioutil.ReadAll(resp.Body)
	assert.NoError(err)
	assert.Equal("Hello Go", string(bodyData))
}

func TestDecoratorHandler(t *testing.T) {
	//prepared
	assert := assert.New(t)
	mux := app.NewHandler()

	//given
	ts := httptest.NewServer(mux)
	defer ts.Close()

	buf := &bytes.Buffer{}
	log.SetOutput(buf)

	//when
	resp, err := http.Get(ts.URL)
	resp.Body.Close()
	assert.Equal(http.StatusOK, resp.StatusCode)
	assert.NoError(err)

	r := bufio.NewReader(buf)
	line, _, err := r.ReadLine()
	assert.NoError(err)
	assert.Contains(string(line), "[LOGGER1] Started")
}
