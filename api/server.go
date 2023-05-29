package api

import (
	"fmt"
	"net/http"
)

func Server() {
	http.Handle("/foo", fooHandler{})
	http.ListenAndServe(":8000", nil)
}

type fooHandler struct{}

func (h fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GO API says Hi")
	fmt.Println("server stuff")
}

