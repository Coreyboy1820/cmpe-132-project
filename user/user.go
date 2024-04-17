package user

import (
	_ "database/sql"
	"cmpe-132-project/dbutil"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	FirstName string
	LastName string
	StudentId string
	PasswordHash string
	PasswordSet bool
	Salt string
	Email string
	LoggedIn bool
	RoleName string
	CheckoutBook bool
	CheckinBook bool
	ReserveRoom bool
	ReserveBooks bool
	AddBooks bool
	CreateUser bool
	UpdateUser bool
	DeleteUser bool
};

func (u User) Read(whereStmt string, args... interface{}) (users []User, err error) {
	// Execute a query that returns rows
    rows, err := dbutil.DB.Query("SELECT * FROM usersAndPerms " + whereStmt, args...)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // Iterate over the result set
    for rows.Next() {
		user := User{}
		err = rows.Scan(
			&user.FirstName,
			&user.LastName,
			&user.StudentId,
			&user.PasswordHash,
			&user.Salt,
			&user.Email,
			&user.PasswordSet,
			&user.LoggedIn,
			&user.RoleName,
			&user.CheckoutBook,
			&user.CheckinBook,
			&user.ReserveRoom,
			&user.ReserveBooks,
			&user.AddBooks,
			&user.CreateUser,
			&user.UpdateUser,
			&user.DeleteUser)
        if err != nil {
            log.Fatal(err)
        }
		users = append(users, user)
    }
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
	return
}

var CurrUser User