package webfunctions

import (
	"io"
	"log"
	"net/http"
	"os"
)

func getWeb(siteurl string) {
	response, err := http.Get(siteurl)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	// copy data from the response to the standard output
	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Number of bytes copied to STDOUT:", n)
}
