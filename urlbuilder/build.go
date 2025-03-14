package urlbuilder

import (
	"log"
	"net"
	"net/url"
	"strconv"
)

func Build(insecure bool, host string, port int) *url.URL {
	u, err := url.Parse("https://example.org")
	if err != nil {
		log.Fatal(err)
	}

	u.Scheme = "https"
	if (insecure) {
		u.Scheme = "http"
	}

	u.Host = host
	u.Path = "2/votd"

	if port > 0 { // override, non-standard
		if (u.Scheme == "https" && port != 443) || (u.Scheme == "http" && port != 80) {
			u.Host = net.JoinHostPort(host, strconv.Itoa(port))
		}
	}

	q := u.Query()
	q.Set("translations", "asv")
	u.RawQuery = q.Encode()

	return u
}
