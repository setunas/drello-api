package cardHandler

import (
	"drello-api/pkg/presentation/rest/handler"
	"drello-api/pkg/presentation/rest/resttest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	resttest.InitTest()
	resttest.Router().Handle("/cards/{id:[0-9]+}", handler.Handler(CardHandler))
	exitCode := m.Run()
	resttest.CleanupAllTable()
	os.Exit(exitCode)
}
