package boardHandler

import (
	"drello-api/pkg/presentation/rest/handler"
	"drello-api/pkg/presentation/rest/test"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	test.InitTest()
	test.Router().Handle("/boards/{id:[0-9]+}", handler.Handler(BoardHandler))
	os.Exit(m.Run())
}
