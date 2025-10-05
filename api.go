package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

// define server struct
type api struct {
	addr string
}

var users = []User{}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")


	// encode user slice to json 
	err := json.NewEncoder(w).Encode(users)
	if err != nil{ 
		http.Error(w,err.Error(), http.StatusInternalServerError)
	}


	w.WriteHeader(http.statusOK)
}

func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request){
	// DECODE REQUEST BODY TO USER  STRUCT 
	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u := User{
		FirstName: payload.FirstName,
		LastName: payload.LastName,
	}


if err = inserUser(u); err != nil{
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
}

w.WriteHeader(http.StatusAccepted)

}




func inserUser(u User)error{
	if u.FirstName == ""{
		return errors.New("First Name is Required")
	}

	if u.LastName == ""{
		return errors.New("Last Name is Required")
	}


	// storage validation

	for _, user := range users{
		if user.FirstName == u.FirstName && user.LastName == u.Lastname { 
			return errors.New("user already exists")
		} 
	}

	users = append(users, u)
	return u

}


