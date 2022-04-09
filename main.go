package main
import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

type info struct {
	status 	string
	sensex 	string
	nifty	string
}

func main() {
	link := "https://economictimes.indiatimes.com/markets"
	c := colly.NewCollector(
		colly.CacheDir("./cache"),
	)
	
	c.OnHTML("div.content_area",func(e *colly.HTMLElement) {
		i	:= info {
			status:	e.DOM.Find("span.mktStatus").Text(),
			sensex: e.DOM.Find("div[data-pos='1']").Find("span.change").Text(),
			nifty:	e.DOM.Find("div[data-pos='2']").Find("span.change").Text(),
		}

		fmt.Println("Status: ", i.status)
		fmt.Println("SENSEX: ", i.sensex)
		fmt.Println("NIFTY: ", 	i.nifty)
	})
	
	c.Visit(link)
}