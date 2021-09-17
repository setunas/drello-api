package utils

import (
	"log"
	"os"
)

func MustGetenv(k string) string {
	v := os.Getenv(k)
	log.Println("ENV", k, v)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}
