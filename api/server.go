package api

import (
	"fmt"
	"net/http"
)
const PORT string = ":8000"

func Server() {
	fmt.Printf("Server listening on port %v/foo", PORT)
	http.Handle("/foo", fooHandler{})
	http.Handle("/mockData", mockJsonResponse{})
	http.ListenAndServe(PORT, nil)

}

type fooHandler struct{}
type mockJsonResponse struct{}

func (h fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GO API says Hi")
	fmt.Println("server stuff")
}

func (h mockJsonResponse) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	byteData := []byte(`[{"name":"frank"}]`)
	w.Write(byteData)
}
