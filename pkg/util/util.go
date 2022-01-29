package util

import (
	"drello-api/pkg/util/log"
	"os"
)

func MustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Warn("%s environment variable not set.", k).Write()
	}
	return v
}
