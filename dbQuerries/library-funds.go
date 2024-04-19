package dbQuerries

import(
	_ "database/sql"
	"cmpe-132-project/dbutil"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type LibraryFunds struct {
	NameOfFund string
	Funds int
}

func (lf LibraryFunds) Read(whereStmt string, args... interface{}) (funds []LibraryFunds, err error) {
		// Execute a query that returns rows
		rows, err := dbutil.DB.Query("SELECT * FROM libraryFunds " + whereStmt, args...)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
	
		// Iterate over the result set
		for rows.Next() {
			fund := LibraryFunds{}
			err = rows.Scan(
				&fund.NameOfFund,
				&fund.Funds)
			if err != nil {
				log.Fatal(err)
			}
			funds = append(funds, fund)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		return
}
