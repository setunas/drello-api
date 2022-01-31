package columnHandler

import (
	"drello-api/pkg/presentation/rest/handler"
	"drello-api/pkg/presentation/rest/resttest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	resttest.InitTest()
	resttest.Router().Handle("/columns/{id:[0-9]+}", handler.Handler(ColumnHandler))
	os.Exit(m.Run())
}
