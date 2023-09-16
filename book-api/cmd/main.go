package main

import (
	"book-api/internal/database"
	"fmt"
)

func main() {
	fmt.Println(database.ViewBook(2))
}
