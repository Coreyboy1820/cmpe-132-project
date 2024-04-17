package main

import(
	"html/template"
	"net/http"
	userpkg "cmpe-132-project/user"

	_ "github.com/mattn/go-sqlite3"

)


func HomeHandler (w http.ResponseWriter, r *http.Request){
	tm := make(map[string]interface{})
	tm["CurrUser"] = userpkg.CurrUser
	t, err := template.ParseFiles("home.html")
	if err != nil {
		panic(err.Error())
	}
	t.Execute(w, tm)
}