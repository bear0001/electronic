package controllers

import (
	"DataCertProject/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

//直接访问login.html页面的请求
func (l *LoginController) Get() {
	//	设置login.html为模板文件
	l.TplName = "login.html"
}

//用户登录接口
func (l *LoginController) Post() {
	var user models.User
	err := l.ParseForm(&user)
		if err != nil {
		//l.tplName="error.html"
		l.Ctx.WriteString("抱歉，用户信息解析失败")
		return
}
//	查询数据库的用户信息
	u, err := user.QueryUser()
	if err != nil {
		l.Ctx.WriteString("抱歉，用户登录失败，请重试！")
		fmt.Println(err.Error())
	}
	//  登录成功,跳转项目核心功能页面（home.html）
	l.Data["Phone"] = u.Phone
	l.TplName = "home.html"

}
