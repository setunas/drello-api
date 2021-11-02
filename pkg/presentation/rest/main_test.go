package rest

import (
	"drello-api/pkg/infrastructure/mysql"
	"log"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
	err := mysql.TestOpen()
	if err != nil {
		log.Println(err)
	}

	router = mux.NewRouter()
	setHandlers()
	os.Exit(m.Run())
}
