// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Hello is a hello, world program, demonstrating
// how to write a simple command-line program.
//
// Usage:
//
//	hello [options] [name]
//
// The options are:
//
//	-g greeting
//		Greet with the given greeting, instead of "Hello".
//
//	-r
//		Greet in reverse.
//
// By default, hello greets the world.
// If a name is specified, hello greets that name instead.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/example/bible-votd/remote"
	"golang.org/x/example/bible-votd/urlbuilder"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: hello [options] [name]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

var (
	hostname = flag.String("H", "chleb-api.daybologic.co.uk", "The hostname for the remote Chleb Bible Search server")
	port = flag.Int("p", 0, "connect to port, default depends whether -k is specified");
	insecureFlag = flag.Bool("k", false, "connect via legacy HTTP (insecure)");
	htmlFlag = flag.Bool("h", false, "Use text/html")
)

func fetch(respond chan<- string, query string) {
	response, ok := remote.Fetch(query, *htmlFlag)
	if !ok {
		os.Exit(1)
	}

	respond <- response
}

func main() {
	// Configure logging for a command-line program.
	log.SetFlags(0)
	log.SetPrefix("bible-votd: ")

	// Parse flags.
	flag.Usage = usage
	flag.Parse()

	// Parse and validate arguments.
	args := flag.Args()
	if len(args) >= 2 {
		usage()
	}

	if *hostname == "" {
		log.Fatalf("invalid hostname: %q", *hostname)
	}

	query := urlbuilder.Build(*insecureFlag, *hostname, *port).String()
	fmt.Printf("URL '%s'\n", query);

	respond := make(chan string)

	go fetch(respond, query)

	queryResp := <-respond

	fmt.Printf("Sent query:\t\t %s\n", query)
	fmt.Printf("Got Response:\t\t %s\n", queryResp)
}
