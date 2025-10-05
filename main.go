package main

import (
	"net/http"
)

// define server struct
// type api struct {
// 	addr string
// }

// implement ServeHTTP method for server struct
// func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	// handle requests here
// 	switch r.Method {
// 	case http.MethodGet:
// 		switch r.URL.Path {
// 		case "/":
// 			w.Write([]byte("index page"))
// 			return
// 		case "/users":
// 			w.Write([]byte("user page"))
// 			return

// 		default:
// 			w.Write([]byte("404 not found"))
// 			return
// 		}
// 	}
// }

// use server struct in main function
func main() {
	api := &api{addr: ":8080"}
	
	// initialize Mux
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr: api.addr,
		Handler : mux,
	}

	mux.HandleFunc("Get /users", api.getUsersHandler)
	mux.HandleFunc("Post /users", api.createUsersHandler)

	srv.ListenAndServe()
}
