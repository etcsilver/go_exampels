package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Nota struct {
	Id          string   `json:"id"`
	Category    string   `json:"category"`
	Deporte     []string `json:"deporte"`
	Type        string   `json:"type"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Slug        string   `json:"slug"`
	Link        string   `json:"link"`
	Image       string   `json:"image"`
	Contenido   string   `json:"contenido"`
	Pais        []string `json:"pais"`
	Author      string   `json:"author"`
	Date        string   `json:"date"`
	Modified    string   `json:"modified"`
}

func main() {

	u, err := url.Parse("https://dev-admin-olimpicos.marcaclaro.com/wp-json/wp/v2/entradas/673")
	if err != nil {
		log.Fatal(err)
	}

	text, err := getUrlContent(u.String())
	textBytes := []byte(text)

	nota := Nota{}
	json.Unmarshal(textBytes, &nota)
	fmt.Println(nota.Title)
	fmt.Println(nota.Date)
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
