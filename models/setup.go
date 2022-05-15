package models

import (
	"fmt"
)

func ConnectDatabase(id string, pass string) {
	fmt.Println("Connected To Database !")
	fmt.Println(id, pass)
}
