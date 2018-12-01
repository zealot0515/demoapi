package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["demoapi/controllers:DemoController"] = append(beego.GlobalControllerRouter["demoapi/controllers:DemoController"],
        beego.ControllerComments{
            Method: "DemoApi",
            Router: `/demoapi`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
