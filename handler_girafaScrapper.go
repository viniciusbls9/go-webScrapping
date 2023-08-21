package main

import (
	"encoding/json"

	"github.com/gocolly/colly"
)

func handlerGirafaScrapper(c *colly.Collector, productName string) ([]byte, []item, error) {
	var items []item

	c.OnHTML("div.bloco-produto", func(h *colly.HTMLElement) {
		item := item{
			Name: h.ChildText("a.informacao-produto h2"),
			FromPrice: h.ChildText("div.valores span.preco-de"),
			ToPrice: h.ChildText("div.valores span.valor-vista"),
			ImgUrl: h.ChildAttr("div.imagens-list-produto img", "src"),
			ProductUrl: h.ChildAttr("a", "href"),
		}

		items = append(items, item)
	})

	err := c.Visit("https://www.girafa.com.br/busca/?q="+productName)
	if err != nil {
		return nil, nil, err
	}
	
	content, err := json.Marshal(items)
	if err != nil {
		return nil, nil, err
	}

	return content, items, nil
}