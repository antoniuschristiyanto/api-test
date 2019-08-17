package models

import (
	"database/sql"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Handphone string `json:"handphone"`
	Status    int    `json:"status"`
}

type UserCollection struct {
	Users []User `json:"result"`
}

func GetUser(db *sql.DB) UserCollection {
	sql := "SELECT * FROM users"
	rows, err := db.Query(sql)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		panic(err)
	}
	// make sure to cleanup when the program exits
	defer rows.Close()

	result := UserCollection{}
	for rows.Next() {
		user := User{}
		err2 := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Handphone, &user.Status)
		// Exit if we get an error
		if err2 != nil {
			panic(err2)
		}
		result.Users = append(result.Users, user)
	}
	return result
}

func InsertUser(db *sql.DB, name string, email string, handphone string, status int) (int64, error) {
	sql := "INSERT INTO users(name, email, handphone, status) VALUES(?,?,?,?)"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'name, email, handphone and status'
	result, err2 := stmt.Exec(name, email, handphone, status)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

func UpdateHandphoneUser(db *sql.DB, userId int, handphone string) (int64, error) {
	sql := "UPDATE users set handphone = ? WHERE id = ?"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(handphone, userId)

	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}

func UpdateStatusUser(db *sql.DB, userId int, status int) (int64, error) {
	sql := "UPDATE users set status = ? WHERE id = ?"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(status, userId)

	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}

func DeleteUser(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM users WHERE id = ?"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}

	// Replace the '?' in our prepared statement with 'id'
	result, err2 := stmt.Exec(id)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
