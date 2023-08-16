package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gocolly/colly"
)

type item struct {
	Name 				string `json:"name"`
	Price 			string `json:"price"`
	ImgUrl 			string `json:"imgurl"`
	ProductUrl 	string `json:"producturl"`
}

func scrapper(responseWritter http.ResponseWriter, request *http.Request) {
	productName := chi.URLParam(request, "productName")
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

	c.Visit("https://lista.mercadolivre.com.br/"+productName)
	
	content, err := json.Marshal(items)

	if err != nil {
		respondWithError(responseWritter, 400, fmt.Sprintf("Couldn't scrappe product: %v", err))
	}

	handlerSaveProducts(responseWritter, request, content, productName)
	respondWithJSON(responseWritter, 200, items)

}