package admin

import (
	"admin-api/models"
	"admin-api/models/redis"
	"strconv"
	"time"

	"github.com/ch-123/helper"
)

type Admin struct {
	Base
}

// 登录
// @router /admin/login [get]
func (this *Admin) Login() {
	param := this.checkParam(arrs2{
		{`account`, `请填写账号`},
		{`password`, `请填写密码`},
	})
	uid, msg := models.Login(param[`account`], param[`password`])
	if msg != "" {
		this.msg(msg)
	}
	token := helper.RandomString(32)
	value := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(uid)
	err := redis.Hset(`admin_token`, token, value)
	if err != nil {
		this.msg(`登录失败，服务器出错`)
	}
	this.ajaxdata(token)
}
