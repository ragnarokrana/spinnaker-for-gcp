package main

import (
	"io/ioutil"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	htmlContent, err := ioutil.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Error reading HTML file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(htmlContent)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":80", nil)
}
