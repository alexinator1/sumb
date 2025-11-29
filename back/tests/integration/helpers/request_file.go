package helpers

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func ParseHttpRequestFromFile(httpRequestFile string) (*http.Request, error) {

	httpReqContent, err := os.ReadFile(httpRequestFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read http request from file %s: %w", httpRequestFile, err)
	}

	reader := bufio.NewReader(bytes.NewBuffer(httpReqContent))
	req, err := http.ReadRequest(reader)
	if err != nil {
		return nil, err
	}
	return req, nil
}
