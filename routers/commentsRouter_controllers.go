package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["metrics-collector/controllers:VideoReproductionController"] = append(beego.GlobalControllerRouter["metrics-collector/controllers:VideoReproductionController"],
		beego.ControllerComments{
			Method:           "PostReproduction",
			Router:           "/video-reproduction",
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
			Filters: nil,
			Params:  nil})

}
