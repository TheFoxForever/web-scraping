package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

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

func main() {
	fmt.Println("working")
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org", "en.wikipedia.org/wiki"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
		// e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	for i := 0; i < len(urls); i++ {
		c.Visit(urls[i])
	}
}
