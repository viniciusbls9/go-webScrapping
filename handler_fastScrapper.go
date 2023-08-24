package main

import (
	"encoding/json"

	"github.com/gocolly/colly"
)

func handlerFastScrapper(c *colly.Collector, productName string) ([]byte, []item, error) {
	var items []item

	c.OnHTML("div.ui-search-result--core", func(h *colly.HTMLElement) {
		item := item{
			Name: h.ChildText("h2.ui-search-item__title"),
			FromPrice: h.ChildText(".ui-search-price__original-value .andes-money-amount__fraction"),
			ToPrice: h.ChildText("span.ui-search-price__part--medium span.andes-money-amount__fraction"),
			ImgUrl: h.ChildAttr("img", "src"),
			ProductUrl: h.ChildAttr("a.ui-search-link", "href"),
		}

		items = append(items, item)
	})

	err := c.Visit("https://lista.mercadolivre.com.br/"+productName)
	if err != nil {
		return nil, nil, err
	}
	
	content, err := json.Marshal(items)
	if err != nil {
		return nil, nil, err
	}

	return content, items, nil
}