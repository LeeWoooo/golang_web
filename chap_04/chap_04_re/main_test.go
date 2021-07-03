package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileServer(t *testing.T) {
	assert := assert.New(t)
	path := "/Users/leewoooo/Pictures/10-8.jpg"
	t.Run("File Upload Test", func(t *testing.T) {
		file, err := os.Open(path)
		assert.NoError(err)
		defer file.Close()

		// prev file Remove
		os.RemoveAll("./uploads")

		//make formfile
		buf := &bytes.Buffer{}
		writer := multipart.NewWriter(buf)
		multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))
		assert.NoError(err)

		io.Copy(multi, file)
		writer.Close()

		//req,res
		resp := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/uploads", buf)
		req.Header.Set("Content-type", writer.FormDataContentType())

		uploadFileHandler(resp, req)
		assert.Equal(http.StatusOK, resp.Code)

		uploadFilePath := "./uploads/" + filepath.Base(path)
		_, err = os.Open(uploadFilePath)
		assert.NoError(err)

		uploadFile, _ := os.Open(uploadFilePath)
		defer uploadFile.Close()
		originFile, _ := os.Open(path)
		defer originFile.Close()

		upload := []byte{}
		origin := []byte{}
		originFile.Read(origin)
		uploadFile.Read(upload)

		assert.Equal(origin, upload)
	})
}
