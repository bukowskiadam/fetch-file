package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Missing URL parameter!")
	}

	if len(os.Args) < 3 {
		log.Fatalln("Missing output filename parameter!")
	}

	for retries := 3; retries > 0; retries -= 1 {
		resp, err := http.Get(os.Args[1])

		if err != nil {
			os.Exit(1)
		}

		if resp.StatusCode != 200 {
			continue
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			os.Exit(1)
		}

		os.WriteFile(os.Args[2], body, 0644)

		nextRefresh := resp.Header.Get("X-Next-Refresh")

		if len(nextRefresh) > 0 {
			fmt.Println(nextRefresh)
		}

		os.Exit(0)
	}

	os.Exit(1)
}
