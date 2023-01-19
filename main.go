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
	pmCardPattern := "tpf-100-pms-card-details-wrap"
	currentRolePattern := "tpf-100-pms-card-role-text"
	previousRolePattern := "tpf-100-pms-card-old-role"
	doc := makeRequestAndParseResponse(requestURL)
	allPms := findAllPmsFromNode(doc, pmCardPattern, currentRolePattern, previousRolePattern)
	fmt.Println(len(allPms))
}

func makeRequestAndParseResponse(requestURL string) *html.Node {
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

	return doc
}

func findAllPmsFromNode(doc *html.Node, pmCardPattern, currentRolePattern, previousRolePattern string) []pm.Pm {
	pmNodes := doTraverse(doc, "div", pmCardPattern)
	var allPms []pm.Pm
	for _, pmNode := range pmNodes {
		name := pmNode.FirstChild.FirstChild.Data
		currentRole := findRoleFromNode(pmNode, currentRolePattern)
		previousRole := findRoleFromNode(pmNode, previousRolePattern)

		p := pm.BuildPmInfo(name, currentRole, previousRole)
		allPms = p.AddPmToList(allPms)
	}
	return allPms
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

func findRoleFromNode(pmNode *html.Node, pattern string) string {
	roleNodes := doTraverse(pmNode, "div", pattern)
	role := ""
	for _, roleNode := range roleNodes {
		if roleNode.FirstChild != nil {
			role = fmt.Sprintf("%v%v", role, roleNode.FirstChild.Data)
		}
	}
	return role
}
