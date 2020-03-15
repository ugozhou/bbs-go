package model

// 站点导航
type SiteNav struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

// 算力配置
type ScoreConfig struct {
	SelfBuyScore          int `json:"postTopicScore"`   // 自身购买算力
	IntroducerRewardScore int `json:"postCommentScore"` // 介绍人奖励获取
}

// 配置返回结构体
type ConfigData struct {
	SiteTitle        string      `json:"siteTitle"`
	SiteDescription  string      `json:"siteDescription"`
	SiteKeywords     []string    `json:"siteKeywords"`
	SiteNavs         []SiteNav   `json:"siteNavs"`
	SiteNotification string      `json:"siteNotification"`
	RecommendTags    []string    `json:"recommendTags"`
	UrlRedirect      bool        `json:"urlRedirect"`
	ScoreConfig      ScoreConfig `json:"scoreConfig"`
}
