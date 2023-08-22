package main

import (
	"encoding/json"

	"github.com/gocolly/colly"
)

func handlerFastScrapper(c *colly.Collector, productName string) ([]byte, []item, error) {
	var items []item

	c.OnHTML("div.category-list", func(h *colly.HTMLElement) {
		item := item{
			Name: h.ChildText("h3.horizontal-grid-short-description"),
			FromPrice: h.ChildText("span.value"),
			ToPrice: h.ChildText("span.price-fraction"),
			ImgUrl: h.ChildAttr("img.grid-image", "src"),
			ProductUrl: h.ChildAttr("a", "href"),
		}

		items = append(items, item)
	})

	err := c.Visit("https://www.fastshop.com.br/web/s/"+productName)
	if err != nil {
		return nil, nil, err
	}
	
	content, err := json.Marshal(items)
	if err != nil {
		return nil, nil, err
	}

	return content, items, nil
}