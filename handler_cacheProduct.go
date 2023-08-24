package main

import (
	"io/ioutil"
	"os"
)

func handlerCacheProduct(productName string) ([]byte, bool, error) {
	content, err := ioutil.ReadFile("products/" + productName + ".json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil, false, nil
		}
		return nil, false, err
	}

	return content, true, nil
}
