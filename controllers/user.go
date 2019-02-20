package controllers

import (
	"beegoTestAPI/models"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"fmt"
)

type UserController struct {
	beego.Controller
}

// @Title 获得所有用户
// @Description 返回所有的用户数据
// @Success 200 {object} models.User
// @router / [get]
func (this *UserController) GetAll() {
	ss := models.GetAllUsers()
	this.Data["json"] = ss
	this.ServeJSON()
}

// @Title 获得一个用户
// @Description 返回某用户数据
// @Param      id            path   int    true          "The key for staticblock"
// @Success 200 {object} models.User
// @router /:id [get]
func (this *UserController) GetById() {
	id, _ := this.GetInt(":id")
	s := models.GetUserById(id)
	this.Data["json"] = s
	this.ServeJSON()
}

// @Title 创建用户
// @Description 创建用户的描述
// @Param      body          body   models.User true          "body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
//func (this *UserController) Post() {
//	user := models.User{}
//	user.Name = this.GetString("name")
//	user.Password = this.GetString("password")
//	user.Avatar = this.GetString("avatar")
//
//	u := models.AddUser(&user)
//	this.Data["json"] = u
//	this.ServeJSON()
//}
func (this *UserController) Post() {
	user := &models.User{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, user)
	if err != nil {
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	u := models.AddUser(user)
	this.Data["json"] = u
	this.ServeJSON()
}

// @Title 修改用户
// @Description 修改用户的内容
// @Param      body          body   models.User true          "body for user content"
// @Success 200 {int} models.User
// @Failure 403 body is empty
// @router / [put]
func (this *UserController) Update() {
	id, _ := this.GetInt(":id")

	fmt.Printf("-------uid: %d", id)

	var user models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	u := models.UpdateUser(id, &user)
	this.Data["json"] = u

	this.ServeJSON()
}

// @Title 删除一个用户
// @Description 删除某用户数据
// @Param      id            path   int    true          "The key for staticblock"
// @Success 200 {object} models.User
// @router /:id [delete]
func (this *UserController) Delete() {
	id, _ := this.GetInt(":id")
	models.DeleteUser(id)
	this.Data["json"] = true
	this.ServeJSON()
}
