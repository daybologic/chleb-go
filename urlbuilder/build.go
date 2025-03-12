package urlbuilder

import (
	"log"
	"net/url"
)

func Build() *url.URL {
	u, err := url.Parse("https://example.org")
	if err != nil {
		log.Fatal(err)
	}

	u.Scheme = "https"
	u.Host = "chleb-api.daybologic.co.uk"
	u.Path = "2/votd"

	q := u.Query()
	q.Set("translations", "asv")
	u.RawQuery = q.Encode()

	return u
}
