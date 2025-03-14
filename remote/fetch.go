package remote

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
)

func Fetch(query string) (response string, ok bool) {
	resp, err := http.Get(query)
	if err != nil {
		log.Fatal(err)
		return "", false
	}

	defer resp.Body.Close()

	rawBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return "", false
	}

	cookedBody := fmt.Sprintf("%s", rawBody)
	//decodedRawBody, err := b64.StdEncoding.DecodeString(cookedBody)
	//if err != nil {
	//	log.Fatal(err)
	//	return "", false
	//}

	return cookedBody, true
}
