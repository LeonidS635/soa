package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func readBodyFromRequest(req *http.Request) ([]byte, error) {
	body := make([]byte, req.ContentLength)
	read, err := req.Body.Read(body)
	defer req.Body.Close()

	if err != io.EOF {
		return nil, fmt.Errorf("%w: %w", ReadBodyError, err)
	}
	if read != int(req.ContentLength) {
		return nil, fmt.Errorf("%w: read less than body length", ReadBodyError)
	}
	return body, nil
}
