package cardsHandler

import (
	"drello-api/pkg/presentation/rest/handler"
	"drello-api/pkg/presentation/rest/resttest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	resttest.InitTest()
	resttest.Router().Handle("/cards", handler.Handler(CardsHandler))
	os.Exit(m.Run())
}
