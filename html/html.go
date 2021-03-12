package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {

	u, err := url.Parse("https://dev-admin-olimpicos.marcaclaro.com/noticias/thomas-bach-sin-oposicion-para-reelegirse-como-presidente-del-comite-olimpico-internacional/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(getUrlContent(u.String()))
}

func getUrlContent(urlToGet string) (string, error) {
	var (
		err     error
		content []byte
		resp    *http.Response
	)

	// GET content of URL
	if resp, err = http.Get(urlToGet); err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check if request was successful
	if resp.StatusCode != 200 {
		return "", err
	}

	// Read the body of the HTTP response
	if content, err = ioutil.ReadAll(resp.Body); err != nil {
		return "", err
	}

	return string(content), err
}
