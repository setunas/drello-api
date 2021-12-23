package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func logHTTPRequest(r *http.Request) {
	var body interface{}
	json.NewDecoder(r.Body).Decode(&body)

	requestLog := map[string]string{
		"Method":    r.Method,
		"URI":       r.URL.String(),
		"Referer":   r.Header.Get("Referer"),
		"UserAgent": r.Header.Get("User-Agent"),
		"Body":      fmt.Sprintln(body),
	}

	message := "HTTP Request:"
	for k, v := range requestLog {
		message += fmt.Sprintf("\n%v: %v", k, v)
	}

	log.Print(message)
}
