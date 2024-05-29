package main

import (
	"fmt"
	database "student-plan/db"
)

func main() {
	fmt.Println("Merhaba Dünya!")
	database.Connect()
}
