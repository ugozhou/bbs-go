package admin

import (
	"strconv"
	"strings"

	"bbs-go/common"
	"bbs-go/model"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"

	"bbs-go/services"
)

type UserController struct {
	Ctx iris.Context
}

//func (c *UserController) GetSynccount() *simple.JsonResult {
//	go func() {
//		services.UserService.SyncUserCount()
//	}()
//	return simple.JsonSuccess()
//}

func (c *UserController) GetBy(id int64) *simple.JsonResult {
	t := services.UserService.Get(id)
	if t == nil {
		return simple.JsonErrorMsg("Not found, id=" + strconv.FormatInt(id, 10))
	}
	return simple.JsonData(c.buildUserItem(t))
}

func (c *UserController) AnyList() *simple.JsonResult {
	list, paging := services.UserService.FindPageByParams(simple.NewQueryParams(c.Ctx).EqByReq("id").LikeByReq("nickname").EqByReq("username").PageByReq().Desc("id"))
	var itemList []map[string]interface{}
	for _, user := range list {
		itemList = append(itemList, c.buildUserItem(&user))
	}
	return simple.JsonData(&simple.PageResult{Results: itemList, Page: paging})
}

func (c *UserController) PostCreate() *simple.JsonResult {
	phone := simple.FormValue(c.Ctx, "phone")
	inviteCode := simple.FormValue(c.Ctx, "inviteCode")
	password := simple.FormValue(c.Ctx, "password")

	user, err := services.UserService.SignUp(phone, inviteCode, password, password)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(c.buildUserItem(user))
}

func (c *UserController) PostCreate2() *simple.JsonResult {
	phone := simple.FormValue(c.Ctx, "phone")
	inviteCode := simple.FormValue(c.Ctx, "inviteCode")
	password := simple.FormValue(c.Ctx, "password")

	user, err := services.UserService.SignUp(phone, inviteCode, password, password)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}

	token, err := services.UserTokenService.Generate(user.Id)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}

	ret := c.buildUserItem(user)
	ret["token"] = token

	return simple.JsonData(ret)
}

func (c *UserController) PostUpdate() *simple.JsonResult {
	id, err := simple.FormValueInt64(c.Ctx, "id")
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	t := services.UserService.Get(id)
	if t == nil {
		return simple.JsonErrorMsg("entity not found")
	}

	password := simple.FormValue(c.Ctx, "password")
	nickname := simple.FormValue(c.Ctx, "nickname")
	roles := simple.FormValueStringArray(c.Ctx, "roles")
	status := simple.FormValueIntDefault(c.Ctx, "status", -1)

	if len(password) > 0 {
		t.Password = simple.EncodePassword(t.Password)
	}
	if len(nickname) > 0 {
		t.Nickname = nickname
	}

	if status != -1 {
		t.Status = status
	}

	t.Roles = strings.Join(roles, ",")

	err = services.UserService.Update(t)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(t)
}

func (c *UserController) buildUserItem(user *model.User) map[string]interface{} {
	return simple.NewRspBuilder(user).
		Put("roles", common.GetUserRoles(user.Roles)).
		Put("phone", user.Mobile).
		Put("nickName", user.Nickname).
		Build()
}
