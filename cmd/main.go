package main

import (
	"fmt"
	"net/url"
	"os"

	"go-metadata"
)

func main() {
	metadataBaseURL, err := url.Parse("http://169.254.169.254")
	if err != nil {
		panic(err)
	}

	meta := metadata.NewClient(metadata.WithBaseURL(metadataBaseURL))
	all, err := meta.Metadata()

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fip := all.FloatingIP.IPv4.IPAddress
	fmt.Println(fip)
}
