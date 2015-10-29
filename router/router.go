package router

import (
	"../handler"
	"github.com/Unknwon/macaron"
)

func SetRouters(m *macaron.Macaron) {
	m.Get("/", handler.IndexHandler)
	m.Get("/news", handler.NewsHandler)
	m.Post("/githubhook", handler.GithubHookHandler)
}
