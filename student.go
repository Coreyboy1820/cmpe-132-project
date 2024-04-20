package main

import (
	userpkg "cmpe-132-project/user"
	_ "database/sql"
	"encoding/json"
	"net/http"
	"log"

	_ "golang.org/x/crypto/bcrypt"
	_ "github.com/mattn/go-sqlite3"
)

/// CheckoutBook will make sure the user has the correct perms to checkout a book and will the check it out

func CheckoutBook(w http.ResponseWriter, r *http.Request) {
	if(!userpkg.CurrUser.LoggedIn) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if(!userpkg.CurrUser.CheckoutBook) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Parse the JSON request body
	book := struct {
		BookId string `json:"bookId"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Print(book)
}