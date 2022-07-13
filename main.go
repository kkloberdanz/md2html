package main

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"io/ioutil"
	"os"
)

func main() {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	md, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	html := markdown.ToHTML(md, parser, nil)
	fmt.Printf("%s", html)
}
