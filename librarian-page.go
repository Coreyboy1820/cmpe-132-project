package main

import(
	"html/template"
	"net/http"
	"log"
	_ "encoding/json"
	"strconv"
	userpkg "cmpe-132-project/user"
	"cmpe-132-project/dbQuerries"
	"cmpe-132-project/dbutil"

	_ "github.com/mattn/go-sqlite3"

)

/// HandleLibrarianPage makes sure the user has the credentials to be on the page then renders the librarian page

func HandleLibrarianPage(w http.ResponseWriter, r *http.Request){
	if(!userpkg.CurrUser.LoggedIn || !userpkg.CurrUser.AddBooks) {
		errString := "You do not have access"
		log.Print(errString)
		http.Error(w, errString, http.StatusUnauthorized)
		return
	}
	var err error
	tm := make(map[string]interface{})
	tm["CurrUser"] = userpkg.CurrUser

	booksInCart, err := dbQuerries.BooksInCart{}.Read("WHERE userId=?", userpkg.CurrUser.UserId)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tm["BooksInCartCount"] = len(booksInCart)

	buyRequests, err := dbQuerries.BuyRequests{}.Read("ORDER BY approved")
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tm["BuyRequests"] = buyRequests

	t, err := template.ParseFiles("static/librarian-page.html", "static/header.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err.Error())
	}
	t.Execute(w, tm)
}

/// AddNewBook simulates an email exchange, again in the console, of a librarian requesting to add a book to the library,
/// 		   it will also check if the library has the funds to buy the book and will then add the book to the home page

func AddNewBook(w http.ResponseWriter, r *http.Request){
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		errString := "Failed to parse form data"
		log.Print(errString)
		http.Error(w, errString, http.StatusBadRequest)
		return
	}
	// Retrieve form values
	link := r.Form.Get("link")
	bookName := r.Form.Get("bookName")
	isbn := r.Form.Get("isbn")
	cost := r.Form.Get("cost")
	quantity := r.Form.Get("amount")

	// auto approve the request due to there being no other interaction so set the final argument to true
	err = InsertNewBuyRequest(link, bookName, isbn, cost, quantity, true)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf(`
	This is a simulated email:
	A librarian wishes to buy a book with the following information 
	link %s
	Book Name: %s
	ISBN: %s
	Cost: %s
	Amount: %s`, 
	link, bookName, isbn, cost, quantity)

	log.Print("\n\nApproved\n")

	fundsAvailable, err := UpdateFunds(cost, quantity, "Book Fund")
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if(!fundsAvailable) {
		http.Redirect(w, r, "/librarianPage/", http.StatusSeeOther)
		return
	}

	err = InsertNewBook(bookName, quantity, isbn)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/librarianPage/", http.StatusSeeOther)
}

/// UpdateFunds alows for the decrease or increase of money in the different funds the library may have

/*
	params:
		cost - the cost per book
		quantity - the number of books to buy
		fundName - the name of the fund to take the total amount out of
	returns:
		fundsAvailable - a boolean indicating whether or not the fund requested has enough money for the purchase
		err - an error if one is thrown
	
*/

func UpdateFunds(cost string, quantity string, fundName string) (fundsAvailable bool, err error){

	funds, err := dbQuerries.LibraryFunds{}.Read("WHERE nameOfFund=?", fundName)
	if err != nil {
		log.Print(err)
		return
	}
	if len(funds) == 0 {
		errString := "Fund does not exist"
		log.Print(errString)
		return
	}
	fundToUpdate := funds[0]

	costInt,err := strconv.Atoi(cost)
	if err != nil {
		log.Print(err)
		return
	}
	quantityInt, err := strconv.Atoi(quantity)
	if err != nil {
		log.Print(err)
		return
	}

	totalCost := (costInt*quantityInt)
	newFundTotal := fundToUpdate.Funds-totalCost
	updateStmt := "UPDATE libraryFunds SET funds=? WHERE nameOfFund=?"
	if newFundTotal < 0  {
		log.Print("Not enough funds")
		fundsAvailable = false
		return
	} else {
		fundsAvailable = true
	}
	_, err = dbutil.DB.Exec(updateStmt, newFundTotal, fundName)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

/// InsertNewBook will insert a new book into the library to be checked out by other students

/*
	params:
		bookName - the name of the book to insert into the database
		bookCount - the quantity of books to be added to the library
		isbn - the isbn of the book
	returns:
		err - an error if one is thrown
*/

func InsertNewBook(bookName string, bookCount string, isbn string) (err error){
	insertStatement := "INSERT INTO books (bookName, count, isbn) VALUES (?, ?, ?)"
	_, err = dbutil.DB.Exec(insertStatement, bookName, bookCount, isbn)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

/// InsertNewBuyRequest will insert a buy request into the database when the librarian requests

/*
	params:
		link - the link to where to buy the book at
		bookName - the name of the book
		isbn - the isbn of the book 
		cost - the cost of each book 
		amount - quantity of books requested
		status - the status of the request
*/

func InsertNewBuyRequest(link string, bookName string, isbn string, cost string, amount string, status bool) (err error) {
	// finally, we insert the book into the cart
	insertStmt := "INSERT INTO buyRequests (bookLink, bookName, isbn, bookCost, bookCount, approved) VALUES (?, ?, ?, ?, ?, ?)"

	_, err = dbutil.DB.Exec(insertStmt, link, bookName, isbn, cost, amount, status)
	if err != nil {
		log.Print(err)
		return
	}
	return
}