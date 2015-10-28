package handler

import (
	"../md2html"
	"github.com/Unknwon/macaron"
	"os"
	"os/exec"
)

func IndexHandler(ctx *macaron.Context) {
	ctx.HTML(200, "index")
}

func NewsHandler(ctx *macaron.Context) {
	ctx.HTML(200, "news")
}

func NewsUpdateHandler(ctx *macaron.Context) {
	tpl := "md2html/template/news.tpl"
	newsDir := "md2html/news"

	c := exec.Command("/bin/sh", "-c", "git pull")
	c.Run()

	newsHtml := md2html.NewsGenerator(tpl, newsDir)
	if len(newsHtml) > 0 {
		f, err := os.Create("views/news.html")
		if err == nil {
			defer f.Close()
			_, err := f.WriteString(newsHtml)
			if err == nil {
				f.Sync()
			}
		}
	}
	ctx.HTML(200, "news")
}
