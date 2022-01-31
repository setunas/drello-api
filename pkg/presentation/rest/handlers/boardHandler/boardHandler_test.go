package boardHandler

import (
	"drello-api/pkg/presentation/rest/handler"
	"drello-api/pkg/presentation/rest/resttest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	resttest.InitTest()
	resttest.Router().Handle("/boards/{id:[0-9]+}", handler.Handler(BoardHandler))
	os.Exit(m.Run())
}
