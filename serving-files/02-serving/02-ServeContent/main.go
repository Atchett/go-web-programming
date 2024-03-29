package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
		<img src="toby.jpg" />
	`)

}

func dogPic(w http.ResponseWriter, r *http.Request) {

	f, err := os.Open("toby.jpg")
	if err != nil {
		// can also use a constant
		http.Error(w, "File not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	http.ServeContent(w, r, f.Name(), fi.ModTime(), f)

}
