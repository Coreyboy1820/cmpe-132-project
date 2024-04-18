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
		log.Println(err)
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
		return
	}

	// make sure the studentId was a valid student Id
	if len(users) == 0 {
		userpkg.CurrUser = userpkg.User{}
		log.Print("login unsuccessful")
		// TODO: throw some error here that login was unsuccesful
		return
	}
	
	// make sure the student Id isn't duplicated
	if(len(users) > 1) {
		userpkg.CurrUser = userpkg.User{}
		log.Print("Unique identifier did not hold fix the database")
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
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	hashedPassword := crypto.HashPassword(Credentials.Password+userpkg.CurrUser.Salt)
	if(hashedPassword != userpkg.CurrUser.PasswordHash) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	
	err = UpdateLogin(true)
	if(err != nil) {
		log.Print(err)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	err := UpdateLogin(false)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/", http.StatusFound)
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
