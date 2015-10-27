package router

import (
	"../handler"
	"github.com/Unknwon/macaron"
)

func SetRouters(m *macaron.Macaron) {
	m.Get("/", handler.IndexHandler)
	m.Get("/news", handler.NewsHandler)
	//	m.Get("/admin/auth", handler.AdminAuthHandler)
	//	m.Get("/dashboard", handler.DashboardHandler)
	//	m.Get("/setting", handler.SettingHandler)
}
