package main

import (
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

	resp, err := http.Get(os.Args[1])

	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	os.WriteFile(os.Args[2], body, 0644)
}
