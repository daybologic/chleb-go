package remote

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
)

type JSONResponseAttributes struct {
	text string `json:"text"`
}

type JSONResponseData struct {
	attributes JSONResponseAttributes `json:"attributes"`
}

type JSONResponseLinks struct {
	prev string `json:"prev"`
	self string `json:"self"`
	next string `json:"next"`
}

type JSONResponse struct {
	data [1]JSONResponseData `json:"data"`
	links JSONResponseLinks `json:"links"`
}

func Fetch(query string, htmlFlag bool) (response string, ok bool) {
	// TODO: You need cookies or you'll be subject to the speed limit
	req, err := http.NewRequest("GET", query, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", false
	}

	req.Header.Set("User-Agent", "chleb-go/experimental")

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

	var cookedBody string
	if (htmlFlag) {
		cookedBody = fmt.Sprintf("%s", rawBody)
	} else {
		var jsonResponse JSONResponse
		err = json.Unmarshal([]byte(rawBody), &jsonResponse)
		if err != nil {
			log.Fatalf("Unable to marshal JSON due to %s", err)
		}
		//cookedBody = fmt.Sprintf("%s", jsonResponse.data[0].attributes.text)
		cookedBody = fmt.Sprintf("%s", jsonResponse)
		cookedBody = jsonResponse.links.self
	}

	return cookedBody, true
}
