package render

import (
	"bbs-go/common/urls"
	"github.com/PuerkitoBio/goquery"
	"html/template"
	"strconv"
	"strings"

	"bbs-go/common/avatar"
	"bbs-go/model"
	"bbs-go/services"
	"bbs-go/services/cache"
	"github.com/mlogclub/simple"
)

func BuildUserDefaultIfNull(id int64) *model.UserInfo {
	user := cache.UserCache.Get(id)
	if user == nil {
		user = &model.User{}
		user.Id = id
		user.Nickname = strconv.FormatInt(id, 10)
		user.Avatar = avatar.DefaultAvatar
		user.CreateTime = simple.NowTimestamp()
	}
	return BuildUser(user)
}

func BuildUserById(id int64) *model.UserInfo {
	user := cache.UserCache.Get(id)
	return BuildUser(user)
}

func BuildUser(user *model.User) *model.UserInfo {
	if user == nil {
		return nil
	}
	a := user.Avatar
	if len(a) == 0 {
		a = avatar.DefaultAvatar
	}
	roles := strings.Split(user.Roles, ",")
	ret := &model.UserInfo{
		Id:          user.Id,
		Nickname:    user.Nickname,
		Avatar:      a,
		Phone:       user.Mobile,
		Type:        user.Type,
		Roles:       roles,
		PasswordSet: len(user.Paypassword) > 0,
		Status:      user.Status,
		CreateTime:  user.CreateTime,
	}
	return ret
}

func BuildUsers(users []model.User) []model.UserInfo {
	if len(users) == 0 {
		return nil
	}
	var responses []model.UserInfo
	for _, user := range users {
		item := BuildUser(&user)
		if item != nil {
			responses = append(responses, *item)
		}
	}
	return responses
}

func BuildArticle(article *model.Article) *model.ArticleResponse {
	if article == nil {
		return nil
	}

	rsp := &model.ArticleResponse{}
	rsp.ArticleId = article.Id
	rsp.Title = article.Title
	rsp.Summary = article.Summary
	rsp.Share = article.Share
	rsp.SourceUrl = article.SourceUrl
	rsp.ViewCount = article.ViewCount
	rsp.CreateTime = article.CreateTime

	rsp.User = BuildUserDefaultIfNull(article.UserId)

	tagIds := cache.ArticleTagCache.Get(article.Id)
	tags := cache.TagCache.GetList(tagIds)
	rsp.Tags = BuildTags(tags)

	if article.ContentType == model.ContentTypeMarkdown {
		mr := simple.NewMd(simple.MdWithTOC()).Run(article.Content)
		rsp.Content = template.HTML(BuildHtmlContent(mr.ContentHtml))
		rsp.Toc = template.HTML(mr.TocHtml)
		if len(rsp.Summary) == 0 {
			rsp.Summary = mr.SummaryText
		}
	} else {
		rsp.Content = template.HTML(BuildHtmlContent(article.Content))
		if len(rsp.Summary) == 0 {
			rsp.Summary = simple.GetSummary(article.Content, 256)
		}
	}

	return rsp
}

func BuildArticles(articles []model.Article) []model.ArticleResponse {
	if articles == nil || len(articles) == 0 {
		return nil
	}
	var responses []model.ArticleResponse
	for _, article := range articles {
		responses = append(responses, *BuildArticle(&article))
	}
	return responses
}

func BuildSimpleArticle(article *model.Article) *model.ArticleSimpleResponse {
	if article == nil {
		return nil
	}

	rsp := &model.ArticleSimpleResponse{}
	rsp.ArticleId = article.Id
	rsp.Title = article.Title
	rsp.Summary = article.Summary
	rsp.Share = article.Share
	rsp.SourceUrl = article.SourceUrl
	rsp.ViewCount = article.ViewCount
	rsp.CreateTime = article.CreateTime

	rsp.User = BuildUserDefaultIfNull(article.UserId)

	tagIds := cache.ArticleTagCache.Get(article.Id)
	tags := cache.TagCache.GetList(tagIds)
	rsp.Tags = BuildTags(tags)

	if article.ContentType == model.ContentTypeMarkdown {
		if len(rsp.Summary) == 0 {
			mr := simple.NewMd(simple.MdWithTOC()).Run(article.Content)
			rsp.Summary = mr.SummaryText
		}
	} else {
		if len(rsp.Summary) == 0 {
			rsp.Summary = simple.GetSummary(simple.GetHtmlText(article.Content), 256)
		}
	}

	return rsp
}

func BuildSimpleArticles(articles []model.Article) []model.ArticleSimpleResponse {
	if articles == nil || len(articles) == 0 {
		return nil
	}
	var responses []model.ArticleSimpleResponse
	for _, article := range articles {
		responses = append(responses, *BuildSimpleArticle(&article))
	}
	return responses
}

//func BuildNode(node *model.TopicNode) *model.NodeResponse {
//	if node == nil {
//		return nil
//	}
//	return &model.NodeResponse{
//		NodeId:      node.Id,
//		Name:        node.Name,
//		Description: node.Description,
//	}
//}
//
//func BuildNodes(nodes []model.TopicNode) []model.NodeResponse {
//	if len(nodes) == 0 {
//		return nil
//	}
//	var ret []model.NodeResponse
//	for _, node := range nodes {
//		ret = append(ret, *BuildNode(&node))
//	}
//	return ret
//}

//func BuildTopic(topic *model.Topic) *model.TopicResponse {
//	if topic == nil {
//		return nil
//	}
//
//	rsp := &model.TopicResponse{}
//
//	rsp.TopicId = topic.Id
//	rsp.Title = topic.Title
//	rsp.User = BuildUserDefaultIfNull(topic.UserId)
//	rsp.LastCommentTime = topic.LastCommentTime
//	rsp.CreateTime = topic.CreateTime
//	rsp.ViewCount = topic.ViewCount
//	rsp.CommentCount = topic.CommentCount
//	rsp.LikeCount = topic.LikeCount
//
//	if topic.NodeId > 0 {
//		node := services.TopicNodeService.Get(topic.NodeId)
//		rsp.Node = BuildNode(node)
//	}
//
//	tags := services.TopicService.GetTopicTags(topic.Id)
//	rsp.Tags = BuildTags(tags)
//
//	mr := simple.NewMd(simple.MdWithTOC()).Run(topic.Content)
//	rsp.Content = template.HTML(BuildHtmlContent(mr.ContentHtml))
//	rsp.Toc = template.HTML(mr.TocHtml)
//
//	if len(topic.ImageList) > 0 {
//		if err := simple.ParseJson(topic.ImageList, &rsp.ImageList); err != nil {
//			logrus.Error(err)
//		}
//	}
//
//	return rsp
//}

//func BuildSimpleTopic(topic *model.Topic) *model.TopicSimpleResponse {
//	if topic == nil {
//		return nil
//	}
//
//	rsp := &model.TopicSimpleResponse{}
//
//	rsp.TopicId = topic.Id
//	rsp.Title = topic.Title
//	rsp.User = BuildUserDefaultIfNull(topic.UserId)
//	rsp.LastCommentTime = topic.LastCommentTime
//	rsp.CreateTime = topic.CreateTime
//	rsp.ViewCount = topic.ViewCount
//	rsp.CommentCount = topic.CommentCount
//	rsp.LikeCount = topic.LikeCount
//
//	if len(topic.ImageList) > 0 {
//		if err := simple.ParseJson(topic.ImageList, &rsp.ImageList); err != nil {
//			logrus.Error(err)
//		}
//	}
//
//	if topic.NodeId > 0 {
//		node := services.TopicNodeService.Get(topic.NodeId)
//		rsp.Node = BuildNode(node)
//	}
//
//	tags := services.TopicService.GetTopicTags(topic.Id)
//	rsp.Tags = BuildTags(tags)
//	return rsp
//}

//func BuildSimpleTopics(topics []model.Topic) []model.TopicSimpleResponse {
//	if topics == nil || len(topics) == 0 {
//		return nil
//	}
//	var responses []model.TopicSimpleResponse
//	for _, topic := range topics {
//		responses = append(responses, *BuildSimpleTopic(&topic))
//	}
//	return responses
//}
//构造矿机产品信息
func BuildProject(project *model.Project) *model.ProjectResponse {
	if project == nil {
		return nil
	}
	rsp := &model.ProjectResponse{}
	rsp.Id = project.Id
	rsp.Name = project.Name
	rsp.Logo = project.Logo
	rsp.Type = project.Type
	rsp.Price = project.Price
	rsp.Capacity = project.Capacity
	rsp.Content = project.Content
	//if project.ContentType == model.ContentTypeHtml {
	//	rsp.Content = template.HTML(BuildHtmlContent(project.Content))
	//	rsp.Summary = simple.GetSummary(simple.GetHtmlText(project.Content), 256)
	//} else {
	//	mr := simple.NewMd().Run(project.Content)
	//	rsp.Content = template.HTML(BuildHtmlContent(mr.ContentHtml))
	//	rsp.Summary = mr.SummaryText
	//}

	return rsp
}

//func BuildSimpleProjects(projects []model.Project) []model.ProjectSimpleResponse {
//	if projects == nil || len(projects) == 0 {
//		return nil
//	}
//	var responses []model.ProjectSimpleResponse
//	for _, project := range projects {
//		responses = append(responses, *BuildSimpleProject(&project))
//	}
//	return responses
//}
//
//func BuildSimpleProject(project *model.Project) *model.ProjectSimpleResponse {
//	if project == nil {
//		return nil
//	}
//	rsp := &model.ProjectSimpleResponse{}
//	rsp.ProjectId = project.ID
//	rsp.Name = project.Name
//	rsp.Logo = project.Logo
//	rsp.CreateTime = project.CreatedAt.Unix()
//	rsp.Price = project.Price
//	rsp.Capacity = project.Capacity
//	rsp.Summary = project.Content
//	rsp.Type = project.Type
//
//	//if project.ContentType == model.ContentTypeHtml {
//	//	rsp.Summary = simple.GetSummary(simple.GetHtmlText(project.Content), 256)
//	//} else {
//	//	rsp.Summary = common.GetMarkdownSummary(project.Content)
//	//}
//
//	return rsp
//}

//func BuildComments(comments []model.Comment) []model.CommentResponse {
//	var ret []model.CommentResponse
//	for _, comment := range comments {
//		ret = append(ret, *BuildComment(comment))
//	}
//	return ret
//}
//
//func BuildComment(comment model.Comment) *model.CommentResponse {
//	return _buildComment(&comment, true)
//}
//
//func _buildComment(comment *model.Comment, buildQuote bool) *model.CommentResponse {
//	if comment == nil {
//		return nil
//	}
//
//	ret := &model.CommentResponse{
//		CommentId:  comment.Id,
//		User:       BuildUserDefaultIfNull(comment.UserId),
//		EntityType: comment.EntityType,
//		EntityId:   comment.EntityId,
//		QuoteId:    comment.QuoteId,
//		Status:     comment.Status,
//		CreateTime: comment.CreateTime,
//	}
//
//	if comment.ContentType == model.ContentTypeMarkdown {
//		markdownResult := simple.NewMd().Run(comment.Content)
//		ret.Content = template.HTML(BuildHtmlContent(markdownResult.ContentHtml))
//	} else {
//		ret.Content = template.HTML(BuildHtmlContent(comment.Content))
//	}
//
//	if buildQuote && comment.QuoteId > 0 {
//		quote := _buildComment(services.CommentService.Get(comment.QuoteId), false)
//		if quote != nil {
//			ret.Quote = quote
//			ret.QuoteContent = template.HTML(quote.User.Nickname+"：") + quote.Content
//		}
//	}
//	return ret
//}

func BuildTag(tag *model.Tag) *model.TagResponse {
	if tag == nil {
		return nil
	}
	return &model.TagResponse{TagId: tag.Id, TagName: tag.Name}
}

func BuildTags(tags []model.Tag) *[]model.TagResponse {
	if len(tags) == 0 {
		return nil
	}
	var responses []model.TagResponse
	for _, tag := range tags {
		responses = append(responses, *BuildTag(&tag))
	}
	return &responses
}

//func BuildFavorite(favorite *model.Favorite) *model.FavoriteResponse {
//	rsp := &model.FavoriteResponse{}
//	rsp.FavoriteId = favorite.Id
//	rsp.EntityType = favorite.EntityType
//	rsp.CreateTime = favorite.CreateTime
//
//	if favorite.EntityType == model.EntityTypeArticle {
//		article := services.ArticleService.Get(favorite.EntityId)
//		if article == nil || article.Status != model.StatusOk {
//			rsp.Deleted = true
//		} else {
//			rsp.Url = urls.ArticleUrl(article.Id)
//			rsp.User = BuildUserById(article.UserId)
//			rsp.Title = article.Title
//			if article.ContentType == model.ContentTypeMarkdown {
//				rsp.Content = common.GetMarkdownSummary(article.Content)
//			} else {
//				doc, err := goquery.NewDocumentFromReader(strings.NewReader(article.Content))
//				if err == nil {
//					text := doc.Text()
//					rsp.Content = simple.GetSummary(text, 256)
//				}
//			}
//		}
//	} else {
//		topic := services.TopicService.Get(favorite.EntityId)
//		if topic == nil || topic.Status != model.StatusOk {
//			rsp.Deleted = true
//		} else {
//			rsp.Url = urls.TopicUrl(topic.Id)
//			rsp.User = BuildUserById(topic.UserId)
//			rsp.Title = topic.Title
//			rsp.Content = common.GetMarkdownSummary(topic.Content)
//		}
//	}
//	return rsp
//}
//
//func BuildFavorites(favorites []model.Favorite) []model.FavoriteResponse {
//	if favorites == nil || len(favorites) == 0 {
//		return nil
//	}
//	var responses []model.FavoriteResponse
//	for _, favorite := range favorites {
//		responses = append(responses, *BuildFavorite(&favorite))
//	}
//	return responses
//}
//
//func BuildMessage(message *model.Message) *model.MessageResponse {
//	if message == nil {
//		return nil
//	}
//
//	detailUrl := ""
//	if message.Type == model.MsgTypeComment {
//		entityType := gjson.Get(message.ExtraData, "entityType")
//		entityId := gjson.Get(message.ExtraData, "entityId")
//		if entityType.String() == model.EntityTypeArticle {
//			detailUrl = urls.ArticleUrl(entityId.Int())
//		} else if entityType.String() == model.EntityTypeTopic {
//			detailUrl = urls.TopicUrl(entityId.Int())
//		}
//	}
//	from := BuildUserDefaultIfNull(message.FromId)
//	if message.FromId <= 0 {
//		from.Nickname = "系统通知"
//		from.Avatar = avatar.DefaultAvatar
//	}
//	return &model.MessageResponse{
//		MessageId:    message.Id,
//		From:         from,
//		UserId:       message.UserId,
//		Content:      message.Content,
//		QuoteContent: message.QuoteContent,
//		Type:         message.Type,
//		DetailUrl:    detailUrl,
//		ExtraData:    message.ExtraData,
//		Status:       message.Status,
//		CreateTime:   message.CreateTime,
//	}
//}
//
//func BuildMessages(messages []model.Message) []model.MessageResponse {
//	if len(messages) == 0 {
//		return nil
//	}
//	var responses []model.MessageResponse
//	for _, message := range messages {
//		responses = append(responses, *BuildMessage(&message))
//	}
//	return responses
//}
//
func BuildHtmlContent(htmlContent string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return htmlContent
	}

	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		href := selection.AttrOr("href", "")

		if len(href) == 0 {
			return
		}

		// 不是内部链接
		if !urls.IsInternalUrl(href) {
			selection.SetAttr("target", "_blank")
			selection.SetAttr("rel", "external nofollow") // 标记站外链接，搜索引擎爬虫不传递权重值

			config := services.SysConfigService.GetConfig()
			if config.UrlRedirect { // 开启非内部链接跳转
				newHref := simple.ParseUrl(urls.AbsUrl("/redirect")).AddQuery("url", href).BuildStr()
				selection.SetAttr("href", newHref)
			}
		}

		// 如果是锚链接
		if urls.IsAnchor(href) {
			selection.ReplaceWithHtml(selection.Text())
		}

		// 如果a标签没有title，那么设置title
		title := selection.AttrOr("title", "")
		if len(title) == 0 {
			selection.SetAttr("title", selection.Text())
		}
	})

	// 处理图片
	doc.Find("img").Each(func(i int, selection *goquery.Selection) {
		src := selection.AttrOr("src", "")
		if strings.Contains(src, "qpic.cn") {
			newSrc := simple.ParseUrl("/api/img/proxy").AddQuery("url", src).BuildStr()
			selection.SetAttr("src", newSrc)
		}
	})

	html, err := doc.Html()
	if err != nil {
		return htmlContent
	}
	return html
}
