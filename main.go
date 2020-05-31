package main

import (
	"flag"
	"fmt"
	"github.com/Comic-Ghost/client"
	"log"
	"time"
)

// Reads command line arguments
// Instantiate XKCDClient
// Fetch from API using XKCDClient
// Output

func main() {
	comicNumber := flag.Int(
		"num", int(client.LatestComic), "Comic number to fetch (default latest)",
	)

	clientTimeout := flag.Int64(
		"timeout", int64(client.DefaultClientTimeout.Seconds()), "Client Timeout in seconds",
	)

	saveImage := flag.Bool(
		"save", false, "Save image to current directory",
	)

	outputType := flag.String(
		"output", "text", "Print output in format: text/json",
	)

	flag.Parse()

	xkcdClient := client.NewXKCDClient()
	xkcdClient.SetTimeout(time.Duration(*clientTimeout) * time.Second)

	comic, err := xkcdClient.Fetch(client.ComicNumber(*comicNumber), *saveImage)
	if err != nil {
		log.Println(err)
	}

	if *outputType == "json" {
		fmt.Println(comic.JSON())
	} else {
		fmt.Println(comic.PrettyString())
	}
}
