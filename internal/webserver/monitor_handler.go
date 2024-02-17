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

func (wmh *WebMonitorHandler) GetResource(path string) {
	startTime := time.Now()

	res, err := http.Get(path)
	if err != nil {
		log.Printf("GOT: %s — LOG [Error]: %v\n", path, err)
		return
	}

	elapsedTime := time.Since(startTime).Milliseconds()

	expectedStatusCode := http.StatusOK

	if res.StatusCode == expectedStatusCode {
		log.Printf(
			"[GOT]: %s — [LOG]: Passed: Response as expected — [StatusCode]: %d — [TIME]: %dms\n",
			path,
			res.StatusCode,
			elapsedTime,
		)
	} else {
		log.Printf(
			"[GOT]: %s — [LOG]: Warning: Unexpected response - [StatusCode]: %d, Expected: %d — [TIME]: %dms\n",
			path,
			expectedStatusCode,
			res.StatusCode,
			elapsedTime,
		)
	}
}
