// Copyright © 2016 Ladislav Prskavec <ladislav@prskavec.net>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"encoding/json"
	"io"
	// "io/ioutil"
	"os"
	"github.com/spf13/cobra"
	"net/http"
)

var subdomain, token string

type ClientResponse struct {
	ErrorCode 	bool `json:"error,omitempty"`
	Message 		string `json:"message,omitempty"`
	Code 				string `json:"code"`
}

func Decode(r io.Reader) (x *ClientResponse, err error) {
    x = new(ClientResponse)
    err = json.NewDecoder(r).Decode(x)
    return
}

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		token = os.Getenv("APIARY_API_KEY")
		if (len(args)>0) {
			subdomain = args[0]
		}
		if (len(token)>0) && (len(subdomain)>0) {
			client := &http.Client{}
			req, _ := http.NewRequest("GET", "https://api.apiary.io/blueprint/get/"+subdomain, nil)

			req.Header.Set("accept", "text/html")
			req.Header.Set("content_type", "text/plain")
			req.Header.Set("authentication", "Token "+token)
			req.Header.Set("user_agent", "User Agent golang")
			res, _ := client.Do(req)
			output, err := Decode(res.Body)
			if(err != nil){
	        fmt.Println("whoops:", err)
	    }
			fmt.Println(output.Code)
			res.Body.Close()
		}
	},
}

func init() {
	RootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("subdomain", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
