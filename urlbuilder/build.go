package urlbuilder

import (
	"log"
	"net/url"
)

func Build(host string) *url.URL {
	u, err := url.Parse("https://example.org")
	if err != nil {
		log.Fatal(err)
	}

	u.Scheme = "https"

	//if host == "" {
	//	host = "chleb-api.daybologic.co.uk"
	//}
	u.Host = host

	u.Path = "2/votd"

	q := u.Query()
	q.Set("translations", "asv")
	u.RawQuery = q.Encode()

	return u
}
