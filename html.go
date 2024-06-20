package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)


func getHtml() []byte{
	resp,err := http.Get("https://www.amazon.in/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(-1)
	}

	defer resp.Body.Close()

	body , err := io.ReadAll(resp.Body);
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(-1)
	}
	return body
	
}

var raw = string(getHtml())

func visit(n *html.Node, pwords , ppics *int) {
	if n.Type == html.TextNode {
		*pwords += len(strings.Fields(n.Data))
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		*ppics++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, pwords, ppics)
	}
}

func countWordsAndImages(n *html.Node) (int,int) {

	words := 0
	pics := 0

	visit(n, &words, &pics)

	return words,pics
}

func main() {

	doc, err := html.Parse(bytes.NewReader([]byte(raw)))

	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(-1)
	}

	words,pics := countWordsAndImages(doc)
	fmt.Printf("words: %d\n", words)
	fmt.Printf("pics: %d\n", pics)

}
