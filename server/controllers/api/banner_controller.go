package api

import (
	"bbs-go/services"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
	"strconv"
)

type BannerController struct {
	Ctx iris.Context
}

func (c *BannerController) GetBy(id int64) *simple.JsonResult {
	t := services.BannerService.Get(id)
	if t == nil {
		return simple.JsonErrorMsg("Not found, id=" + strconv.FormatInt(id, 10))
	}
	return simple.JsonData(t)
}

func (c *BannerController) AnyList() *simple.JsonResult {
	list, paging := services.BannerService.FindPageByParams(simple.NewQueryParams(c.Ctx).PageByReq().Desc("id"))
	return simple.JsonData(&simple.PageResult{Results: list, Page: paging})
}
