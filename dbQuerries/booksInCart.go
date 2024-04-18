package dbQuerries

import(
	_ "database/sql"
	"cmpe-132-project/dbutil"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type BooksInCart struct {
	CartId int
	UserId int
	BookId int
	BookName string
	Count int
	Isbn string 
}

func (b BooksInCart) Read(whereStmt string, args... interface{}) (books []BooksInCart, err error) {
		// Execute a query that returns rows
		rows, err := dbutil.DB.Query("SELECT * FROM booksInCart " + whereStmt, args...)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
	
		// Iterate over the result set
		for rows.Next() {
			book := BooksInCart{}
			err = rows.Scan(
				&book.CartId,
				&book.UserId,
				&book.BookId,
				&book.BookName,
				&book.Count,
				&book.Isbn)
			if err != nil {
				log.Fatal(err)
			}
			books = append(books, book)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		return
}
