package main

import (
	"log"
	"net/http"
)

// define server struct
type server struct{
	addr string
}

// implement ServeHTTP method for server struct
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle requests here
	w.Write([]byte("hello from server"))
}

// use server struct in main function
func main() {
	s := &server{addr: ":8080"}
 if err := 	http.ListenAndServe(s.addr, s); err != nil {
	log.Fatal(err)
 }

}