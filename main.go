package main

import (
	"os"
	"fmt"
	"github.com/anaskhan96/soup"
)

func main () {
	URL := "https://economictimes.indiatimes.com/markets"

	html, err := soup.Get(URL)
	if err != nil {
    		os.Exit(1)
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
