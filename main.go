package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/gocolly/colly"
)

type UrlJSON struct {
	URL  string `json:"url"`
	Text string `json:"text"`
}

// validate that all URLs provided are on the allowed domain
func isValidURL(urlStr string, allowedDomains []string) bool {
	url, err := url.Parse(urlStr)
	if err != nil {
		return false
	}

	if url.Scheme != "https" {
		fmt.Printf("Wrong protocol: %s (must be https)", url.Scheme)
		return false
	}

	if url.Scheme == "" || url.Host == "" {
		return false
	}

	for _, domain := range allowedDomains {
		fmt.Println(url.Host)
		if url.Host == domain {
			return true
		}
	}

	return false
}

func main() {

	var urls = []string{
		"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robot",
		"https://en.wikipedia.org/wiki/Reinforcement_learning",
		"https://en.wikipedia.org/wiki/Robot_Operating_System",
		"https://en.wikipedia.org/wiki/Intelligent_agent",
		"https://en.wikipedia.org/wiki/Software_agent",
		"https://en.wikipedia.org/wiki/Robotic_process_automation",
		"https://en.wikipedia.org/wiki/Chatbot",
		"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
		"https://en.wikipedia.org/wiki/Android_(robot)",
	}

	allowedDomains := []string{"en.wikipedia.org"}

	for _, url := range urls {
		if isValidURL(url, allowedDomains) {
			continue
		} else {
			fmt.Println("Invalid:", url)
		}
	}

	c := colly.NewCollector(
		colly.AllowedDomains(allowedDomains...),
	)

	file, err := os.Create("output.jl")
	if err != nil {
		log.Fatal("Could not create output file", err)
	}
	defer file.Close()

	c.OnHTML("#content #bodyContent p", func(e *colly.HTMLElement) {
		data := UrlJSON{
			URL:  e.Request.URL.String(),
			Text: e.Text,
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Printf("JSON encoding failed: %v", err)
			return
		}
		fmt.Fprintln(file, string(jsonData))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	for i := 0; i < len(urls); i++ {
		c.Visit(urls[i])
	}
}
