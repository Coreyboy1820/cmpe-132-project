package main

import (
	"html/template"
	"net/http"
	"log"
	userpkg "cmpe-132-project/user"
	"cmpe-132-project/dbQuerries"

	_ "github.com/mattn/go-sqlite3"

)

func HandleCart(w http.ResponseWriter, r *http.Request) {
	tm := make(map[string]interface{})
	tm["CurrUser"] = userpkg.CurrUser

	// --------- Read all books ---------

	booksInCart, err := dbQuerries.BooksInCart{}.Read("WHERE userId=?", userpkg.CurrUser.UserId)
	if err != nil {
		log.Print(err)
	}
	tm["BooksInCart"] = booksInCart
	tm["BooksInCartCount"] = len(booksInCart)

	t, err := template.ParseFiles("static/cart.html", "static/header.html")
	if err != nil {
		panic(err.Error())
	}
	t.Execute(w, tm)
}