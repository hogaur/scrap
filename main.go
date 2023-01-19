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

	parseAndPrintResponse(res)

}

func parseAndPrintResponse(response *http.Response) {
	tokenizedResponse := html.NewTokenizer(response.Body)

	for {
		nextToken := tokenizedResponse.Next()

		switch {
		case nextToken == html.ErrorToken:
			// End of the document, we're done
			fmt.Printf("EOF reached: %v\n", nextToken)
			return
		case nextToken == html.StartTagToken:
			linkToken := tokenizedResponse.Token()

			isAnchor := linkToken.Data == "a"
			if isAnchor {
				fmt.Printf("Next Token found %v\n", linkToken)
				fmt.Println("We found a link!")
			}
			for _, a := range linkToken.Attr {
				if a.Key == "href" {
					fmt.Println("Found href:", a.Val)
					break
				}
			}
		}
	}
}
