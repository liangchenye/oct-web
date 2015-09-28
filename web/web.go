package web

import (
	"../middleware"
	"../router"
	"github.com/Unknwon/macaron"
)

func SetWharfMacaron(m *macaron.Macaron) {

	middleware.SetMiddlewares(m)

	//Setting Router
	router.SetRouters(m)
}
