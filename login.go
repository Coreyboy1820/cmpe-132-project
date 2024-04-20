package main

import (
	userpkg "cmpe-132-project/user"
	"cmpe-132-project/crypto"
	"cmpe-132-project/dbutil"
	_ "database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	_ "golang.org/x/crypto/bcrypt"
	_ "github.com/mattn/go-sqlite3"
)

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/logIn.html", "static/header.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err.Error())
	}
	tm := make(map[string]interface{})
	t.Execute(w, tm)
}

func Login(w http.ResponseWriter, r *http.Request) {
	
	// Parse the JSON request body
	Credentials := struct {
		StudentId string `json:"studentId"`
		Password string  `json:"password"`
		Email    string  `json:"email"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&Credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	whereStmt := "WHERE studentId=?"
	users, err := userpkg.User{}.Read(whereStmt, Credentials.StudentId);
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// make sure the studentId was a valid student Id
	if len(users) == 0 {
		errString := "login unsuccessful"
		log.Print(errString)
		http.Error(w, errString, http.StatusUnauthorized)
		userpkg.CurrUser = userpkg.User{}
		return
	}

	userpkg.CurrUser = users[0]
	// if the password hasn't been set yet, then we redirect to setting up the password
	if(!userpkg.CurrUser.PasswordSet) {
		w.WriteHeader(http.StatusFound)
		return
	}

	// now we can check the password and email to make sure they line up with the studentId
	if(userpkg.CurrUser.Email != Credentials.Email) {
		http.Error(w, "Email is incorrect", http.StatusUnauthorized)
		return
	}

	hashedPassword := crypto.HashPassword(Credentials.Password+userpkg.CurrUser.Salt)
	if(hashedPassword != userpkg.CurrUser.PasswordHash) {
		http.Error(w, "Password is incorrect", http.StatusUnauthorized)
		return
	}
	
	err = UpdateLogin(true)
	if(err != nil) {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	err := UpdateLogin(false)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func UpdateLogin(status bool) (err error) {
	updateStmt := "UPDATE users SET loggedIn=? WHERE studentId=?"
	_, err = dbutil.DB.Exec(updateStmt, status, userpkg.CurrUser.StudentId)
	if err != nil {
		log.Print(err)
		return
	}
	userpkg.CurrUser.LoggedIn = status
	return
}
