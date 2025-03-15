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
	htmlFlag = flag.Bool("h", true, "Use text/html")
	translations = flag.String("t", "asv", "The translation(s) required (asv, kjv)")
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

	query := urlbuilder.Build(*insecureFlag, *hostname, *port, *translations).String()

	respond := make(chan string)

	go fetch(respond, query)

	queryResp := <-respond

	fmt.Printf("%s\n", queryResp)
}
