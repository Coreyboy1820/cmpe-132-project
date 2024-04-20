package dbQuerries

import(
	_ "database/sql"
	"cmpe-132-project/dbutil"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	BookId int
	BookName string
	Count int
	Isbn string
}

func (b Book) Read(whereStmt string, args... interface{}) (books []Book, err error) {
		// Execute a query that returns rows
		rows, err := dbutil.DB.Query("SELECT * FROM books " + whereStmt, args...)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
	
		// Iterate over the result set
		for rows.Next() {
			book := Book{}
			err = rows.Scan(
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