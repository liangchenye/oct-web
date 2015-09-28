package web

import (
	"../middleware"
	"../router"
	"github.com/Unknwon/macaron"
)

func SetOCTMacaron(m *macaron.Macaron) {

	middleware.SetMiddlewares(m)

	//Setting Router
	router.SetRouters(m)
}
