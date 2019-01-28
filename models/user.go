package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/astaxie/beego"
)

type User struct {
	Id       int    `json:"user_id"`                  //用户编号
	Name     string `orm:"size(32)"  json:"name"`     //用户昵称
	Password string `orm:"size(128)" json:"password"` //用户密码
	Avatar   string `orm:"size(256)" json:"avatar"`   //用户头像路径
}

func GetAllUsers() []*User {
	o := orm.NewOrm()
	o.Using("default")
	var user []*User
	q := o.QueryTable("user")
	q.All(&user)
	return user

}
func GetUserById(id int) User {
	u := User{Id: id}
	o := orm.NewOrm()
	o.Using("default")
	err := o.Read(&u)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}
	return u
}

func AddUser(u *User) int {
	o := orm.NewOrm()
	o.Using("default")
	id, err := o.Insert(u)
	if err != nil {
		fmt.Println("inert error = ", err)
	}
	beego.Info("reg insert succ id = ", id)
	return u.Id
}

//func UpdateUser(u *User) {
//	o := orm.NewOrm()
//	o.Using("default")
//	o.Update(u)
//}

func UpdateUser(uid int, uu *User) User {

	//fmt.Printf("-------uid: %d", uid)

	//u := GetUserById(uid)
	u := User{Id: uid}
	o := orm.NewOrm()
	o.Using("default")

	fmt.Printf("-------name: %s, ------u: %s", uu.Name, u)

	if uu.Name != "" {
		u.Name = uu.Name
	}
	if uu.Password != "" {
		u.Password = uu.Password
	}
	if uu.Avatar != "" {
		u.Avatar = uu.Avatar
	}

	o.Update(&u)
	return u
}

func DeleteUser(id int) {
	o := orm.NewOrm()
	o.Using("default")
	o.Delete(&User{Id: id})
}
