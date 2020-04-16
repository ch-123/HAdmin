package admin

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type arrs2 [][]string

type JsonData struct {
	Code int8        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Base struct {
	beego.Controller
}

//返回错误信息
func (this *Base) msg(s string) {
	this.Ctx.WriteString(`{"code":0,"msg":"` + s + `","data":{}}`)
	this.StopRun()
}

//返回json数据
func (this *Base) ajaxdata(data interface{}) {
	s, err := json.Marshal(JsonData{1, "", data})
	if err != nil {
		this.msg(`json error`)
		this.StopRun()
		return
	}
	this.Ctx.WriteString(string(s))
	this.StopRun()
}

//返回json字符串
func (this *Admin) Json(d JsonData) {
	s, err := json.Marshal(d)
	if err == nil {
		this.Ctx.WriteString(`{"code":0,"msg":"error","data":{}}`)
		this.StopRun()
		return
	}
	fmt.Printf("%+v\n", string(s))
	this.Ctx.WriteString("asdadas")
	this.StopRun()
}

//校验参数
func (this *Base) checkParam(arr arrs2) map[string]string {
	ret := make(map[string]string)
	for _, v := range arr {
		val := this.GetString(v[0])
		if val == "" && v[1] != "" {
			this.msg(v[1])
		}
		ret[v[0]] = val
	}
	return ret
	// this.Ctx.Input.Param(":key")
}
