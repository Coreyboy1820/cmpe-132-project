package main

import(
	"html/template"
	"net/http"
	"log"
	userpkg "cmpe-132-project/user"
	"cmpe-132-project/dbQuerries"

	_ "github.com/mattn/go-sqlite3"

)


func HomeHandler (w http.ResponseWriter, r *http.Request){
	var err error
	tm := make(map[string]interface{})
	tm["CurrUser"] = userpkg.CurrUser

	// --------- Read all books ---------

	tm["Books"], err =dbQuerries.Book{}.Read("")
	if err != nil {
		log.Print(err)
	}

	t, err := template.ParseFiles("static/home.html", "static/header.html")
	if err != nil {
		panic(err.Error())
	}
	t.Execute(w, tm)
}