// @APIVersion 1.0.0
// @Title Video Service MORA Metrics collector
// @Description Lorem ipsum
// @Contact alessandro.distefano@phd.unict.it
package routers

import (
	"github.com/astaxie/beego"
	"metrics-collector/controllers"
)

func init() {
	b := beego.NewNamespace("/v1",
		beego.NSInclude(
			&controllers.VideoReproductionController{},
		),
	)
	beego.AddNamespace(b)
}
