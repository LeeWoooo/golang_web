package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	uploadfile, header, err := r.FormFile("upload_file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "파일을 읽어오지 못했습니다.", err)
		return
	}
	defer uploadfile.Close()

	// 저장할 파일 만들기
	dirName := "./uploads"
	os.Mkdir(dirName, 0777)

	filePath := fmt.Sprintf("%s/%s", dirName, header.Filename)
	// 비어있는 파일을 생성한다.
	file, err := os.Create(filePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	defer file.Close()

	// upload된 파일을 생성할 파일에 복사
	_, err = io.Copy(file, uploadfile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "파일 전송 완료")
}

func main() {
	/*
		특정 폴터안에 있는 정적 파일들을 웹서버에서 클라이언트로 그대로 전달하기 위해 내장기능인 http.FileServer를 사용할 수 있다.
		http.FileServer()는 해당 디렉토리 내의 모든 정적 리소스를 1:1로 매핑하여 그대로 전달하는 일을 한다.
		parameter로는 static file이 들어있는 dir를 받게 된다.
	*/
	http.HandleFunc("/uploads", uploadFileHandler)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.ListenAndServe(":3000", nil)
}
