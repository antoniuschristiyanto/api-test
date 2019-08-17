package main

import (
	"api-test/services"
	"database/sql"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()
	db := initDB("users.db")
	migrate(db)

	e.GET("/users", services.GetUser(db))
	e.POST("/users", services.InsertUser(db))
	e.PATCH("/users/update-handphone/:id", services.UpdateHandphoneUser(db))
	e.PATCH("/users/update-status/:id", services.UpdateStatusUser(db))
	e.DELETE("/users/delete/:id", services.DeleteUser(db))

	e.Logger.Fatal(e.Start(":9000"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS users(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		email VARCHAR NOT NULL,
		handphone VARCHAR NOT NULL,
		status INTEGER
    );
    `

	_, err := db.Exec(sql)

	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
