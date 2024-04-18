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


func HomeHandler(w http.ResponseWriter, r *http.Request){
	var err error
	tm := make(map[string]interface{})
	tm["CurrUser"] = userpkg.CurrUser

	// --------- Read all books ---------

	tm["Books"], err =dbQuerries.Book{}.Read("")
	if err != nil {
		log.Print(err)
	}

	booksInCart, err := dbQuerries.BooksInCart{}.Read("WHERE userId=?", userpkg.CurrUser.UserId)
	if err != nil {
		log.Print(err)
	}
	tm["BooksInCartCount"] = len(booksInCart)

	t, err := template.ParseFiles("static/home.html", "static/header.html")
	if err != nil {
		panic(err.Error())
	}
	t.Execute(w, tm)
}

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
		insertStmt := "INSERT INTO cart (userId, bookId) VALUES (?, ?)"
		_, err = dbutil.DB.Exec(insertStmt, userpkg.CurrUser.UserId, book.BookId)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
}
