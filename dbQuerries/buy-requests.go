package dbQuerries

import(
	_ "database/sql"
	"cmpe-132-project/dbutil"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type BuyRequests struct {
	BookLink string
	BookName string
	Isbn string
	BookCost int
	BookCount int
	Approved bool
}

func (br BuyRequests) Read(whereStmt string, args... interface{}) (requests []BuyRequests, err error) {
		// Execute a query that returns rows
		rows, err := dbutil.DB.Query("SELECT * FROM buyRequests " + whereStmt, args...)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
	
		// Iterate over the result set
		for rows.Next() {
			request := BuyRequests{}
			err = rows.Scan(
				&request.BookLink,
				&request.BookName,
				&request.Isbn,
				&request.BookCost,
				&request.BookCount,
				&request.Approved)
			if err != nil {
				log.Fatal(err)
			}
			requests = append(requests, request)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		return
}
