package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	html := "<form method=post action=/save enctype=multipart/form-data><input type=file name=upload><input type=submit>"
	h := w.Header()
	h.Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprintf(w, html)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	uploadFile, _, err := r.FormFile("upload")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer uploadFile.Close()

	localFile, err := os.OpenFile("/tmp/a", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer localFile.Close()

	size, err := io.Copy(localFile, uploadFile)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, "ok save size=%d", size)
}

func main() {
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
