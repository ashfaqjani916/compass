package utils

import (
"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func fetchHTML(url string) (*http.Response, error)  {
	resp, err := http.Get(url)
	if err!=nil {
		return nil,err 
	}

if !strings.Contains(resp.Header.Get("Content-Type"), "charset=utf-8") {
		return nil, fmt.Errorf("non-UTF-8 content received")
	}
	return resp, nil	
}



func parseHTML(resp *http.Response) ([]string, []string, error) {
	// Parse the HTML content using the `html` package
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	var titles []string
	var links []string

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		// Traverse through the HTML nodes to extract links to hackathons
		if n.Type == html.TextNode{
			// Check if the link contains a hackathon-related URL
			for _, attr := range n.Attr {
				fmt.Println(attr)
			}
		}
		// Recursively traverse child nodes
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	// Start traversing from the root node
	traverse(doc)
	return titles, links, nil
}


func Data()  {
	// URL of the open hackathons page
	url := "https://devfolio.co/hackathons/open"

	// Fetch the HTML data
	resp, err := fetchHTML(url)
	if err != nil {
		fmt.Println("Error fetching HTML:", err)
		return
	}
	defer resp.Body.Close()

	// Parse the HTML and extract titles and links of the hackathons
	titles, links, err := parseHTML(resp)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}

	fmt.Println("fetching done")

	// Print the titles and links of the hackathons
	for i, title := range titles {
		fmt.Printf("Hackathon: %s\nLink: %s\n\n", title, links[i])
	}
}
