package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	timeout = 10
)

func SendGetRequest(url string, result interface{}) error {
	// req.Header.Add(SecretKeyName, s.secretKey)

	client := &http.Client{
		Timeout: timeout * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	// Call Remote API
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// Read response Body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// If result if nil, do not unmarshall result
	if result == nil {
		return nil
	}

	// Read body into desired type
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	return nil
}

func GetBytes(key interface{}) ([]byte, error) {
	buf, ok := key.([]byte)
	if !ok {
		return nil, fmt.Errorf("ooops, did not work")
	}

	return buf, nil
}
