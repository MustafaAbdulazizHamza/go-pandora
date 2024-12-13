package goPandora

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func NewPandoraClient(url, username, password, privateKey, publicKey string) *PandoraClient {
	return &PandoraClient{
		url:        url,
		username:   username,
		password:   password,
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}
func sendHTTPRequest(method, url, username, password string, body string) (response *http.Response, err error) {
	req, err := http.NewRequest(method, url, bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("username", username)
	req.Header.Set("password", password)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // This skips certificate verification
		},
	}

	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: tr,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func toJson(i interface{}) ([]byte, error) {
	j, err := json.Marshal(i)
	return j, err
}
func ParseResponse(inp *http.Response) (response Response, err error) {
	defer inp.Body.Close()
	body, err := ioutil.ReadAll(inp.Body)
	if err != nil {
		return response, err
	}
	if string(body) == "404 page not found" {
		body, err = toJson(Response{
			Status: fmt.Sprint(http.StatusNotFound),
			Text:   "Page not found.",
		})
	}
	err = json.Unmarshal(body, &response)

	return response, err
}
