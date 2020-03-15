package api

import (
	"bbs-go/services"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
	"strconv"
)

type NoticeController struct {
	Ctx iris.Context
}

func (c *NoticeController) GetBy(id int64) *simple.JsonResult {
	t := services.NoticeService.Get(id)
	if t == nil {
		return simple.JsonErrorMsg("Not found, id=" + strconv.FormatInt(id, 10))
	}
	return simple.JsonData(t)
}

func (c *NoticeController) AnyList() *simple.JsonResult {
	list, paging := services.NoticeService.FindPageByParams(simple.NewQueryParams(c.Ctx).PageByReq().Desc("id"))
	return simple.JsonData(&simple.PageResult{Results: list, Page: paging})
}
