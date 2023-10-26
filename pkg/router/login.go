package router

import (
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/open-beagle/awecloud-bmq-server/pkg/conf"
	"github.com/open-beagle/awecloud-bmq-server/pkg/model"
	"github.com/open-beagle/awecloud-bmq-server/pkg/util"
)

func Login(c *gin.Context) {
	res := model.APIResponse{}
	user := &model.User{}
	err := c.BindJSON(user)
	if err != nil {
		res.Success = 0
		res.ErrMsg = `传参错误!`
		c.JSON(200, res)
		return
	}
	if user.Name == `` || user.Password == `` {
		res.Success = 0
		res.ErrMsg = `用户名或密码不能为空!`
		c.JSON(200, res)
		return
	}

	loginUser := conf.Server.LoginUser(user.Name, user.Password)
	if loginUser == nil {
		res.Success = 0
		res.ErrMsg = `用户不存在或密码错误!`
		c.JSON(200, res)
		return
	}

	res.Success = 1
	res.Data = loginUser

	session := sessions.Default(c)
	session.Set("User", loginUser)
	session.Save()

	c.JSON(200, res)
}

func Logout(c *gin.Context) {
	res := model.APIResponse{
		Success: 1,
	}
	session := sessions.Default(c)
	session.Clear()

	c.JSON(200, res)
}

func LoginUser(c *gin.Context) {
	res := model.APIResponse{}

	session := sessions.Default(c)
	user := session.Get("User")

	if user != nil {
		res.Success = 1
		res.Data = user
	} else {
		res.Success = 0
		res.ErrMsg = "用户未登录."
	}

	c.JSON(200, res)
}

func LoginCaptchaGet(c *gin.Context) {
	res := model.APIResponse{}
	widthS := c.Query("width")
	heightS := c.Query("height")
	var width, height int
	if widthS == "" {
		width = conf.API.CaptchaWidth
	} else {
		width, _ = strconv.Atoi(widthS)
	}
	if heightS == "" {
		height = conf.API.CaptchaHeight
	} else {
		height, _ = strconv.Atoi(heightS)
	}

	captcha, err := util.GetCaptcha(width, height)
	if err != nil {
		res.Success = 0
		res.ErrMsg = err.Error()
	}

	res.Success = 1
	res.Data = captcha

	session := sessions.Default(c)
	session.Set("Captcha", captcha)
	session.Save()

	c.JSON(200, res)
}

func LoginCaptchaCheck(c *gin.Context) {
	res := model.APIResponse{}

	captcha := &model.Captcha{}
	err := c.BindJSON(captcha)
	if err != nil {
		res.Success = 0
		res.ErrMsg = `传参错误!`
		c.JSON(200, res)
		return
	}

	session := sessions.Default(c)
	captchaSession, ok := session.Get("Captcha").(model.Captcha)
	if !ok {
		res.Success = 0
		res.ErrMsg = `验证码丢失!`
		c.JSON(200, res)
		return
	}

	if captcha.Id != captchaSession.Id || captcha.Captcha != captchaSession.Captcha {
		res.Success = 0
		res.ErrMsg = `验证码错误!`
		c.JSON(200, res)
		return
	}

	res.Success = 1
	c.JSON(200, res)
}
