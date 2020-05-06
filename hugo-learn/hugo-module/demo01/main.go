package main

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Extensions int

const (
	NoExtensions           Extensions = 0
	NoIntraEmphasis        Extensions = 1 << iota // Ignore emphasis markers inside words
	Tables                                        // Parse tables
	FencedCode                                    // Parse fenced code blocks
	Autolink                                      // Detect embedded URLs that are not explicitly marked
	Strikethrough                                 // Strikethrough text using ~~test~~
	LaxHTMLBlocks                                 // Loosen up HTML block parsing rules
	SpaceHeadings                                 // Be strict about prefix heading rules
	HardLineBreak                                 // Translate newlines into line breaks
	NonBlockingSpace                              // Translate backspace spaces into line non-blocking spaces
	TabSizeEight                                  // Expand tabs to eight spaces instead of four
	Footnotes                                     // Pandoc-style footnotes
	NoEmptyLineBeforeBlock                        // No need to insert an empty line to start a (code, quote, ordered list, unordered list) block
	HeadingIDs                                    // specify heading IDs  with {#id}
	Titleblock                                    // Titleblock ala pandoc
	AutoHeadingIDs                                // Create the heading ID from the text
	BackslashLineBreak                            // Translate trailing backslashes into line breaks
	DefinitionLists                               // Parse definition lists
	MathJax                                       // Parse MathJax
	OrderedListStart                              // Keep track of the first number used when starting an ordered list.
	Attributes                                    // Block Attributes
	SuperSubscript                                // Super- and subscript support: 2^10^, H~2~O.
	EmptyLinesBreakList                           // 2 empty lines break out of list
	Includes                                      // Support including other files.
	Mmark                                         // Support Mmark syntax, see https://mmark.nl/syntax

	CommonExtensions Extensions = NoIntraEmphasis | Tables | FencedCode |
		Autolink | Strikethrough | SpaceHeadings | HeadingIDs |
		BackslashLineBreak | DefinitionLists | MathJax
)

func main() {
	fmt.Println("start!")
	fmt.Printf("| NoExtensions | %d | %032b |\n", NoExtensions, NoExtensions)
	fmt.Printf("| NoIntraEmphasis | %d | %032b |\n", NoIntraEmphasis, NoIntraEmphasis)
	fmt.Printf("| Tables | %d | %032b |\n", Tables, Tables)
	fmt.Printf("| FencedCode | %d | %032b |\n", FencedCode, FencedCode)
	fmt.Printf("| Autolink | %d | %032b |\n", Autolink, Autolink)
	fmt.Printf("| Strikethrough | %d | %032b |\n", Strikethrough, Strikethrough)
	fmt.Printf("| LaxHTMLBlocks | %d | %032b |\n", LaxHTMLBlocks, LaxHTMLBlocks)
	fmt.Printf("| SpaceHeadings | %d | %032b |\n", SpaceHeadings, SpaceHeadings)
	fmt.Printf("| HardLineBreak | %d | %032b |\n", HardLineBreak, HardLineBreak)
	fmt.Printf("| NonBlockingSpace | %d | %032b |\n", NonBlockingSpace, NonBlockingSpace)
	fmt.Printf("| TabSizeEight | %d | %032b |\n", TabSizeEight, TabSizeEight)
	fmt.Printf("| Footnotes | %d | %032b |\n", Footnotes, Footnotes)
	fmt.Printf("| NoEmptyLineBeforeBlock | %d | %032b |\n", NoEmptyLineBeforeBlock, NoEmptyLineBeforeBlock)
	fmt.Printf("| HeadingIDs | %d | %032b |\n", HeadingIDs, HeadingIDs)
	fmt.Printf("| Titleblock | %d | %032b |\n", Titleblock, Titleblock)
	fmt.Printf("| AutoHeadingIDs | %d | %032b |\n", AutoHeadingIDs, AutoHeadingIDs)
	fmt.Printf("| BackslashLineBreak | %d | %032b |\n", BackslashLineBreak, BackslashLineBreak)
	fmt.Printf("| DefinitionLists | %d | %032b |\n", DefinitionLists, DefinitionLists)
	fmt.Printf("| MathJax | %d | %032b |\n", MathJax, MathJax)
	fmt.Printf("| OrderedListStart | %d | %032b |\n", OrderedListStart, OrderedListStart)
	fmt.Printf("| Attributes | %d | %032b |\n", Attributes, Attributes)
	fmt.Printf("| SuperSubscript | %d | %032b |\n", SuperSubscript, SuperSubscript)
	fmt.Printf("| EmptyLinesBreakList | %d | %032b |\n", EmptyLinesBreakList, EmptyLinesBreakList)
	fmt.Printf("| Includes | %d | %032b |\n", Includes, Includes)
	fmt.Printf("| Mmark | %d | %032b |\n", Mmark, Mmark)
	fmt.Printf("| CommonExtensions | %d | %032b |\n", CommonExtensions, CommonExtensions)
}

func f1() {
	md := []byte("## markdown document")
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	render := html.NewRenderer(opts)

	output := markdown.ToHTML(md, nil, render)
	fmt.Println("out:", string(output))
}

func f2() {
	md := []byte("## markdown document")
	doc := parser.New().Parse(md)
	fmt.Println(doc)
}
