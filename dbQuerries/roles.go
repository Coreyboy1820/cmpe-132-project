package dbQuerries

import(
	_ "database/sql"
	"cmpe-132-project/dbutil"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Roles struct {
	RoleId int
	PermsId int
	RoleName string
}

func (r Roles) Read(whereStmt string, args... interface{}) (roles []Roles, err error) {
		// Execute a query that returns rows
		rows, err := dbutil.DB.Query("SELECT * FROM roles " + whereStmt, args...)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
	
		// Iterate over the result set
		for rows.Next() {
			role := Roles{}
			err = rows.Scan(
				&role.RoleId,
				&role.PermsId,
				&role.RoleName)
			if err != nil {
				log.Fatal(err)
			}
			roles = append(roles, role)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		return
}
