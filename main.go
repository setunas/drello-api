package main

import (
	"drello-api/pkg/infrastracture/mysql"
	"drello-api/pkg/presentation/http/rest"
	"fmt"
)

func main() {
	fmt.Println("booting drello-api")
	fmt.Println("Listening on http://127.0.0.1:8080")

	_, err := mysql.Open()
	if err != nil {
		fmt.Println(err)
	}
	rest.HandleRequests()
}
