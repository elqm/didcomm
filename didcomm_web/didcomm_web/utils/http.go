package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func HttpRequest(url string, content interface{}) ([]byte, error) {
	jsonContent, err := json.Marshal(content)
	if err != nil {
		log.Printf("json err: %v\n", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonContent))
	if err != nil {
		log.Printf("request err: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("resp err: %v\n", err)
	}
	defer resp.Body.Close()

	// ioutil.ReadAll(deprecated) -> io.ReadAll 대체
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("respBody err: %v\n", err)
	}

	return respBody, nil
}
