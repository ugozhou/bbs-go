package model

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

var Models = []interface{}{
	&User{}, &UserToken{}, &Tag{}, &Article{}, &ArticleTag{}, &SysConfig{}, &Project{},
	&ThirdAccount{}, &UserScore{}, &UserScoreLog{}, &UserAssert{}, &UserAssertLog{}, &Waiter{}, &Notice{},
}

type Model struct {
	Id int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id" form:"id"`
}

type ModelInt struct {
	Id int `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id" form:"id"`
}

const (
	StatusOk      = 0 // 正常
	StatusDeleted = 1 // 删除
	StatusPending = 2 // 待审核

	UserTypeNormal = 0 // 普通用户
	UserTypeGzh    = 1 // 公众号用户

	ContentTypeHtml     = "html"
	ContentTypeMarkdown = "markdown"

	EntityTypeArticle = "article"
	EntityTypeTopic   = "topic"
	EntityTypeComment = "comment"

	MsgStatusUnread = 0 // 消息未读
	MsgStatusReaded = 1 // 消息已读

	MsgTypeComment = 0 // 回复消息

	ThirdAccountTypeGithub = "github"
	ThirdAccountTypeQQ     = "qq"
	ThirdAccountTypeWechat = "wechat"

	ScoreTypeIncr = 0 // 积分+
	ScoreTypeDecr = 1 // 积分-

	TopicTypeNormal  = 0 // 普通帖子
	TopicTypeTwitter = 1 // 推文
)

type User struct {
	Model
	Mobile      string `gorm:"size:32;unique;" json:"mobile" form:"mobile"`                // 用户名
	Nickname    string `gorm:"size:16;" json:"nickname" form:"nickname"`                   // 昵称
	Avatar      string `gorm:"type:text" json:"avatar" form:"avatar"`                      // 头像
	Password    string `gorm:"size:512" json:"password" form:"password"`                   // 密码
	Paypassword string `gorm:"size:16" json:"payPassword" form:"payPassword"`              // 支付密码
	Status      int    `gorm:"index:idx_user_status;not null" json:"status" form:"status"` // 状态
	Invitecode  string `gorm:"not null;size:512" json:"invitecode" form:"invitecode"`      // 邀请码
	Roles       string `gorm:"type:text" json:"roles" form:"roles"`                        // 角色
	Type        int    `gorm:"not null" json:"type" form:"type"`                           // 用户类型
	CreateTime  int64  `json:"createTime" form:"createTime"`                               // 创建时间
	UpdateTime  int64  `json:"updateTime" form:"updateTime"`                               // 更新时间
	Introducer  int64  `gorm:not null; json:"introducer" form:"introducer"`                //介绍人
}

type UserToken struct {
	Model
	Token      string `gorm:"size:32;unique;not null" json:"token" form:"token"`
	UserId     int64  `gorm:"not null;index:idx_user_token_user_id;" json:"userId" form:"userId"`
	ExpiredAt  int64  `gorm:"not null" json:"expiredAt" form:"expiredAt"`
	Status     int    `gorm:"not null;index:idx_user_token_status" json:"status" form:"status"`
	CreateTime int64  `gorm:"not null" json:"createTime" form:"createTime"`
}

type ThirdAccount struct {
	Model
	UserId     sql.NullInt64 `gorm:"unique_index:idx_user_id_third_type;" json:"userId" form:"userId"`                                  // 用户编号
	Avatar     string        `gorm:"size:1024" json:"avatar" form:"avatar"`                                                             // 头像
	Nickname   string        `gorm:"size:32" json:"nickname" form:"nickname"`                                                           // 昵称
	ThirdType  string        `gorm:"size:32;not null;unique_index:idx_user_id_third_type,idx_third;" json:"thirdType" form:"thirdType"` // 第三方类型
	ThirdId    string        `gorm:"size:64;not null;unique_index:idx_third;" json:"thirdId" form:"thirdId"`                            // 第三方唯一标识，例如：openId,unionId
	ExtraData  string        `gorm:"type:longtext" json:"extraData" form:"extraData"`                                                   // 扩展数据
	CreateTime int64         `json:"createTime" form:"createTime"`                                                                      // 创建时间
	UpdateTime int64         `json:"updateTime" form:"updateTime"`                                                                      // 更新时间
}

// 标签
type Tag struct {
	Model
	Name        string `gorm:"size:32;unique;not null" json:"name" form:"name"`
	Description string `gorm:"size:1024" json:"description" form:"description"`
	Status      int    `gorm:"index:idx_tag_status;not null" json:"status" form:"status"`
	CreateTime  int64  `json:"createTime" form:"createTime"`
	UpdateTime  int64  `json:"updateTime" form:"updateTime"`
}

// 文章
type Article struct {
	Model
	UserId      int64  `gorm:"index:idx_article_user_id" json:"userId" form:"userId"`             // 所属用户编号
	Title       string `gorm:"size:128;not null;" json:"title" form:"title"`                      // 标题
	Summary     string `gorm:"type:text" json:"summary" form:"summary"`                           // 摘要
	Content     string `gorm:"type:longtext;not null;" json:"content" form:"content"`             // 内容
	ContentType string `gorm:"type:varchar(32);not null" json:"contentType" form:"contentType"`   // 内容类型：markdown、html
	Status      int    `gorm:"int;not null;index:idx_article_status" json:"status" form:"status"` // 状态
	Share       bool   `gorm:"not null" json:"share" form:"share"`                                // 是否是分享的文章，如果是这里只会显示文章摘要，原文需要跳往原链接查看
	SourceUrl   string `gorm:"type:text" json:"sourceUrl" form:"sourceUrl"`                       // 原文链接
	ViewCount   int64  `gorm:"not null;index:idx_view_count;" json:"viewCount" form:"viewCount"`  // 查看数量
	CreateTime  int64  `gorm:"index:idx_article_create_time" json:"createTime" form:"createTime"` // 创建时间
	UpdateTime  int64  `json:"updateTime" form:"updateTime"`                                      // 更新时间
}

// 文章标签
type ArticleTag struct {
	Model
	ArticleId  int64 `gorm:"not null;index:idx_article_id;" json:"articleId" form:"articleId"`  // 文章编号
	TagId      int64 `gorm:"not null;index:idx_article_tag_tag_id;" json:"tagId" form:"tagId"`  // 标签编号
	Status     int64 `gorm:"not null;index:idx_article_tag_status" json:"status" form:"status"` // 状态：正常、删除
	CreateTime int64 `json:"createTime" form:"createTime"`                                      // 创建时间
}

// 评论
type Comment struct {
	Model
	UserId      int64  `gorm:"index:idx_comment_user_id;not null" json:"userId" form:"userId"`             // 用户编号
	EntityType  string `gorm:"index:idx_comment_entity_type;not null" json:"entityType" form:"entityType"` // 被评论实体类型
	EntityId    int64  `gorm:"index:idx_comment_entity_id;not null" json:"entityId" form:"entityId"`       // 被评论实体编号
	Content     string `gorm:"type:text;not null" json:"content" form:"content"`                           // 内容
	ContentType string `gorm:"type:varchar(32);not null" json:"contentType" form:"contentType"`            // 内容类型：markdown、html
	QuoteId     int64  `gorm:"not null"  json:"quoteId" form:"quoteId"`                                    // 引用的评论编号
	Status      int    `gorm:"int;index:idx_comment_status" json:"status" form:"status"`                   // 状态：0：待审核、1：审核通过、2：审核失败、3：已发布
	CreateTime  int64  `json:"createTime" form:"createTime"`                                               // 创建时间
}

// 收藏
type Favorite struct {
	Model
	UserId     int64  `gorm:"index:idx_favorite_user_id;not null" json:"userId" form:"userId"`                     // 用户编号
	EntityType string `gorm:"index:idx_favorite_entity_type;size:32;not null" json:"entityType" form:"entityType"` // 收藏实体类型
	EntityId   int64  `gorm:"index:idx_favorite_entity_id;not null" json:"entityId" form:"entityId"`               // 收藏实体编号
	CreateTime int64  `json:"createTime" form:"createTime"`                                                        // 创建时间
}

// 话题节点
type TopicNode struct {
	Model
	Name        string `gorm:"size:32;unique" json:"name" form:"name"`        // 名称
	Description string `json:"description" form:"description"`                // 描述
	SortNo      int    `gorm:"index:idx_sort_no" json:"sortNo" form:"sortNo"` // 排序编号
	Status      int    `gorm:"not null" json:"status" form:"status"`          // 状态
	CreateTime  int64  `json:"createTime" form:"createTime"`                  // 创建时间
}

// 话题节点
type Topic struct {
	Model
	Type            int    `gorm:"not null;index:idx_topic_type" json:"type" form:"type"`                           // 类型
	NodeId          int64  `gorm:"not null;index:idx_node_id;" json:"nodeId" form:"nodeId"`                         // 节点编号
	UserId          int64  `gorm:"not null;index:idx_topic_user_id;" json:"userId" form:"userId"`                   // 用户
	Title           string `gorm:"size:128" json:"title" form:"title"`                                              // 标题
	Content         string `gorm:"type:longtext" json:"content" form:"content"`                                     // 内容
	ImageList       string `gorm:"type:longtext" json:"imageList" form:"imageList"`                                 // 图片
	Recommend       bool   `gorm:"not null;index:idx_recommend" json:"recommend" form:"recommend"`                  // 是否推荐
	ViewCount       int64  `gorm:"not null" json:"viewCount" form:"viewCount"`                                      // 查看数量
	CommentCount    int64  `gorm:"not null" json:"commentCount" form:"commentCount"`                                // 跟帖数量
	LikeCount       int64  `gorm:"not null" json:"likeCount" form:"likeCount"`                                      // 点赞数量
	Status          int    `gorm:"index:idx_topic_status;" json:"status" form:"status"`                             // 状态：0：正常、1：删除
	LastCommentTime int64  `gorm:"index:idx_topic_last_comment_time" json:"lastCommentTime" form:"lastCommentTime"` // 最后回复时间
	CreateTime      int64  `gorm:"index:idx_topic_create_time" json:"createTime" form:"createTime"`                 // 创建时间
	ExtraData       string `gorm:"type:text" json:"extraData" form:"extraData"`                                     // 扩展数据
}

// 主题标签
type TopicTag struct {
	Model
	TopicId         int64 `gorm:"not null;index:idx_topic_tag_topic_id;" json:"topicId" form:"topicId"`                // 主题编号
	TagId           int64 `gorm:"not null;index:idx_topic_tag_tag_id;" json:"tagId" form:"tagId"`                      // 标签编号
	Status          int64 `gorm:"not null;index:idx_topic_tag_status" json:"status" form:"status"`                     // 状态：正常、删除
	LastCommentTime int64 `gorm:"index:idx_topic_tag_last_comment_time" json:"lastCommentTime" form:"lastCommentTime"` // 最后回复时间
	CreateTime      int64 `json:"createTime" form:"createTime"`                                                        // 创建时间
}

// 话题点赞
type TopicLike struct {
	Model
	UserId     int64 `gorm:"not null;index:idx_topic_like_user_id;" json:"userId" form:"userId"`    // 用户
	TopicId    int64 `gorm:"not null;index:idx_topic_like_topic_id;" json:"topicId" form:"topicId"` // 主题编号
	CreateTime int64 `json:"createTime" form:"createTime"`                                          // 创建时间
}

// 消息
type Message struct {
	Model
	FromId       int64  `gorm:"not null" json:"fromId" form:"fromId"`                            // 消息发送人
	UserId       int64  `gorm:"not null;index:idx_message_user_id;" json:"userId" form:"userId"` // 用户编号(消息接收人)
	Content      string `gorm:"type:text;not null" json:"content" form:"content"`                // 消息内容
	QuoteContent string `gorm:"type:text" json:"quoteContent" form:"quoteContent"`               // 引用内容
	Type         int    `gorm:"not null" json:"type" form:"type"`                                // 消息类型
	ExtraData    string `gorm:"type:text" json:"extraData" form:"extraData"`                     // 扩展数据
	Status       int    `gorm:"not null" json:"status" form:"status"`                            // 状态：0：未读、1：已读
	CreateTime   int64  `json:"createTime" form:"createTime"`                                    // 创建时间
}

// 系统配置
type SysConfig struct {
	Model
	Key         string `gorm:"not null;size:128;unique" json:"key" form:"key"` // 配置key
	Value       string `gorm:"type:text" json:"value" form:"value"`            // 配置值
	Name        string `gorm:"not null;size:32" json:"name" form:"name"`       // 配置名称
	Description string `gorm:"size:128" json:"description" form:"description"` // 配置描述
	CreateTime  int64  `gorm:"not null" json:"createTime" form:"createTime"`   // 创建时间
	UpdateTime  int64  `gorm:"not null" json:"updateTime" form:"updateTime"`   // 更新时间
}

// 矿机数据
type Project struct {
	ModelInt
	Name     string `gorm:"type:varchar(128)" json:"name" form:"name"`   //矿机名字
	Logo     string `gorm:"type:varchar(1024)" json:"logo" form:"logo"`  //图片地址
	Type     uint   `gorm:"not null;" json:"type" form:"type"`           //矿机类型
	Price    int64  `gorm:"not null;" json:"price" form:"price"`         //矿机价格
	Capacity int    `gorm:"type:varchar(32);" json:"price" form:"price"` //矿机容量，1T，300T
	Content  string `gorm:"type:longtext" json:"content" form:"content"` //描述信息
}

// 友链
type Link struct {
	Model
	Url        string `gorm:"not null;type:text" json:"url" form:"url"`     // 链接
	Title      string `gorm:"not null;size:128" json:"title" form:"title"`  // 标题
	Summary    string `gorm:"size:1024" json:"summary" form:"summary"`      // 站点描述
	Logo       string `gorm:"type:text" json:"logo" form:"logo"`            // LOGO
	Status     int    `gorm:"not null" json:"status" form:"status"`         // 状态
	CreateTime int64  `gorm:"not null" json:"createTime" form:"createTime"` // 创建时间
}

// 站点地图
type Sitemap struct {
	Model
	Loc        string `gorm:"not null;size:1024" json:"loc" form:"loc"`              // loc
	Lastmod    int64  `gorm:"not null" json:"lastmod" form:"lastmod"`                // 最后更新时间
	LocName    string `gorm:"not null;size:32;unique" json:"locName" form:"locName"` // loc的md5
	CreateTime int64  `gorm:"not null" json:"createTime" form:"createTime"`          // 创建时间
}

// 用户算力矿机信息
type UserScore struct {
	Model
	UserId      int64 `gorm:"unique;not null" json:"userId" form:"userId"`    // 用户编号
	Score       int   `gorm:"not null" json:"score" form:"score"`             // 自购算力
	RewardScore int   `gorm:"not null" json:"rewardScore" form:"rewardScore"` // 奖励算力
	Level       int   `gorm:"not null" json:"level" form:"level"`             // 等级
	CreateTime  int64 `json:"createTime" form:"createTime"`                   // 创建时间
	UpdateTime  int64 `json:"updateTime" form:"updateTime"`                   // 更新时间
}

// 用户算力矿机流水信息
type UserScoreLog struct {
	Model
	UserId      int64  `gorm:"not null;index:idx_user_score_log_user_id" json:"userId" form:"userId"` // 用户编号
	Catalog     int    `gorm:"not null;index:idx_user_score_catalog" json:"catalog" form:"catalog"`   // 流水分类
	Description string `json:"description" form:"description"`                                        // 描述
	Type        int    `json:"type" form:"type"`                                                      // 类型(增加、减少)
	Score       int    `json:"score" form:"score"`                                                    // 积分
	CreateTime  int64  `json:"createTime" form:"createTime"`                                          // 创建时间
}

// 用户资产
type UserAssert struct {
	Model
	UserId  int64  `gorm:"unique;not null;idx_user_assert_user_id" json:"user_id" form:"userId"` //用户编号
	USDT    int64  `gorm:"default:0"json:"usdt" form:"usdt"`                                     //usdt数量
	Address string `gorm:"unique;notnull;idx_user_usdt_address" json:"address" form:"address"`   //usdt钱包充值地址
}

// USDT充值记录
type UserAssertLog struct {
	Model
	UserId      int64  `gorm:"not null;index:idx_user_assert_log_id" json:"user_id" form:"userId"` //用户编号
	Amount      int64  `gorm:"not null;"`                                                          //交易数量
	Type        uint8  `gorm:"not null" json:"type" form:"type"`                                   //交易类型（增加，减少）
	Description string `json:"description" form:"description"`                                     // 描述
}

// 用户FIL流水
type UserFILLog struct {
	gorm.Model
	UserId int64 `gorm:"unique;not null;idx_user_assert_user_id" json:"user_id" form:"userId"` //用户编号
	Amount int32 `gorm:"not null;"`                                                            //交易数量
	Type   uint8 `gorm:"not null"`                                                             //记录类型
}

//	客服信息
type Waiter struct {
	ModelInt
	Name   string `gorm:"unique;not null;" json:"name" form:"name"`     //客服名称
	Wechat string `gorm:"unique;not null;" json:"wechat" form:"wechat"` //微信号
	Phone  int    `gorm:"unique;not null;" json:"phone" form:"phone"`   //电话号码
}

// 公告信息
type Notice struct {
	ModelInt
	Title      string `gorm:"not null;" json:"title" form:"title"`           //公告标题
	Content    string `gorm:"not null;" json:"content" form:"content"`       //内容详情
	CreateTime int64  `gorm:"not null;" json:"createTime" form:"createTime"` //创建时间
}
