package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var subdomain, token, apiHost, apiKey string

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

func CommonHeaders(req http.Request, token string, contentType string) {
	req.Header.Set("Accept", "text/html")
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authentication", "Token "+token)
	req.Header.Set("User-Agent", "bee / golang")
	return
}

func Fetch(subdomain string, token string, apiHost string) {
	if len(token) == 0 {
		fmt.Println("Missing APIARY_API_KEY")
		os.Exit(1)
	}

	if len(subdomain) > 0 {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", apiHost+"/blueprint/get/"+subdomain, nil)

		CommonHeaders(*req, token, "text/plain")

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
	if len(token) == 0 {
		fmt.Println("Missing APIARY_API_KEY")
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(filename)

	if err == nil {
		if len(subdomain) > 0 {
			client := &http.Client{}

			form := url.Values{}
			form.Add("code", string(data))
			form.Add("messageToSave", "Update")

			req, _ := http.NewRequest("POST", apiHost+"/blueprint/publish/"+subdomain, bytes.NewBufferString(form.Encode()))
			req.PostForm = form
			CommonHeaders(*req, token, "application/x-www-form-urlencoded")
			res, _ := client.Do(req)
			_, error := Decode(res.Body)
			if err != nil {
				fmt.Println("whoops:", error)
			}
			defer res.Body.Close()
			fmt.Println("Response Status:", res.Status)
			fmt.Println("Response Headers:", res.Header)
		}
	} else {
		fmt.Println("whoops:", err)
	}

}
