package router

import (
	"../handler"
	"github.com/Unknwon/macaron"
)

func SetRouters(m *macaron.Macaron) {
	m.Get("/", handler.IndexHandler)
	//	m.Get("/auth", handler.AuthHandler)
	//	m.Get("/admin/auth", handler.AdminAuthHandler)
	//	m.Get("/dashboard", handler.DashboardHandler)
	//	m.Get("/setting", handler.SettingHandler)
}
