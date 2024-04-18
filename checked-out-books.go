package main

import (
	"cmpe-132-project/dbQuerries"
	"cmpe-132-project/dbutil"
	userpkg "cmpe-132-project/user"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func HandleCheckedOutBooks(w http.ResponseWriter, r *http.Request) {
	// check for correct perms
	if !userpkg.CurrUser.CheckoutBook {
		log.Print("You do not have permissions for this")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	tm := make(map[string]interface{})
	tm["CurrUser"] = userpkg.CurrUser

	// --------- Read all books ---------

	booksInCart, err := dbQuerries.BooksInCart{}.Read("WHERE userId=?", userpkg.CurrUser.UserId)
	if err != nil {
		log.Print(err)
		return
	}
	tm["BooksInCartCount"] = len(booksInCart)

	// Read all books checked out

	checkedOutBooks, err := dbQuerries.BooksAndCheckedOut{}.Read("WHERE userId=?", userpkg.CurrUser.UserId)
	if err != nil {
		log.Print(err)
		return
	}
	tm["CheckedOutBooks"] = checkedOutBooks
	tm["CheckedOutBooksCount"] = len(checkedOutBooks)

	t, err := template.ParseFiles("static/checked-out-books.html", "static/header.html")
	if err != nil {
		panic(err.Error())
	}
	t.Execute(w, tm)
}

// despite its name, reserve a book only works after a professor has already checked that book out, it allows them to select students to check out the book for them
func ReserveBook(w http.ResponseWriter, r *http.Request) {
	// check for correct perms
	if !userpkg.CurrUser.ReserveBooks {
		log.Print("You do not have permissions for this")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Parse the JSON request body
	CheckedOutBook := struct {
		StudentId string `json:"studentId"`
		CheckedOutBookId   int `json:"checkedOutBooksId"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&CheckedOutBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	users, err := userpkg.User{}.Read("WHERE studentId=?", CheckedOutBook.StudentId)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(users) == 0 {
		log.Print("user does not exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	studentToTransferBookTo := users[0]
	// then we update the stock to make sure it can't be further reduced
	updateStmt := "UPDATE checkedOutBooks SET userId=? WHERE checkedOutBooksId=?"
	_, err = dbutil.DB.Exec(updateStmt, studentToTransferBookTo.UserId, CheckedOutBook.CheckedOutBookId)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}


}
