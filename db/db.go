package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InsertUser(id string, name string, surname string, othername string) (int64, error) {
	fmt.Println("Go MySQL database")
	db, err := sql.Open("mysql", "root:QWEfghUIO0!@tcp(127.0.0.1:3306)/mydb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users VALUES(?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}

	res, err := stmt.Exec(id, name, surname, othername)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.RowsAffected()

	if err != nil {
		panic(err.Error())
	}
	return lastId, nil
}

func DeleteUser(id string) (int64, error) {
	fmt.Println("Go MySQL database")
	db, err := sql.Open("mysql", "root:QWEfghUIO0!@tcp(127.0.0.1:3306)/mydb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")

	if err != nil {
		panic(err.Error())
	}

	res, err := stmt.Exec(id)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.RowsAffected()

	if err != nil {
		panic(err.Error())
	}
	return lastId, nil
}

/*
func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:QWEfghIOP0!@tcp(127.0.0.1:3306)/mydb")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO users VALUES ( 1, 'TEST' , 'Test1','Test2' )")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}
*/
