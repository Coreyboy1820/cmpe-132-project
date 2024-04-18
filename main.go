package main

import (
	"database/sql"
	_ "html/template"
	"log"
	"net/http"
	"cmpe-132-project/dbutil"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	// Open a connection to the SQLite database
	var err error
	dbutil.DB, err = sql.Open("sqlite3", "rbac.db")
	if err != nil {
		log.Fatal(err)
	}

	cssf := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css", cssf))
	
	http.HandleFunc("/logIn/", LoginPageHandler)
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/signin/", Login)
	http.HandleFunc("/logout/", Logout)
	http.HandleFunc("/firstTimeLogin/", FirstTimeLoginHandler)
	http.HandleFunc("/submitTemporaryPassword/", SubmitTempPasword)
	http.HandleFunc("/newPassword/", HandleNewPassword)
	http.HandleFunc("/submitNewPassword/", SubmitNewPassword)
	http.HandleFunc("/checkoutBook/", CheckoutBook)

	log.Println("Starting server on localhost:8080")
	http.ListenAndServe(":8080", nil)
}