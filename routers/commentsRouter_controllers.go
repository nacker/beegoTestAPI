package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["beegoTestAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["beegoTestAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoTestAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["beegoTestAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoTestAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["beegoTestAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoTestAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["beegoTestAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetById",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beegoTestAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["beegoTestAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
