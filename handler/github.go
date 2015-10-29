package handler

import (
	"../md2html"
	"github.com/Unknwon/macaron"
	"os"
	"os/exec"
)

func NewsUpdate() {
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
}

func GithubHookHandler(ctx *macaron.Context) {
	NewsUpdate()
	ctx.Resp.Write([]byte("ok"))
}
