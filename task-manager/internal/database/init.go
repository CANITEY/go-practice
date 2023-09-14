package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type task struct {
	Id int64
	Title string
	Description string
	Status bool
}

var db, err = sql.Open("sqlite3", "tasks.db")

func init() {
	dbCreate()
}



// database creation process
func dbCreate() {
	if err != nil {
		log.Fatal(err)
	}

	//create database
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS tasks(title varchar(255), description varchar(500), status bool)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err)
	}else {
		fmt.Println("Database created successfully")
	}
}

func getLastID() (int64, error) {
	statement, err := db.Query("select MAX(ROWID) from tasks")
	if err != nil {
		return 0, err
	}
	var id int64
	statement.Next()
	statement.Scan(&id)
	return id, nil
}


//Adds a task into the tasks list
func Add(title, description string) (error) {
	statement, err := db.Prepare("INSERT INTO tasks(title, description, status) VALUES(?, ?, 0)")
	if err != nil {
		return err
	}

	_, err = statement.Exec(title, description)
	if err != nil {
		 return err
	}

	return nil
}


//Reads the tasks
func Read() ([]task, error){
	statement, err := db.Query("SELECT ROWID, * FROM tasks")
	if err != nil {
		return nil, err
	}
	var tasks []task
	for statement.Next() {
		var tsk task
		statement.Scan(&tsk.Id, &tsk.Title, &tsk.Description, &tsk.Status)
		tasks = append(tasks, tsk)
	}

	return tasks,nil
}

//Update status
func ToggleStatus(id int64) (error) {
	row := db.QueryRow("SELECT status FROM tasks WHERE ROWID=?", id)

	var status bool

	if err := row.Scan(&status); err != nil {
		return err
	}

	statement, err := db.Prepare("UPDATE tasks SET status=? WHERE ROWID=?")
	if err != nil {
		return err
	}

	if _, err := statement.Exec(!status, id); err != nil {
		return err
	}
	return nil
}

//Delete task
func Delete(id int64) (error) {
	statement, err := db.Prepare("DELETE FROM tasks WHERE ROWID=?")
	if err != nil {
		return nil
	}

	if _, err := statement.Exec(id); err != nil {
		return err
	}

	return nil
}
