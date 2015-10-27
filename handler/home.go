package handler

import (
	"github.com/Unknwon/macaron"
)

func IndexHandler(ctx *macaron.Context) {
	ctx.HTML(200, "index")
}

func NewsHandler(ctx *macaron.Context) {
	ctx.HTML(200, "news")
}
