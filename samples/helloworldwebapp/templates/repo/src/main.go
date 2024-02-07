package main

import (
  "io"
  "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "<body style='background-color: teal; border: 5px; border-style: dashed; border-color: brown;'><center><h1>This is First Deployment</h1></center></body>")
}

func main() {
  http.HandleFunc("/", hello)
  http.ListenAndServe(":80", nil)
}
