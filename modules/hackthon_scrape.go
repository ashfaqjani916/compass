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
//
// var mu sync.Mutex // Mutex to prevent race conditions while writing to CSV
//
// // fetchJob scrapes job data from given URLs and writes it to a CSV file
// func fetchJob(jobs <-chan string, wg *sync.WaitGroup, csvFile *os.File) {
// 	defer wg.Done()
//
// 	writer := csv.NewWriter(csvFile)
//
// 	for url := range jobs {
// 		res, err := http.Get(url)
// 		if err != nil {
// 			log.Printf("Error while fetching data from %s: %v", url, err)
// 			continue
// 		}
// 		defer res.Body.Close()
//
// 		doc, err := goquery.NewDocumentFromReader(res.Body)
// 		if err != nil {
// 			log.Printf("Error while parsing data from %s: %v", url, err)
// 			continue
// 		}
//
// 		// Find relevant job data
// 		doc.Find("*").Each(func(index int, element *goquery.Selection) {
// 			tag := element.Nodes[0].Data
// 			attrs := ""
// 			for _, attr := range element.Nodes[0].Attr {
// 				attrs += fmt.Sprintf(` %s="%s"`, attr.Key, attr.Val)
// 			}
//
// 			text := strings.TrimSpace(element.Text())
//
// 			// Write to CSV with thread safety
// 			mu.Lock()
// 			writer.Write([]string{url, tag, attrs, text})
// 			mu.Unlock()
// 		})
// 	}
//
// 	writer.Flush() // Ensure all data is written
// }

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

		// sc-hKMtZM sc-jSMfEi hackathonStatus-__StyledGrid-sc-347abc02-2 cwGkJu kGTbcK iGtZjp
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

// csvFile, err := os.Create("jobs.csv")
// 	if err != nil {
// 		log.Fatalf("Could not create CSV file: %v", err)
// 	}
// 	defer csvFile.Close()
//
// 	// Write CSV headers
// 	writer := csv.NewWriter(csvFile)
// 	writer.Write([]string{"URL", "Tag", "Attributes", "Text Content"})
// 	writer.Flush()
//
// 	for _, ele := range sources {
// 		fmt.Println(ele)
// 		jobs <- ele
// 	}
// 	close(jobs) 
//
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
