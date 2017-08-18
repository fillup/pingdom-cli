package main

import (
	"net/http"
	"os"
)

// NewHTTPRequest - Create a new HTTP Request object and attach auth info
func NewHTTPRequest(url string, method string) *http.Request {
	req, _ := http.NewRequest(method, url, nil)
	req.SetBasicAuth(os.Getenv("PINGDOM_USER"), os.Getenv("PINGDOM_PASS"))
	req.Header.Set("App-Key", os.Getenv("PINGDOM_APPKEY"))

	return req
}

// CallAPI - Call API
func CallAPI(url string, method string) (resp *http.Response, err error) {
	req := NewHTTPRequest(url, method)
	client := &http.Client{}
	resp, err = client.Do(req)
	return
}
