package server

import (
	"github.com/canitey/simpleTaskManager/internal/database"
	"fmt"
	"log"
	"net/http"
	"strconv"
)



func Index(w http.ResponseWriter, r *http.Request) {
	content := `
	<html>
	<head>
	<style>
	* {
		font-family: arial;
	}
	table, td {
		border: 1px solid black;
	}
	td {
		padding: 5px 10px; !important
	}
	</style>
	</head>
	<body>
	<h1>Task manager</h1>
	<table>
	<thead>
	<tr>
	<td>title</td>
	<td>description</td>
	<td>done ?</td>
	<td>delete</td>
	</tr>
	`
	
	taskTemplate := `
	<tr>
	<td>%v</td>
	<td>%v</td>
	<td><a href="/toggle/?id=%v">%v</a></td>
	<td><a href="/delete/?id=%v">delete</a></td>
	</tr>
	`


	form := `
	</table>
	<form method=post action="add/">
	<input type=text placeholder="title" name="title">
	<br>
	<textarea name="description" placeholder="description"></textarea>
	<br>
	<input type=submit>
	</form>
	</body>
	</html>
	`
	tasks, _ := database.Read()
	fmt.Fprint(w, content)
	for _, task := range tasks {
		var status string
		if task.Status {
			status = "done"
		}else {
			status = "not yet"
		}
		fmt.Fprintf(w, taskTemplate, task.Title, task.Description, task.Id, status, task.Id)
	}
	fmt.Fprint(w, form)
}


func Toggle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("id")
	id, err := strconv.Atoi(query)
	id64 := int64(id)
	if err != nil {
		log.Printf("%v", err)
	}
	database.ToggleStatus(id64)
	fmt.Fprint(w, "<script>window.location='/'</script>")
}

func Add(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	queries := r.PostForm
	title := queries.Get("title")
	description := queries.Get("description")
	database.Add(title, description)
	fmt.Fprint(w, "<script>window.location='/'</script>")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Print(err)
	}
	database.Delete(int64(intId))
	fmt.Fprint(w, "<script>window.location='/'</script>")
}
