package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type item struct {
	Name 		string `json:"name"`
	Price 	string `json:"price"`
	ImgUrl 	string `json:"imgurl"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("lista.mercadolivre.com.br"),
	)

	c.OnHTML("div[class=ui-search-result__content-wrapper]", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

}