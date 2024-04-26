package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Task struct {
	ID        int
	Name      string
	Completed bool
}

const (
	host     = "localhost"
	user     = "postgres"
	port     = 5432
	password = "19082004"
	dbname   = "data"
)

func main() {
	pq := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", pq)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var input int

	for {
		fmt.Print("Enter option (1: Create, 2: Delete, 3: Update, 4: Read, 0: Exit): ")
		fmt.Scan(&input)

		switch input {
		case 1:
			var a int
			fmt.Print("Enter id: ")
			fmt.Scan(&a)
			var s string
			fmt.Print("Enter task: ")
			fmt.Scan(&s)
			err = createTask(db, a, s)
		case 2:
			var a int
			fmt.Print("id to delete: ")
			fmt.Scan(&a)
			err = deleteTask(db, a)
		case 3:
			var s int
			fmt.Print("Enter id to update(set TRUE): ")
			fmt.Scan(&s)
			err = updateTask(db, s, true)
		case 4:
			rows, err := readTasks(db)
			if err != nil {
				fmt.Printf("Failed to read tasks: %v", err)
				continue
			}
			defer rows.Close()

			for rows.Next() {
				var id int
				var name string
				var completed bool
				if err := rows.Scan(&id, &name, &completed); err != nil {
					fmt.Printf("Failed to scan task: %v", err)
					continue
				}
				fmt.Printf("ID: %d, Name: %s, Completed: %t\n", id, name, completed)
			}
		case 0:
			return
		default:
			fmt.Println("Invalid option")
		}

		if err != nil {
			fmt.Printf("Operation failed: %v", err)
		}
	}
}

func createTask(db *sql.DB, id int, name string) error { //works
	_, err := db.Query("Insert INTO tasks (id, name) VALUES ($1, $2)", id, name)
	return err
}
func deleteTask(db *sql.DB, id int) error { //works
	_, err := db.Query("DELETE FROM tasks WHERE id = $1", id)
	return err
}
func updateTask(db *sql.DB, id int, cmp bool) error {

	_, err := db.Exec("UPDATE tasks SET completed = $1 WHERE id = $2", cmp, id)

	return err
}
func readTasks(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}

	return rows, nil
}
