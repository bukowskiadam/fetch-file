package main

import (
	"fmt"
	"io/ioutil"
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

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := file.Write(body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Wrote %d bytes\n", bytes)
}
