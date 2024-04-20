package main

import(
	"html/template"
	"net/http"
	"log"
	"encoding/json"
	userpkg "cmpe-132-project/user"
	"cmpe-132-project/dbQuerries"
	"cmpe-132-project/dbutil"

	_ "github.com/mattn/go-sqlite3"

)

/// HomeHandler renders the home page

func HomeHandler(w http.ResponseWriter, r *http.Request){
	var err error
	tm := make(map[string]interface{})
	tm["CurrUser"] = userpkg.CurrUser

	// --------- Read all books ---------

	tm["Books"], err =dbQuerries.Book{}.Read("")
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	booksInCart, err := dbQuerries.BooksInCart{}.Read("WHERE userId=?", userpkg.CurrUser.UserId)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tm["BooksInCartCount"] = len(booksInCart)

	t, err := template.ParseFiles("static/home.html", "static/header.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err.Error())
	}
	t.Execute(w, tm)
}

/// AddToCart will update the users cart, and decrease the stock of the book if it is still in stock, if it is not in stock then it wont allow the book to be checked out

func AddToCart(w http.ResponseWriter, r *http.Request){
	// Parse the JSON request body
	book := struct {
		BookId string `json:"bookId"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Then we querry to check if the book is in stock or exists
	books, err := dbQuerries.Book{}.Read("WHERE bookId=?", book.BookId)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(books) == 0 {
		errString := "book does not exist"
		http.Error(w, errString, http.StatusBadRequest)
		return
	} 
	requestedBook := books[0]
	if books[0].Count == 0 {
		errString := "Book out of stock"
		log.Print(errString)
		http.Error(w, errString, http.StatusBadRequest)
		return
	}

	// then we update the stock to make sure it can't be further reduced
	updateStmt := "UPDATE books SET count=? WHERE bookId=?"
	_, err = dbutil.DB.Exec(updateStmt, (requestedBook.Count-1), book.BookId)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// finally, we insert the book into the cart
	insertStmt := "INSERT INTO cart (userId, bookId) VALUES (?, ?)"
	_, err = dbutil.DB.Exec(insertStmt, userpkg.CurrUser.UserId, book.BookId)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
