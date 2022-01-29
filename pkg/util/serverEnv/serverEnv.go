package serverEnv

import (
	"log"
	"os"
)

type ServerEnvType string

const (
	Dev  ServerEnvType = "dev"
	Stg  ServerEnvType = "stg"
	Prod ServerEnvType = "prod"
)

func ServerEnv() ServerEnvType {
	switch env := os.Getenv("SERVER_ENV"); env {
	case "prod":
		return Prod
	case "stg":
		return Stg
	case "dev":
		return Dev
	default:
		log.Fatalf("SERVER_ENV not set properly. %s", env)
		return Dev
	}
}
