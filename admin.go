package main

import (
	"cmpe-132-project/dbutil"
	"cmpe-132-project/dbQuerries"
	userpkg "cmpe-132-project/user"
	_"crypto/rand"
	_ "database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)


func HandleAdmin(w http.ResponseWriter, r *http.Request) {
	// check for correct perms
	if !userpkg.CurrUser.CreateUser {
		errString := "You do not have permissions for this"
		log.Print(errString)
		http.Error(w, errString, http.StatusUnauthorized)
		return
	}
	tm := make(map[string]interface{})
	tm["CurrUser"] = userpkg.CurrUser

	// --------- Read all books ---------

	booksInCart, err := dbQuerries.BooksInCart{}.Read("WHERE userId=?", userpkg.CurrUser.UserId)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tm["BooksInCartCount"] = len(booksInCart)

	// --------- Read all users ---------

	users, err := userpkg.User{}.Read("ORDER BY lastName, firstName, studentId", userpkg.CurrUser.UserId)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tm["Users"] = users

	t, err := template.ParseFiles("static/admin.html", "static/header.html")
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err.Error())
	}
	t.Execute(w, tm)
}

func UpdateRole(w http.ResponseWriter, r *http.Request) {
	UserToUpdate := struct {
		UserId int `json:"userId"`
		Role string `json:"role"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&UserToUpdate)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if userpkg.CurrUser.UserId == UserToUpdate.UserId {
		var users []userpkg.User
		users, err = userpkg.User{}.Read("WHERE userId=?", UserToUpdate.UserId)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(users) == 0 {
			errString := "User Does Not Exist"
			log.Print(errString)
			http.Error(w, errString, http.StatusInternalServerError)
			return
		}
		userpkg.CurrUser = users[0]
	}
	if !userpkg.CurrUser.UpdateUser {
		errString := "You do not have permissions for this"
		http.Error(w, errString, http.StatusUnauthorized)
		log.Print(errString)
		return
	}

	roles, err := dbQuerries.Roles{}.Read("WHERE roleName=?", UserToUpdate.Role)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(roles) == 0 {
		errString := "Role does not exist"
		log.Print(errString)
		http.Error(w, errString, http.StatusInternalServerError)
		return
	}
	role := roles[0]
	// then we update the stock to make sure it can't be further reduced
	updateStmt := "UPDATE users SET roleId=? WHERE userId=?"
	_, err = dbutil.DB.Exec(updateStmt, role.RoleId, UserToUpdate.UserId)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/", http.StatusSeeOther)
}