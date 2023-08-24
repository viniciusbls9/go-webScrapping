package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gocolly/colly"
)

type item struct {
	Name       string `json:"name"`
	FromPrice  string `json:"fromPrice"`
	ToPrice    string `json:"toPrice"`
	ImgUrl     string `json:"imgurl"`
	ProductUrl string `json:"producturl"`
}

func scrapper(responseWritter http.ResponseWriter, request *http.Request) {
	productName := chi.URLParam(request, "productName")
	c := colly.NewCollector()

	content, exists, err := handlerCacheProduct(productName)
	if err != nil {
		respondWithError(responseWritter, 500, fmt.Sprintf("Error checking JSON file: %v", err))
		return
	}

	if exists {
		var items []item
		if err := json.Unmarshal(content, &items); err != nil {
			respondWithError(responseWritter, 500, fmt.Sprintf("Error unmarshaling JSON content: %v", err))
			return
		}
		respondWithJSON(responseWritter, 200, items)
		return
	}

	content, items, err := handlerGirafaScrapper(c, productName)

	if err != nil {
		respondWithError(responseWritter, 400, fmt.Sprintf("Couldn't scrappe product: %v", err))
	}

	handlerSaveProducts(responseWritter, request, content, productName)
	respondWithJSON(responseWritter, 200, items)

}
