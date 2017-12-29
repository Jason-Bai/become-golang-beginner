package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/jackdanger/collectlinks"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Please specify start page")
		os.Exit(1)
	}

	retrieve(args[0])
}

func retrieve(uri string) {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	client := http.Client{Transport: transport}

	resp, err := client.Get(uri)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)

	// fmt.Println(string(body))

	links := collectlinks.All(resp.Body)

	for _, link := range links {
		fmt.Println(link)
	}
}
