package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("upload_file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer file.Close()

	dirName := "./uploads"
	os.MkdirAll(dirName, 0777) //0777 권한 부여
	filepath := fmt.Sprintf("%s/%s", dirName, header.Filename)
	f, err := os.Create(filepath)
	defer f.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "파일 전송 성공")
}

func main() {
	// fileserver public direct에 있는 파일들을 access 할 수 있는 fileServer
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/upload", uploadHandler)

	http.ListenAndServe(":3000", nil)
}
