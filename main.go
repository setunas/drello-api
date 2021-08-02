package main

import (
	"drello-api/pkg/infrastracture/mysql"
	"drello-api/pkg/presentation/http/rest"
	"fmt"
)

func main() {
	_, err := mysql.Open()
	if err != nil {
		fmt.Println(err)
	}

	rest.HandleRequests()
}
