package main

import (
	"html/template"
	"net/http"
	"log"
	"encoding/json"
	"strconv"
    "time"
	userpkg "cmpe-132-project/user"
	"cmpe-132-project/dbQuerries"
	"cmpe-132-project/dbutil"

	_ "github.com/mattn/go-sqlite3"

)

func HandleCart(w http.ResponseWriter, r *http.Request) {
	// check for correct perms
	if(!userpkg.CurrUser.CheckoutBook) {
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Print(err)
		return
	}
	tm["BooksInCart"] = booksInCart
	tm["BooksInCartCount"] = len(booksInCart)

	t, err := template.ParseFiles("static/cart.html", "static/header.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err.Error())
	}
	t.Execute(w, tm)
}

func DeleteFromCart(w http.ResponseWriter, r *http.Request) {
	// check for correct perms
	if(!userpkg.CurrUser.CheckoutBook) {
		errString := "You do not have permissions for this"
		log.Print(errString)
		http.Error(w, errString, http.StatusUnauthorized)
		return
	}

	// Parse the JSON request body
	Cart := struct {
		CartId int `json:"cartId"`
		BookId int `json:"bookId"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&Cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Then we querry to check if the book is in stock or exists
	books, err := dbQuerries.Book{}.Read("WHERE bookId=?", Cart.BookId)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(books) == 0 {
		errString := "book does not exist"
		log.Print("book does not exist")
		http.Error(w, errString, http.StatusBadRequest)
		return
	} 
	requestedBook := books[0]

	// then we update the stock to make sure it can't be further reduced
	updateStmt := "UPDATE books SET count=? WHERE bookId=?"
	_, err = dbutil.DB.Exec(updateStmt, (requestedBook.Count+1), Cart.BookId)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Then delete the book from the cart based on the cartId
	deleteStmt := "DELETE FROM cart WHERE userId=? AND cartId=?"
	_, err = dbutil.DB.Exec(deleteStmt, userpkg.CurrUser.UserId, Cart.CartId)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Checkout(w http.ResponseWriter, r *http.Request) {
	booksInCart, err := dbQuerries.BooksInCart{}.Read("WHERE userId=?", userpkg.CurrUser.UserId)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Then delete the book from the cart based on the cartId
	deleteStmt := "DELETE FROM cart WHERE userId=?"
	_, err = dbutil.DB.Exec(deleteStmt, userpkg.CurrUser.UserId)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	insertStmt := "INSERT INTO checkedOutBooks (userId, bookId, checkedOutDate, dueDate) VALUES (?, ?, ?, ?)"
	
	for _, book := range booksInCart {
		currentTime := time.Now()
		futureTime := currentTime.Add(30 * 24 * time.Hour)
		currentTimeAsString := strconv.Itoa(int(currentTime.Unix()))
		futureTimeAsString := strconv.Itoa(int(futureTime.Unix()))
		// first insert the books into the checked out books table
		_, err = dbutil.DB.Exec(insertStmt, userpkg.CurrUser.UserId, book.BookId, currentTimeAsString, futureTimeAsString)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}