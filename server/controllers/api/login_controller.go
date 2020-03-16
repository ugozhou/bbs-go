package api

import (
	"github.com/goburrow/cache"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
	"time"

	"bbs-go/common"
	"bbs-go/common/github"
	"bbs-go/common/qq"
	"bbs-go/controllers/render"
	"bbs-go/model"
	"bbs-go/services"
)

type LoginController struct {
	Ctx iris.Context
}

//type RegistForm struct {
//	phone string	`json:"phone"`
//	inviateCode  string	`json:"inviateCode"`
//	password	string	`json:"password"`
//	rePassword string	`json:"rePassword"`
//	captchald string	`json:"captchald"`
//	ret string	`json:"ret"`
//}
var phonecache = cache.New(cache.WithMaximumSize(1000),
	cache.WithExpireAfterAccess(10*time.Second),
	cache.WithRefreshAfterWrite(90*time.Second))

// 注册
func (c *LoginController) PostSignup() *simple.JsonResult {

	//var requestData RegistForm
	////c.Ctx.JSON(&requestData)
	//error := c.Ctx.ReadJSON(&requestData)
	//if error != nil {
	//	return simple.JsonErrorMsg(error.Error())
	//}
	//println(requestData.phone)
	var (
		//captchaId   = c.Ctx.PostValueTrim("captchaId")
		//captchaCode = c.Ctx.PostValueTrim("captchaCode")
		//email       = c.Ctx.PostValueTrim("email")
		mobile      = c.Ctx.PostValueTrim("phone")
		inviteCode  = c.Ctx.PostValueTrim("inviteCode")
		messageCode = c.Ctx.PostValueTrim("messageCode")
		password    = c.Ctx.PostValueTrim("password")
		rePassword  = c.Ctx.PostValueTrim("rePassword")
	)
	if !common.IsValidateMobile(mobile) {
		return simple.JsonError(common.MobileError)
	}
	verifyCode, exist := phonecache.GetIfPresent(mobile)
	if !exist || messageCode != verifyCode {
		return simple.JsonErrorMsg("验证码错误")
	}
	//if !captcha.VerifyString(captchaId, captchaCode) {
	//	return simple.JsonError(common.CaptchaError)
	//}
	user, err := services.UserService.SignUp(mobile, inviteCode, password, password)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return c.GenerateLoginResult(user, "index")
}

// 用户名密码登录
func (c *LoginController) PostSignin() *simple.JsonResult {
	var (
		//captchaId   = c.Ctx.PostValueTrim("captchaId")
		//captchaCode = c.Ctx.PostValueTrim("captchaCode")
		phone    = c.Ctx.PostValueTrim("phone")
		password = c.Ctx.PostValueTrim("password")
		ref      = c.Ctx.FormValue("ref")
	)
	//if !captcha.VerifyString(captchaId, captchaCode) {
	//	return simple.JsonError(common.CaptchaError)
	//}
	user, err := services.UserService.SignIn(phone, password)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return c.GenerateLoginResult(user, ref)
}

func (c *LoginController) GetVerifyinviatecode() *simple.JsonResult {
	invitecode := c.Ctx.FormValue("invitecode")
	if len(invitecode) == 0 {
		return simple.JsonErrorMsg("参数错误")
	}
	result := services.UserService.IsInviteCodeExist(invitecode)
	if !result {
		return simple.JsonErrorMsg("邀请码不存在")
	}
	return simple.JsonSuccess()
}

// 退出登录
func (c *LoginController) GetSignout() *simple.JsonResult {
	err := services.UserTokenService.Signout(c.Ctx)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonSuccess()
}

// 获取Github登录授权地址
func (c *LoginController) GetGithubAuthorize() *simple.JsonResult {
	ref := c.Ctx.FormValue("ref")
	url := github.AuthCodeURL(map[string]string{"ref": ref})
	return simple.NewEmptyRspBuilder().Put("url", url).JsonResult()
}

// 获取Github回调信息获取
//func (c *LoginController) GetGithubCallback() *simple.JsonResult {
//	code := c.Ctx.FormValue("code")
//	state := c.Ctx.FormValue("state")
//
//	thirdAccount, err := services.ThirdAccountService.GetOrCreateByGithub(code, state)
//	if err != nil {
//		return simple.JsonErrorMsg(err.Error())
//	}
//
//	user, codeErr := services.UserService.SignInByThirdAccount(thirdAccount)
//	if codeErr != nil {
//		return simple.JsonError(codeErr)
//	} else {
//		return c.GenerateLoginResult(user, "")
//	}
//}

// 获取QQ登录授权地址
func (c *LoginController) GetQqAuthorize() *simple.JsonResult {
	ref := c.Ctx.FormValue("ref")
	url := qq.AuthorizeUrl(map[string]string{"ref": ref})
	return simple.NewEmptyRspBuilder().Put("url", url).JsonResult()
}

// 获取QQ回调信息获取
//func (c *LoginController) GetQqCallback() *simple.JsonResult {
//	code := c.Ctx.FormValue("code")
//	state := c.Ctx.FormValue("state")
//
//	thirdAccount, err := services.ThirdAccountService.GetOrCreateByQQ(code, state)
//	if err != nil {
//		return simple.JsonErrorMsg(err.Error())
//	}
//
//	user, codeErr := services.UserService.SignInByThirdAccount(thirdAccount)
//	if codeErr != nil {
//		return simple.JsonError(codeErr)
//	} else {
//		return c.GenerateLoginResult(user, "")
//	}
//}

// user: login user, ref: 登录来源地址，需要控制登录成功之后跳转到该地址
func (c *LoginController) GenerateLoginResult(user *model.User, ref string) *simple.JsonResult {
	token, err := services.UserTokenService.Generate(user.Id)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.NewEmptyRspBuilder().
		Put("token", token).
		Put("user", render.BuildUser(user)).
		Put("ref", ref).JsonResult()
}

func (c *LoginController) GetSendsms() *simple.JsonResult {

	//captcha.RandomDigits(4)
	phone := c.Ctx.FormValue("phone")
	verifycode := common.GetRandomString(4)
	phonecache.Put(phone, verifycode)
	return simple.NewEmptyRspBuilder().
		Put("testverifycode", verifycode).
		Put("phone", phone).JsonResult()
}
