package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"time"
)

func scrape () {
	URL := "https://economictimes.indiatimes.com/markets"

	html, err := soup.Get(URL)
	if err != nil {
    		panic(err)
	}

	DOC 		:= soup.HTMLParse(html)
	STATUS 		:= DOC.Find("span", "class", "mktStatus").Text()
	//TIME 		:= DOC.Find("span", "class", "date_format active").Text()
	INDICES 	:= DOC.Find("div", "class", "indices")
	SENSEX_DIFF 	:= INDICES.Find("div", "data-pos", "1").Find("span", "class", "change").Text()
	NIFTY_DIFF 	:= INDICES.Find("div", "data-pos", "2").Find("span", "class", "change").Text()

	fmt.Println("Status: ", STATUS)
	fmt.Println("Sensex: ", SENSEX_DIFF )
	fmt.Println("Nifty: ", NIFTY_DIFF)
}

func main () {
    for {
    scrape()

    time.Sleep(60 * time.Second)

    for j := 0; j < 3; j++ {
        fmt.Print("\033[1A\033[K")
    }
    }
}
