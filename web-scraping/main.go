package main

import (
	"fmt"
	"io"
	// "log"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"os"
	"strings"
	// "time"
)

var (
	fileName    string
	fullURLFile string
)

func main() {
	fullURLFile = "https://apod.nasa.gov/apod/ap190604.html"
	for i := 605; i > 500; i-- {

		fullURLFile = findIMGs()
		// Build fileName from full URL
		buildFileName()

		// Create blank file
		file := createFile()

		// Put content in file
		putFile(file, httpClient())

		// Go to next file
		fullURLFile = "https://apod.nasa.gov/apod/ap190" + string(i) + ".html"
	}
}

func putFile(file *os.File, client *http.Client) {
	resp, err := client.Get(fullURLFile)
	checkError(err)
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)
	defer file.Close()
	checkError(err)

	fmt.Println("Downloaded file...", fileName, size)
}

func buildFileName() {
	fileURL, err := url.Parse(fullURLFile)
	checkError(err)

	path := fileURL.Path
	segments := strings.Split(path, "/")

	fileName = segments[len(segments)-1]
}

func createFile() *os.File {
	file, err := os.Create(fileName)
	checkError(err)
	return file
}

func httpClient() *http.Client {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	return &client
}

func findIMGs() string {
	var ans string

	// Make HTTP request
	response, err := http.Get(fullURLFile)
	checkError(err)
	defer response.Body.Close()

	// Create a goquery doc from HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	checkError(err)

	// Find and print image URLs
	document.Find("img").Each(func(index int, element *goquery.Selection) {
		imgSrc, exists := element.Attr("src")
		if exists {
			ans = imgSrc
		}
	})
	return "https://apod.nasa.gov/" + ans
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
