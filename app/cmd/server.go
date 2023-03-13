package main

import (
	"GoTestingAdvanced/app"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", &app.Server{})
}
