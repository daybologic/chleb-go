/*
Chleb Bible Search
Copyright (c) 2024-2025, Rev. Duncan Ross Palmer (M6KVM, 2E0EOL),
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

    * Redistributions of source code must retain the above copyright notice,
      this list of conditions and the following disclaimer.

    * Redistributions in binary form must reproduce the above copyright
      notice, this list of conditions and the following disclaimer in the
      documentation and/or other materials provided with the distribution.

    * Neither the name of the Daybo Logic nor the names of its contributors
      may be used to endorse or promote products derived from this software
      without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE
LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
POSSIBILITY OF SUCH DAMAGE.
*/

package remote

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"strings"
)

type JSONResponseAttributes struct {
	Book string `json:"book"`
	Chapter json.Number `json:"chapter"`
	Ordinal json.Number `json:"ordinal"`
	Text string `json:"text"`
}

type JSONResponseData struct {
	Attributes JSONResponseAttributes `json:"attributes"`
}

type JSONResponseLinks struct {
	Prev string `json:"prev"`
	Self string `json:"self"`
	Next string `json:"next"`
}

type JSONResponse struct {
	Data []JSONResponseData `json:"data"`
	Links JSONResponseLinks `json:"links"`
}

func Fetch(query string, htmlFlag bool) (response string, ok bool) {
	// TODO: You need cookies or you'll be subject to the speed limit
	req, err := http.NewRequest("GET", query, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", false
	}

	req.Header.Set("User-Agent", "chleb-go/0.1.0")

	if (htmlFlag) {
		req.Header.Set("Accept", "text/html")
	} else {
		req.Header.Set("Accept", "application/json")
	}

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", false
	}

	defer resp.Body.Close()

	rawBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return "", false
	}

	var cookedBody strings.Builder
	if (htmlFlag) {
		cookedBody.WriteString(fmt.Sprintf("%s", rawBody))
	} else {
		var jsonResponse JSONResponse
		err = json.Unmarshal([]byte(rawBody), &jsonResponse)
		if err != nil {
			log.Fatalf("Unable to marshal JSON due to %s", err)
		}
		for _,element := range jsonResponse.Data {
			var attribs JSONResponseAttributes = element.Attributes
			cookedBody.WriteString(fmt.Sprintf(
				"%s %s:%s %s\n",
				attribs.Book,
				attribs.Chapter,
				attribs.Ordinal,
				attribs.Text,
			))
		}
	}

	return cookedBody.String(), true
}
