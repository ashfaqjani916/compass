package modules

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"web-scraper/config"

	"github.com/PuerkitoBio/goquery"
)

func fetchJob(jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for url := range jobs {
		res, err := http.Get(url)
		if err != nil {
			log.Printf("Error while fetching data from %s: %v", url, err)
			continue
		}
		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Printf("Error while parsing data from %s: %v", url, err)
			continue 
		}

		doc.Find("a").Each(func(index int, element *goquery.Selection) {
			tag := element.Nodes[0].Data
			attrs := ""
			for _, attr := range element.Nodes[0].Attr {
				attrs += fmt.Sprintf(` %s="%s"`, attr.Key, attr.Val)
			}

			text := strings.TrimSpace(element.Text())
			fmt.Printf("<%s%s> %s\n", tag, attrs, text)
		})
	}
}

func GetHackathonData() {
	sources := config.FetchSource()
	numJobs := len(sources)

	jobs := make(chan string, numJobs)

	for _, ele := range sources {
		fmt.Println(ele)
		jobs <- ele
	}
	close(jobs) 

	var wg sync.WaitGroup

	// w signifies the number of goroutines at a time, 
	// generally the range can be btw 2x to 5x for x number of cores in a processor
	for w := 1; w <= 13; w++ {
		wg.Add(1) 
		go fetchJob(jobs, &wg)
	}

	wg.Wait()

	fmt.Println("Scraping finished!")
}
