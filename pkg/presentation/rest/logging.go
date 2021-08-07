package rest

import (
	"fmt"
	"log"
	"net/http"
)

func logHTTPRequest(r *http.Request) {
	requestLog := map[string]string{
		"method":    r.Method,
		"uri":       r.URL.String(),
		"referer":   r.Header.Get("Referer"),
		"userAgent": r.Header.Get("User-Agent"),
	}

	message := "#### HTTP Request ####"
	for k, v := range requestLog {
		message += fmt.Sprintf("\n%v: %v", k, v)
	}

	log.Print(message)
}
