package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/ch-123/helper"
)

type Admin struct {
	Id                int
	Account, Password string
	Status, Rstatus   int8
}

//登录
func Login(account string, password string) (int, string) {

	user := Admin{}

	err := orm.NewOrm().Raw(`SELECT a.*,r.name,r.status rstatus from db_admin a LEFT JOIN db_admin_role r on a.role_id = r.id where a.account = ? and a.delete_time = 0`, account).QueryRow(&user)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return 0, "用户名不存在"
	}
	if user.Password != helper.Md5(helper.Md5(password)) {
		return 0, "密码不正确"
	}
	if user.Status == 2 || user.Rstatus == 2 {
		return 0, `账号已停用`
	}
	return user.Id, ""

}
