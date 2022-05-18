package main

import (
	//"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello World from GO HTTP Server </h1>\n\n <h2>Deploying Server version v3 on Cloud RUN </h2>\n"))
	})
	http.ListenAndServe(":80", nil)
}
