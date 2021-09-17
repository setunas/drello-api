package main

import (
	"drello-api/pkg/infrastructure/mysql"
	"drello-api/pkg/presentation/rest"
	"log"
)

func main() {
	err := mysql.Open()
	if err != nil {
		log.Println(err)
	}

	rest.HandleRequests()
}
