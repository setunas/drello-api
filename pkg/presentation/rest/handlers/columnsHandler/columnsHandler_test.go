package columnsHandler

import (
	"drello-api/pkg/presentation/rest/handler"
	"drello-api/pkg/presentation/rest/resttest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	resttest.InitTest()
	resttest.Router().Handle("/columns", handler.Handler(ColumnsHandler))
	os.Exit(m.Run())
}
