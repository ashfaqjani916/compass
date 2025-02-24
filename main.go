
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	source :="https://gobyexample.com/" 

	results := make(chan string, 25)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		res, err := http.Get(source)
		if err != nil {
			log.Println("error while fetching data:", err)
			return
		}
		defer res.Body.Close() 

		data, readError := goquery.NewDocumentFromReader(res.Body)
		if readError != nil {
			log.Println("error while parsing data:", readError)
			return
		}

		data.Find("a").Each(func(index int, element *goquery.Selection) {
			link, exists := element.Attr("href")
			if exists {
				results <- fmt.Sprintf("URL: %s → Link: %s", source, link)
			}
		})

		close(results)
	}()

	for url := range results {
		fmt.Println(url)
	}

	wg.Wait()
}
