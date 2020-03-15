package model

// 系统配置
const (
	SysConfigSiteTitle        = "siteTitle"        // 站点标题
	SysConfigSiteDescription  = "siteDescription"  // 站点描述
	SysConfigSiteKeywords     = "siteKeywords"     // 站点关键字
	SysConfigSiteNavs         = "siteNavs"         // 站点导航
	SysConfigSiteNotification = "siteNotification" // 站点公告
	SysConfigRecommendTags    = "recommendTags"    // 推荐标签
	SysConfigUrlRedirect      = "urlRedirect"      // 是否开启链接跳转
	SysConfigScoreConfig      = "scoreConfig"      // 分数配置
)

// 图片样式
const (
	ImageStyleAvatar = "avatar" // 头像样式
	ImageStyleDetail = "detail" // 图片详情样式
)

//USDT资源变更类型
const (
	Recharge = iota + 1 //充值
	BuyMiner            //购买矿机
)

const (
	SelfBuy        = iota + 1 //自购
	InviatorReward            //邀请奖励
)

const (
	Level0   = "s0"
	Level1   = "s1"
	LevelV1  = "v1"
	Levelv2  = "v2"
	Levelv3  = "v3"
	LevelVip = "vip"
)
