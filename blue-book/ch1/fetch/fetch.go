// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("Fetching...")
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		fmt.Printf("HTTP Status: %v\n", resp.Status)
		fmt.Printf("%v\n", err)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		fmt.Printf("%v\n", b)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:\nreading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
