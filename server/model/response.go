package model

import (
	"html/template"
)

type UserInfo struct {
	Id          int64    `json:"id"`
	Email       string   `json:"email"`
	Nickname    string   `json:"nickname"`
	Avatar      string   `json:"avatar"`
	Type        int      `json:"type"`
	Roles       []string `json:"roles"`
	Phone       string   `json:"phone"`
	Score       int      `json:"score"`       // 积分
	PasswordSet bool     `json:"passwordSet"` // 密码已设置
	Status      int      `json:"status"`
	CreateTime  int64    `json:"createTime"`
}

func (info *UserInfo) HasRole(role string) bool {
	if len(info.Roles) == 0 {
		return false
	}
	for _, r := range info.Roles {
		if r == role {
			return true
		}
	}
	return false
}

type TagResponse struct {
	TagId   int64  `json:"tagId"`
	TagName string `json:"tagName"`
}

type ArticleSimpleResponse struct {
	ArticleId  int64          `json:"articleId"`
	User       *UserInfo      `json:"user"`
	Tags       *[]TagResponse `json:"tags"`
	Title      string         `json:"title"`
	Summary    string         `json:"summary"`
	Share      bool           `json:"share"`
	SourceUrl  string         `json:"sourceUrl"`
	ViewCount  int64          `json:"viewCount"`
	CreateTime int64          `json:"createTime"`
}

type ArticleResponse struct {
	ArticleSimpleResponse
	Content template.HTML `json:"content"`
	Toc     template.HTML `json:"toc"`
}

type NodeResponse struct {
	NodeId      int64  `json:"nodeId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// 帖子列表返回实体
type TopicSimpleResponse struct {
	TopicId         int64          `json:"topicId"`
	User            *UserInfo      `json:"user"`
	Node            *NodeResponse  `json:"node"`
	Tags            *[]TagResponse `json:"tags"`
	Title           string         `json:"title"`
	ImageList       *[]string      `json:"imageList"`
	LastCommentTime int64          `json:"lastCommentTime"`
	ViewCount       int64          `json:"viewCount"`
	CommentCount    int64          `json:"commentCount"`
	LikeCount       int64          `json:"likeCount"`
	Liked           bool           `json:"liked"`
	CreateTime      int64          `json:"createTime"`
}

// 帖子详情返回实体
type TopicResponse struct {
	TopicSimpleResponse
	Content template.HTML `json:"content"`
	Toc     template.HTML `json:"toc"`
}

// 项目详情
type ProjectResponse struct {
	Project
}

type CommentResponse struct {
	CommentId    int64            `json:"commentId"`
	User         *UserInfo        `json:"user"`
	EntityType   string           `json:"entityType"`
	EntityId     int64            `json:"entityId"`
	Content      template.HTML    `json:"content"`
	QuoteId      int64            `json:"quoteId"`
	Quote        *CommentResponse `json:"quote"`
	QuoteContent template.HTML    `json:"quoteContent"`
	Status       int              `json:"status"`
	CreateTime   int64            `json:"createTime"`
}

type FavoriteResponse struct {
	FavoriteId int64     `json:"favoriteId"`
	EntityType string    `json:"entityType"`
	EntityId   int64     `json:"entityId"`
	Deleted    bool      `json:"deleted"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	User       *UserInfo `json:"user"`
	Url        string    `json:"url"`
	CreateTime int64     `json:"createTime"`
}

// 消息
type MessageResponse struct {
	MessageId    int64     `json:"messageId"`
	From         *UserInfo `json:"from"`    // 消息发送人
	UserId       int64     `json:"userId"`  // 消息接收人编号
	Content      string    `json:"content"` // 消息内容
	QuoteContent string    `json:"quoteContent"`
	Type         int       `json:"type"`
	DetailUrl    string    `json:"detailUrl"` // 消息详情url
	ExtraData    string    `json:"extraData"`
	Status       int       `json:"status"`
	CreateTime   int64     `json:"createTime"`
}
