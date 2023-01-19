package main

import (
	"fmt"
	"net/http"
	"os"
	"scrappergo/pm"
	"strings"

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

	pmNodes := doTraverse(doc, "div", "tpf-100-pms-card-details-wrap")
	fmt.Println(pmNodes)
	fmt.Println(len(pmNodes))
	var allPms []pm.Pm
	for _, pmNode := range pmNodes {
		name := pmNode.FirstChild.FirstChild.Data
		fmt.Println(name)

		currentRoleNodes := doTraverse(pmNode, "div", "tpf-100-pms-card-role-text")
		currentRole := ""
		for _, currentRoleNode := range currentRoleNodes {
			if currentRoleNode.FirstChild != nil {
				currentRole = fmt.Sprintf("%v%v", currentRole, currentRoleNode.FirstChild.Data)
			}
		}
		fmt.Println(currentRole)

		previousRoleNodes := doTraverse(pmNode, "div", "tpf-100-pms-card-old-role")
		previousRole := ""
		for _, previousRoleNode := range previousRoleNodes {
			if previousRoleNode.FirstChild != nil {
				previousRole = fmt.Sprintf("%v%v", previousRole, previousRoleNode.FirstChild.Data)
			}
		}
		fmt.Println(previousRole)
		pm := pm.BuildPmInfo(name, currentRole, previousRole)
		allPms = append(allPms, pm)
		fmt.Println(allPms)
	}
}

func doTraverse(doc *html.Node, tag string, pattern string) []*html.Node {
	var pmNodes []*html.Node
	var traverse func(n *html.Node, tag string) *html.Node
	attr := "class"

	traverse = func(n *html.Node, tag string) *html.Node {
		for c := n.FirstChild; c != nil; c = c.NextSibling {

			if c.Data == tag {
				for _, a := range c.Attr {
					if a.Key == attr && strings.Contains(a.Val, pattern) {
						pmNodes = append(pmNodes, c)
					}
				}
			}
			res := traverse(c, tag)

			if res != nil {
				return res
			}
		}

		return nil
	}

	traverse(doc, tag)
	return pmNodes
}
