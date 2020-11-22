package test

import (
	"net/http"
	"encoding/json"
	"strings"
	"io/ioutil"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{httpClient: &http.Client{}}
}

type SlowRequest struct {
	Timeout int
}

type SlowResponse struct {
	Status string `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
	Code   int    `json:"code,omitempty"`
}

func (cl *Client) SlowApiPost(timeout int) (*SlowResponse, error) {
	req := SlowRequest{Timeout: timeout}
	m, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	r := strings.NewReader(string(m))
	resp, err := cl.httpClient.Post("http://localhost:8080/api/slow", "application/json", r)
	if err != nil {
		return nil, err
	}
	var slowResp SlowResponse
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &slowResp)
	if err != nil {
		return nil, err
	}
	slowResp.Code = resp.StatusCode
	return &slowResp, nil
}
