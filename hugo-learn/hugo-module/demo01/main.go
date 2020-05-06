package main

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/gomarkdown/markdown/html"
)

func main()  {
	fmt.Println("start!")
	md := []byte("## markdown document")
	doc :=  parser.New().Parse(md)
	fmt.Println(doc)
}

func f1()  {
	md := []byte("## markdown document")
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	render := html.NewRenderer(opts)

	output := markdown.ToHTML(md,nil, render)
	fmt.Println("out:", string(output))
}