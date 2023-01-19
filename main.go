package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	requestURL := "https://www.theproductfolks.com/100-product-managers"
	makeRequestAndParseResponse(requestURL)
}

func makeRequestAndParseResponse(requestURL string) {
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		fmt.Printf("clientl: error parsing http response: %s\n", err)
		os.Exit(1)
	}

	var data []string

	doTraverse(doc, &data, "div")
	fmt.Println(data)
}

func doTraverse(doc *html.Node, data *[]string, tag string) {

	var traverse func(n *html.Node, tag string) *html.Node

	traverse = func(n *html.Node, tag string) *html.Node {

		for c := n.FirstChild; c != nil; c = c.NextSibling {

			if c.Type == html.TextNode && c.Parent.Data == tag {

				*data = append(*data, c.Data)
			}

			res := traverse(c, tag)

			if res != nil {

				return res
			}
		}

		return nil
	}

	traverse(doc, tag)
}
