package controllers

import (
	"blog/bootstrap/driver"
	"blog/modules"
	"blog/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	_ "github.com/gin-contrib/sessions/cookie"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 表单提交数据
type formUser struct {
	Name string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email string `form:"email" json:"email" binding:"required"`
}

// render 注册页面
func Register(c *gin.Context) {
	auth := Auth{}.GetAuth(c)
	if auth.Id > 0 {
		utils.Redirect(c, "/")
		return
	}
	data := struct {
		Title string
		Auth
	}{
		"注册",
		auth,
	}

	c.HTML(http.StatusOK, "join", data)
}

// 注册用户
func DoRegister(c *gin.Context) {
	var regData formUser
	// 数据验证
	if err := c.ShouldBind(&regData); err != nil {
		response := utils.Response{
			Status: 403,
			Data:   nil,
			Msg:    err.Error(),
		}

		c.JSON(http.StatusBadRequest, response.FailedResponse())
		return
	}
	var user modules.User
	// 保存数据到数据库
	driver.Db.Where("email = ?", regData.Email).First(&user)
	if user.ID != 0 {
		response := utils.Response{
			Status: 1001,
			Data:   nil,
			Msg:    "邮箱已经存在",
		}
		c.JSON(http.StatusOK, response.FailedResponse())
		return
	}

	salt := utils.Salt()
	user = modules.User{
		Name:       regData.Name,
		Password:   utils.CryptUserPassword(regData.Password, salt),
		Salt:		salt,
		Email:      regData.Email,
		Profession: "",
		Avatar:     "",
	}
	// 保存数据
	err := driver.Db.Create(&user).Error
	if err != nil {
		response := utils.Response{
			Status: 500,
			Data:   nil,
			Msg:    err.Error(),
		}
		c.JSON(http.StatusOK, response.FailedResponse())
		return
	}
	// 将用户数据写入session
	auth := &Auth{
		Id:int(user.ID),
		Name:user.Name,
		Avatar:user.Avatar,
		Profession: user.Profession,
	}

	sess := sessions.Default(c)
	sess.Set("auth", auth)
	err  = sess.Save()

	fmt.Println(sess.Get("auth"))
	if err != nil {
		c.JSON(http.StatusOK, (&utils.Response{
			Status: 500,
			Data:   nil,
			Msg:    err.Error(),
		}).FailedResponse())
		return
	}
	// 相应数据
	response := utils.Response{
		Data:   regData,
	}
	c.JSON(http.StatusOK, response.SuccessResponse())
}

// get sign in
func Login(c *gin.Context)  {
	auth := (&Auth{}).GetAuth(c)
	if auth.Id != 0 {
		utils.Redirect(c, "/")
	}
	data := struct {
		Auth
	}{auth}
	c.HTML(http.StatusOK, "login", data)
}
// post sign in
func DoLogin(c *gin.Context)  {
	var logData struct{
		Email string `form:"email" json:"email" binding:"required" `
		Password string `form:"password" json:"password" binding:"required"`
	}
	// 验证数据
	if err := c.ShouldBind(&logData);err != nil {
		response := utils.Response{
			Status: 403,
			Data:   nil,
			Msg:    err.Error(),
		}
		c.JSON(http.StatusOK, response.FailedResponse())
		return
	}
	// 查找用户， 密码是否匹配
	user := modules.User{}
	err := driver.Db.Where("email = ?", logData.Email).First(&user).Error
	if err != nil {
		response := utils.Response{
			Status: 404,
			Data:   nil,
			Msg:    "登录失败， 请确认邮箱是否正确!",
		}
		c.JSON(http.StatusOK, response.FailedResponse())
		return
	}
	// 找到了后验证密码是否匹配
	ok := utils.VerifyUserPassword(&user, logData.Password)
	if !ok {
		c.JSON(http.StatusOK, (&utils.Response{
			Status: 403,
			Data:   nil,
			Msg:    "邮箱或者密码错误.",
		}).FailedResponse())
		return
	}

	// 验证成功后， 写将用户的信息写入session中
	auth := &Auth{
		Id:int(user.ID),
		Name:user.Name,
		Avatar:user.Avatar,
		Profession: user.Profession,
	}

	sess := sessions.Default(c)
	sess.Set("auth", auth)
	err = sess.Save()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, (&utils.Response{
		Status: 0,
		Data:   nil,
		Msg:    "",
	}).SuccessResponse())
	return
}


// logout
func Logout(c *gin.Context)  {
	sess := sessions.Default(c)
	sess.Clear()
	err := sess.Save()
	if err != nil {
		panic(err)
	}
	c.Redirect(http.StatusFound, "/")
	return
}
