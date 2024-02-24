package webserver

import (
	"log"
	"net/http"
	"time"
)

type WebMonitorHandler struct{}

func NewWebMonitorHandler() *WebMonitorHandler {
	return &WebMonitorHandler{}
}

func (wmh *WebMonitorHandler) GetResource(url string, expectedStatusCode int) {
	startTime := time.Now()

	res, err := http.Get(url)
	if err != nil {
		log.Printf("GOT: %s — LOG [Error]: %v\n", url, err)
		return
	}

	elapsedTime := time.Since(startTime).Milliseconds()

	if res.StatusCode == expectedStatusCode {
		log.Printf(
			"[GOT]: %s — [LOG]: Passed: Response as expected — [StatusCode]: %d — [TIME]: %dms\n",
			url,
			res.StatusCode,
			elapsedTime,
		)
	} else {
		log.Printf(
			"[GOT]: %s — [LOG]: Warning: Unexpected response - [StatusCode]: %d, Expected: %d — [TIME]: %dms\n",
			url,
			expectedStatusCode,
			res.StatusCode,
			elapsedTime,
		)
	}
}
