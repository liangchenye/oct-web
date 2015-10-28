package main

import (
	"../"
	"fmt"
)

func main() {
	tpl := "../template/news.tpl"
	newsDir := "../news"
	content := md2html.NewsGenerator(tpl, newsDir)
	fmt.Println(content)
}
