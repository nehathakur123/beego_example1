package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"strconv"
	"time"
)

type Student struct {
	Id        int
	Name      string `valid:"Required"`
	Status    int
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Student))
}

func (u *Student) TableName() string {
	return "students"
}

func GetAllStudents() []*Student {
	var students []*Student
	o := orm.NewOrm()
	qs := o.QueryTable("students")
	num, err := qs.All(&students)
	if num > 0 {
		fmt.Println(err)
	}

	return students
}

func CreateStudent(student Student) (s Student, err error) {
	o := orm.NewOrm()
	valid := validation.Validation{}
	valid.Required(student.Name, "name")
	if valid.HasErrors() {
		// validation does not pass
		// print invalid message
		return s, errors.New("Invalid Input")
	}
	id, err := o.Insert(&student)
	if err == nil {
		fmt.Println(id)
	}
	return student, nil
}

func GetOneStudent(id string) (Student, error) {
	uid, _ := strconv.Atoi(id)
	student := Student{Id: uid}
	o := orm.NewOrm()
	err := o.Read(&student)
	if err == orm.ErrNoRows {
		fmt.Println(errors.New("not"))
		return student, errors.New("404")
	} else {
		return student, nil
	}
}

func UpdateStudent(student Student) (s Student, err error) {
	fmt.Println(student.Id)
	o := orm.NewOrm()
	valid := validation.Validation{}
	valid.Required(student.Name, "name")
	if valid.HasErrors() {
		// validation does not pass
		// print invalid message
		return s, errors.New("Invalid Input")
	}
	id, err := o.Update(&student)
	if err == nil {
		fmt.Println(id)
	} else {
		fmt.Println(err)
	}
	return student, err
}

func DeleteStudent(id int) (bool, error) {
	o := orm.NewOrm()
	if num, err := o.Delete(&Student{Id: id}); err == nil {
		fmt.Println(num)
		return true, nil
	}
	return false, errors.New("Error")
}
