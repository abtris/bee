package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var subdomain, token, apiHost, apiKey string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type ClientResponse struct {
	ErrorCode bool   `json:"error,omitempty"`
	Message   string `json:"message,omitempty"`
	Code      string `json:"code"`
}

func Decode(r io.Reader) (x *ClientResponse, err error) {
	x = new(ClientResponse)
	err = json.NewDecoder(r).Decode(x)
	return
}

func commonHeaders(req http.Request, token string) {
	req.Header.Set("accept", "text/html")
	req.Header.Set("content_type", "text/plain")
	req.Header.Set("authentication", "Token "+token)
	req.Header.Set("user_agent", "User Agent golang")
	return
}

func Fetch(subdomain string, token string, apiHost string) {

	if len(subdomain) > 0 {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", apiHost+"/blueprint/get/"+subdomain, nil)

		commonHeaders(*req, token)

		res, _ := client.Do(req)
		output, err := Decode(res.Body)
		if err != nil {
			fmt.Println("whoops:", err)
		}
		fmt.Println(output.Code)
		res.Body.Close()
	}
}

func Publish(subdomain string, filename string, token string, apiHost string) {

	data, err := ioutil.ReadFile(filename)
	raw := bytes.NewReader(data)

	if err != nil {
		if len(subdomain) > 0 {
			client := &http.Client{}
			req, _ := http.NewRequest("POST", apiHost+"/blueprint/publish/"+subdomain, raw)

			commonHeaders(*req, token)

			res, _ := client.Do(req)
			output, err := Decode(res.Body)
			if err != nil {
				fmt.Println("whoops:", err)
			}
			fmt.Println(output.Code)
			res.Body.Close()
		}
	}

}
