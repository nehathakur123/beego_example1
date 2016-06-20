package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["beego_example1/controllers:ObjectController"] = append(beego.GlobalControllerRouter["beego_example1/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["beego_example1/controllers:ObjectController"] = append(beego.GlobalControllerRouter["beego_example1/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beego_example1/controllers:ObjectController"] = append(beego.GlobalControllerRouter["beego_example1/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beego_example1/controllers:ObjectController"] = append(beego.GlobalControllerRouter["beego_example1/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["beego_example1/controllers:ObjectController"] = append(beego.GlobalControllerRouter["beego_example1/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["beego_example1/controllers:StudentController"] = append(beego.GlobalControllerRouter["beego_example1/controllers:StudentController"],
		beego.ControllerComments{
			"GetAllStudents",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beego_example1/controllers:StudentController"] = append(beego.GlobalControllerRouter["beego_example1/controllers:StudentController"],
		beego.ControllerComments{
			"CreateStudent",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["beego_example1/controllers:StudentController"] = append(beego.GlobalControllerRouter["beego_example1/controllers:StudentController"],
		beego.ControllerComments{
			"GetOneStudent",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beego_example1/controllers:StudentController"] = append(beego.GlobalControllerRouter["beego_example1/controllers:StudentController"],
		beego.ControllerComments{
			"UpdateStudent",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["beego_example1/controllers:StudentController"] = append(beego.GlobalControllerRouter["beego_example1/controllers:StudentController"],
		beego.ControllerComments{
			"DeleteStudent",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["beego_example1/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_example1/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["beego_example1/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_example1/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beego_example1/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_example1/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beego_example1/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_example1/controllers:UserController"],
		beego.ControllerComments{
			"Put",
			`/:uid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["beego_example1/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_example1/controllers:UserController"],
		beego.ControllerComments{
			"Delete",
			`/:uid`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["beego_example1/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_example1/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beego_example1/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_example1/controllers:UserController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

}
