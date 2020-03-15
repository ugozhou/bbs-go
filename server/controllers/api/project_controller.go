package api

import (
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"

	"bbs-go/controllers/render"
	"bbs-go/services"
)

type ProjectController struct {
	Ctx iris.Context
}

func (c *ProjectController) GetBy(projectId int) *simple.JsonResult {
	project := services.ProjectService.Get(projectId)
	if project == nil {
		return simple.JsonErrorMsg("项目不存在")
	}
	return simple.JsonData(render.BuildProject(project))
}

func (c *ProjectController) GetProjects() *simple.JsonResult {
	page := simple.FormValueIntDefault(c.Ctx, "page", 1)

	list, paging := services.ProjectService.FindPageByParams(simple.NewQueryParams(c.Ctx).
		Page(page, 20).Desc("id"))

	return simple.JsonData(&simple.PageResult{Results: list, Page: paging})
}
