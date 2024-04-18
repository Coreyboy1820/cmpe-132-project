package dbQuerries

import(
	_ "database/sql"
	"cmpe-132-project/dbutil"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type BooksAndCheckedOut struct {
	CheckedOutBooksId int
	BookId int
	UserId int
	BookName string
	Count int
	Isbn string
	CheckedOutDate string
	DueDate string
}

func (b BooksAndCheckedOut) Read(whereStmt string, args... interface{}) (books []BooksAndCheckedOut, err error) {
		// Execute a query that returns rows
		rows, err := dbutil.DB.Query("SELECT * FROM booksAndCheckedOut " + whereStmt, args...)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
	
		// Iterate over the result set
		for rows.Next() {
			book := BooksAndCheckedOut{}
			err = rows.Scan(
				&book.CheckedOutBooksId,
				&book.BookId,
				&book.UserId,
				&book.BookName,
				&book.Count,
				&book.Isbn,
				&book.CheckedOutDate,
				&book.DueDate)
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
