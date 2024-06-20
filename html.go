package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var raw = `
<!DOCTYPE html>
<html>
  <body>
    <h1>My First Heading</h1>
      <p>My first paragraph.</p>
      <p>HTML <a href="https://www.w3schools.com/html/html_images.asp">images</a> are defined with the img tag:</p>
      <img src="xxx.jpg" width="104" height="142">
  </body>
</html>`

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
