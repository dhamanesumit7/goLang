package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var db *sql.DB

func connectDB() {

	var err error

	db, err = sql.Open(
		"mysql",
		"root:S@um9594@tcp(127.0.0.1:3306)/golangdb",
	)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("MySQL Connected")
}

func createUser(w http.ResponseWriter, r *http.Request) {

	var user User

	json.NewDecoder(r.Body).Decode(&user)

	query := "INSERT INTO users(name,email,password) VALUES(?,?,?)"

	result, err := db.Exec(
		query,
		user.Name,
		user.Email,
		user.Password,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	id, _ := result.LastInsertId()

	user.ID = int(id)

	json.NewEncoder(w).Encode(user)
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query(
		"SELECT id,name,email,password FROM users",
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer rows.Close()

	var users []User

	for rows.Next() {

		var user User

		rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
		)

		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	query := "DELETE FROM users WHERE id = ?"

	result, err := db.Exec(query, id)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		http.Error(w, "User not found", 404)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "User deleted successfully",
	})

}

func atoi(s string) int {
	var num int
	fmt.Sscanf(s, "%d", &num)
	return num
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	query := `
		UPDATE users
		SET name = ? , email = ?, password = ?
		WHERE id = ?
	`

	result, err := db.Exec(

		query,
		user.Name,
		user.Email,
		user.Password,
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		http.Error(w, "user not found", 404)
		return
	}

	user.ID = atoi(id)

	json.NewEncoder(w).Encode(user)
}

func main() {

	connectDB()

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			getUsers(w, r)
		}

		if r.Method == "POST" {
			createUser(w, r)
		}

		if r.Method == "DELETE" {
			deleteUser(w, r)
		}

		if r.Method == "PUT" {
			updateUser(w, r)
		}
	})

	fmt.Println("Server running on 8080")

	http.ListenAndServe(":8080", nil)
}
