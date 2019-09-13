package scalewayapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)


// response
func (s *ScalewayAPI) response(method, uri string, content io.Reader) (resp *http.Response, err error) {
	var (
		req *http.Request
	)

	req, err = http.NewRequest(method, uri, content)
	if err != nil {
		err = fmt.Errorf("response %s %s", method, uri)
		return
	}
	req.Header.Set("X-Auth-Token", s.Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", s.userAgent)
	resp, err = s.client.Do(req)
	return
}

// GetResponse returns an http.Response object for the updated resource
func (s *ScalewayAPI) GetResponse(apiURL, resource string, data interface{}) (*http.Response, error) {
	payload := new(bytes.Buffer)
	if err := json.NewEncoder(payload).Encode(data); err != nil {
		return nil, err
	}
	return s.response("GET", fmt.Sprintf("%s/%s", strings.TrimRight(apiURL, "/"), resource), payload)
}

// PostResponse returns an http.Response object for the updated resource
func (s *ScalewayAPI) PostResponse(apiURL, resource string, data interface{}) (*http.Response, error) {
	payload := new(bytes.Buffer)
	if err := json.NewEncoder(payload).Encode(data); err != nil {
		return nil, err
	}
	return s.response("POST", fmt.Sprintf("%s/%s", strings.TrimRight(apiURL, "/"), resource), payload)
}

// PatchResponse returns an http.Response object for the updated resource
func (s *ScalewayAPI) PatchResponse(apiURL, resource string, data interface{}) (*http.Response, error) {
	payload := new(bytes.Buffer)
	if err := json.NewEncoder(payload).Encode(data); err != nil {
		return nil, err
	}
	return s.response("PATCH", fmt.Sprintf("%s/%s", strings.TrimRight(apiURL, "/"), resource), payload)
}

// PutResponse returns an http.Response object for the updated resource
func (s *ScalewayAPI) PutResponse(apiURL, resource string, data interface{}) (*http.Response, error) {
	payload := new(bytes.Buffer)
	if err := json.NewEncoder(payload).Encode(data); err != nil {
		return nil, err
	}
	return s.response("PUT", fmt.Sprintf("%s/%s", strings.TrimRight(apiURL, "/"), resource), payload)
}

// DeleteResponse returns an http.Response object for the deleted resource
func (s *ScalewayAPI) DeleteResponse(apiURL, resource string) (*http.Response, error) {
	return s.response("DELETE", fmt.Sprintf("%s/%s", strings.TrimRight(apiURL, "/"), resource), nil)
}