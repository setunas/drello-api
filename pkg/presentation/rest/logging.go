package rest

import (
	"fmt"
	"log"
	"net/http"
)

func logHTTPRequest(r *http.Request) {
	requestLog := map[string]string{
		"Method":    r.Method,
		"URI":       r.URL.String(),
		"Referer":   r.Header.Get("Referer"),
		"UserAgent": r.Header.Get("User-Agent"),
	}

	message := "HTTP Request:"
	for k, v := range requestLog {
		message += fmt.Sprintf("\n%v: %v", k, v)
	}

	log.Print(message)
}
