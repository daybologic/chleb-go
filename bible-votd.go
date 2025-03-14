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
	hostname = flag.String("h", "chleb-api.daybologic.co.uk", "The hostname for the remote Chleb Bible Search server")
	htmlFlag = flag.Bool("t", false, "Use text/html")
)

func fetch(respond chan<- string, query string) {
	response, ok := remote.Fetch(query)
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
	//hostname := "chleb-api.daybologic.co.uk"
	args := flag.Args()
	if len(args) >= 2 {
		usage()
	}
	//if len(args) >= 1 {
	//	hostname = args[0]
	//}
	if *hostname == "" {
		log.Fatalf("invalid hostname %q", hostname)
	}

	// Run actual logic.
	//if *reverseFlag {
	//	fmt.Printf("%s, %s!\n", reverse.String(*greeting), reverse.String(name))
	//	return
	//}
	//fmt.Printf("%s, %s!\n", *greeting, name)

	query := urlbuilder.Build(*hostname).String()
	fmt.Printf("URL '%s'\n", query);

	respond := make(chan string)

	go fetch(respond, query)

	queryResp := <-respond

	fmt.Printf("Sent query:\t\t %s\n", query)
	fmt.Printf("Got Response:\t\t %s\n", queryResp)
}
