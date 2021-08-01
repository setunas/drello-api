package main

import (
	"drello-api/pkg/presentation/http/rest"
	"fmt"
)

func main() {
	fmt.Println("booting drello-api")
	fmt.Println("Listening on tcp://127.0.0.1:8080")
	rest.HandleRequests()
}
