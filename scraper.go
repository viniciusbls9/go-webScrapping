package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type item struct {
	Name 				string `json:"name"`
	Price 			string `json:"price"`
	ImgUrl 			string `json:"imgurl"`
	ProductUrl 	string `json:"producturl"`
}

func scrapper() {
	c := colly.NewCollector(
		colly.AllowedDomains("lista.mercadolivre.com.br"),
	)

	var items []item

	c.OnHTML("div.ui-search-result--core", func(h *colly.HTMLElement) {
		item := item{
			Name: h.ChildText("h2.ui-search-item__title"),
			Price: h.ChildText("div.ui-search-item__group__element span.andes-money-amount"),
			ImgUrl: h.ChildAttr("img", "src"),
			ProductUrl: h.ChildAttr("a.ui-search-link", "href"),
		}

		items = append(items, item)
	})

	c.Visit("https://lista.mercadolivre.com.br/iphone-13s-pro-max")
	
	content, err := json.Marshal(items)

	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("products.json", content, 0644)

}