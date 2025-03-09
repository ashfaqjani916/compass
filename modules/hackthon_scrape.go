package modules

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
	"web-scraper/config"

	"github.com/chromedp/chromedp"
)



func fetchJob(jobs <-chan string, wg *sync.WaitGroup) {
    defer wg.Done()
ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 40*time.Second)
	defer cancel()
    for url := range jobs {


	var names []string

	err := chromedp.Run(ctx,

chromedp.Navigate(url),
		chromedp.WaitVisible(`h3`, chromedp.ByQuery), // Wait for hackathon titles
		chromedp.Sleep(5*time.Second),                // Wait for JS execution
		chromedp.Evaluate(`Array.from(document.querySelectorAll("h3")).map(e => e.innerText)`, &names),
			)

	if err != nil {
		log.Fatal("Scraping failed:", err)
	}


	fmt.Printf("\n=== Open Hackathons from %s  ===",url)
	for i := range names {
		fmt.Println(names[i])
	}
    }
}



func GetHackathonData() {
	sources := config.FetchSource()
	numJobs := len(sources)
	jobs := make(chan string, numJobs)
	var wg sync.WaitGroup
	numWorkers := 13
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1) 
		go fetchJob(jobs, &wg)
	}
	for _,url := range sources {
		jobs <- url
	}
	close(jobs)
	wg.Wait()
	fmt.Println("Scraping finished!")
}
