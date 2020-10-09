package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	debug         = flag.Bool("debug", false, "enable debugging")
	password      = flag.String("password", "", "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", "", "the database server")
	database      = flag.String("database", "", "the database name")
	user          = flag.String("user", "", "the database user")
)

func main() {
	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", *server, *user, *password, *port, *database)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Conneted!\n")

	count, err := ReadBanks(db)
	fmt.Printf("Bank Count: %d\n", count)

	count, err = ReadAccounts(db)
	if count != -1 {
		fmt.Printf("Account Count: %d\n", count)
	}

}

// ReadBanks returns count, error
func ReadBanks(db *sql.DB) (int, error) {
	ctx := context.Background()

	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("SELECT * FROM dbo.Banks;")

	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err
	}
	defer rows.Close()

	var count int

	for rows.Next() {
		var ID, Code int
		var Name string

		err := rows.Scan(&ID, &Code, &Name)
		if err != nil {
			return -1, err
		}

		fmt.Printf("ID: %d, Code: %d, Name: %s\n", ID, Code, Name)
		count++
	}

	return count, nil
}

// ReadAccounts returns count, error
func ReadAccounts(db *sql.DB) (int, error) {
	ctx := context.Background()

	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("SELECT * FROM dbo.Accounts;")

	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err
	}
	defer rows.Close()

	var count int

	for rows.Next() {
		var ID, UserID, BankID int
		var Code string

		err := rows.Scan(&ID, &UserID, &BankID, &Code)
		if err != nil {
			return -1, err
		}

		fmt.Printf("ID: %d, UserID: %d, BankID:%d, Code: %s\n", ID, UserID, BankID, Code)
		count++
	}

	return count, nil
}
