package controllers

import (
	"beego_example1/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type StudentController struct {
	beego.Controller
}

// @router / [get]
func (o *StudentController) GetAllStudents() {
	obs := models.GetAllStudents()
	o.Data["json"] = obs
	o.ServeJSON()

}

// @router / [post]
func (o *StudentController) CreateStudent() {
	var student models.Student
	//o.ParseForm(&student)
	json.Unmarshal(o.Ctx.Input.RequestBody, &student)
	fmt.Println(student)
	result, err := models.CreateStudent(student)
	if err != nil {
		o.Abort("500")
	} else {
		o.Data["json"] = result
	}
	o.ServeJSON()
}

// @router /:id [get]
func (o *StudentController) GetOneStudent() {
	uid := o.GetString(":id")
	if uid == "" {
		o.Abort("403")
	}

	student, err := models.GetOneStudent(uid)
	if err != nil {
		o.Abort("500")
	} else {
		o.Data["json"] = student
	}
	o.ServeJSON()
}

// @router /:id [put]
func (o *StudentController) UpdateStudent() {
	uid := o.GetString(":id")
	if uid == "" {
		o.Abort("403")
	}
	id, _ := strconv.Atoi(uid)
	student := models.Student{Id: id}

	//o.ParseForm(&student)
	json.Unmarshal(o.Ctx.Input.RequestBody, &student)
	result, err := models.UpdateStudent(student)
	if err != nil {
		o.Abort("500")
	} else {
		o.Data["json"] = result
	}

	o.ServeJSON()
}

// @router /:id [delete]
func (o *StudentController) DeleteStudent() {
	uid := o.GetString(":id")
	if uid == "" {
		o.Abort("403")
	}

	id, error := strconv.Atoi(uid)
	if error != nil {
		o.Abort("500")
	}
	response, error := models.DeleteStudent(id)
	if error == nil {
		o.Data["json"] = response
	} else {
		o.Abort("500")
	}

	o.ServeJSON()
}
