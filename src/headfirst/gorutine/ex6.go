package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

/**
  ex1 개선
*/
func responseSize2(url string, channel chan Page) {
	fmt.Println("Greeting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	channel <- Page{URL: url, Size: len(body)}
}

func main() {
	sizes := make(chan Page)
	urls := []string{"https://example.com/", "https://golang.org/", "https://golang.org/doc"}
	for _, urls := range urls {
		go responseSize2(urls, sizes)
	}

	for i := 0; i < len(urls); i++ {
		page := <-sizes
		fmt.Printf("%s : %d\n", page.URL, page.Size)
	}

}
