package main

import (
	"cmpe-132-project/crypto"
	"cmpe-132-project/dbutil"
	userpkg "cmpe-132-project/user"
	"crypto/rand"
	_ "database/sql"
	"encoding/json"
	"html/template"
	"log"
	"math/big"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func FirstTimeLoginHandler(w http.ResponseWriter, r *http.Request) {

	if userpkg.CurrUser.PasswordSet {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tempPassword, err := rand.Int(rand.Reader, big.NewInt(100000000))
	if err != nil {
		log.Print(err)
		return
	}
	err = HashAndUpdate(tempPassword.String())
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusFound)
	}

	log.Println("Your temporary password is:\n" + tempPassword.String())

	tm := make(map[string]interface{})
	t, err := template.ParseFiles("static/first-time-login.html", "static/header.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err.Error())
	}
	t.Execute(w, tm)
}

func SubmitTempPasword(w http.ResponseWriter, r *http.Request) {
	tempPassword := struct {
		TempPassword string `json:"temporaryPassword"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&tempPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users, err := userpkg.User{}.Read("WHERE studentId=?", userpkg.CurrUser.StudentId)
	if err != nil {
		log.Print(err)
		return
	}
	if len(users) == 0 {
		errString := "student does not exist"
		log.Print(errString)
		http.Error(w, errString, http.StatusBadRequest)
		return
	}
	user := users[0]
	hashedTempPassword := crypto.HashPassword(tempPassword.TempPassword + user.Salt)
	if user.PasswordHash != hashedTempPassword {
		http.Error(w, "Incorrect Password", http.StatusUnauthorized)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusFound)
}

func HandleNewPassword(w http.ResponseWriter, r *http.Request) {
	tm := make(map[string]interface{})
	t, err := template.ParseFiles("static/submit-new-password.html", "static/header.html")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, tm)
}

func SubmitNewPassword(w http.ResponseWriter, r *http.Request) {
	newPassword := struct {
		NewPassword string `json:"newPassword"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&newPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = HashAndUpdate(newPassword.NewPassword)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusFound)
	}
	updateStmt := "UPDATE users SET passwordSet=? WHERE studentId=?"
	_, err = dbutil.DB.Exec(updateStmt, true, userpkg.CurrUser.StudentId)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusFound)
}

func HashAndUpdate(password string) (err error){
	var hashedPassword string
	salt, err := crypto.GenerateSalt(crypto.DefaultSaltLength)
	if err != nil {
		log.Print(err)
		return
	}
	hashedPassword = crypto.HashPassword(password+salt)

	updateStmt := "UPDATE users SET passwordHash=?, salt=? WHERE studentId=?"
	_, err = dbutil.DB.Exec(updateStmt, hashedPassword, salt, userpkg.CurrUser.StudentId)
	if err != nil {
		log.Print(err)
	}
	return
}