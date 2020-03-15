package services

import (
	"bbs-go/model"
	"bbs-go/repositories"
	"bbs-go/services/cache"
	"errors"
	"github.com/mlogclub/simple"
)

var UserScoreService = newUserScoreService()

func newUserScoreService() *userScoreService {
	return &userScoreService{}
}

type userScoreService struct {
}

func (s *userScoreService) Get(id int64) *model.UserScore {
	return repositories.UserScoreRepository.Get(simple.DB(), id)
}

func (s *userScoreService) Take(where ...interface{}) *model.UserScore {
	return repositories.UserScoreRepository.Take(simple.DB(), where...)
}

func (s *userScoreService) Find(cnd *simple.SqlCnd) []model.UserScore {
	return repositories.UserScoreRepository.Find(simple.DB(), cnd)
}

func (s *userScoreService) FindOne(cnd *simple.SqlCnd) *model.UserScore {
	return repositories.UserScoreRepository.FindOne(simple.DB(), cnd)
}

func (s *userScoreService) FindPageByParams(params *simple.QueryParams) (list []model.UserScore, paging *simple.Paging) {
	return repositories.UserScoreRepository.FindPageByParams(simple.DB(), params)
}

func (s *userScoreService) FindPageByCnd(cnd *simple.SqlCnd) (list []model.UserScore, paging *simple.Paging) {
	return repositories.UserScoreRepository.FindPageByCnd(simple.DB(), cnd)
}

func (s *userScoreService) Create(t *model.UserScore) error {
	return repositories.UserScoreRepository.Create(simple.DB(), t)
}

func (s *userScoreService) Update(t *model.UserScore) error {
	return repositories.UserScoreRepository.Update(simple.DB(), t)
}

func (s *userScoreService) Updates(id int64, columns map[string]interface{}) error {
	return repositories.UserScoreRepository.Updates(simple.DB(), id, columns)
}

func (s *userScoreService) UpdateColumn(id int64, name string, value interface{}) error {
	return repositories.UserScoreRepository.UpdateColumn(simple.DB(), id, name, value)
}

func (s *userScoreService) Delete(id int64) {
	repositories.UserScoreRepository.Delete(simple.DB(), id)
}

func (s *userScoreService) GetByUserId(userId int64) *model.UserScore {
	return s.FindOne(simple.NewSqlCnd().Eq("user_id", userId))
}

func (s *userScoreService) CreateOrUpdate(t *model.UserScore) error {
	if t.Id > 0 {
		return s.Update(t)
	} else {
		return s.Create(t)
	}
}

//自购买矿机
func (s *userScoreService) IncrementSelfBuyScore(user *model.User, amount int) (error, *model.UserScore) {

	err, userScore := s.addScore(user.Id, amount, model.SelfBuy, "购买矿机")
	if err != nil {
		return err, nil
	}
	s.checkAndLevlup(userScore)
	//查找用户推荐人，处理推荐奖励逻辑
	//此处暂且给介绍人增加1T
	err, userScore = s.addScore(user.Introducer, 1, model.InviatorReward, "邀请奖励")
	if err != nil {
		return err, nil
	}
	s.checkAndLevlup(userScore)
	//检查用户的等级变化

	return nil, userScore
}

//检查并升级用户等级
func (s *userScoreService) checkAndLevlup(userScore *model.UserScore) {
	oldLevel := userScore.Level
	selfScore := userScore.Score
	rewardScore := userScore.RewardScore
	totalScore := selfScore + rewardScore
	level := 0

	if selfScore >= 300 {
		level = 5
	} else if selfScore >= 100 && totalScore >= 2000 {
		level = 4
	} else if selfScore >= 50 && totalScore >= 500 {
		level = 3
	} else if selfScore >= 10 && totalScore >= 100 {
		level = 2
	} else if selfScore >= 0 && totalScore > 0 {
		level = 1
	}
	if oldLevel != level {
		userScore.Level = level
		repositories.UserScoreRepository.Update(simple.DB(), userScore)
	}

}

//// IncrementPostCommentScore 跟帖获积分
//func (s *userScoreService) IncrementPostCommentScore(comment *model.Comment) {
//	// 非话题跟帖，跳过
//	if comment.EntityType != model.EntityTypeTopic {
//		return
//	}
//	config := SysConfigService.GetConfig()
//	if config.ScoreConfig.IntroducerRewardScore <= 0 {
//		logrus.Info("请配置跟帖积分")
//		return
//	}
//	err := s.addScore(comment.UserId, config.ScoreConfig.IntroducerRewardScore, model.EntityTypeComment,
//		strconv.FormatInt(comment.Id, 10), "发表跟帖")
//	if err != nil {
//		logrus.Error(err)
//	}
//}

//// Increment 增加算力
//func (s *userScoreService) Increment(userId int64, score int, sourceType, sourceId, description string) error {
//	if score <= 0 {
//		return errors.New("分数必须为正数")
//	}
//	return s.addScore(userId, score, sourceType, sourceId, description)
//}

//// Decrement 减少分数
//func (s *userScoreService) Decrement(userId int64, score int, sourceType, sourceId, description string) error {
//	if score <= 0 {
//		return errors.New("分数必须为正数")
//	}
//	return s.addScore(userId, -score, sourceType, sourceId, description)
//}

// addScore 加分数，也可以加负数
func (s *userScoreService) addScore(userId int64, score int, catalog int, description string) (error, *model.UserScore) {
	if score == 0 {
		return errors.New("分数不能为0"), nil
	}
	userScore := s.GetByUserId(userId)
	if userScore == nil {
		userScore = &model.UserScore{
			UserId:     userId,
			CreateTime: simple.NowTimestamp(),
		}
	}
	userScore.Score = userScore.Score + score
	userScore.UpdateTime = simple.NowTimestamp()
	if err := s.CreateOrUpdate(userScore); err != nil {
		return err, nil
	}

	scoreType := model.ScoreTypeIncr
	if score < 0 {
		scoreType = model.ScoreTypeDecr
	}
	err := UserScoreLogService.Create(&model.UserScoreLog{
		UserId:      userId,
		Catalog:     catalog,
		Description: description,
		Type:        scoreType,
		Score:       score,
		CreateTime:  simple.NowTimestamp(),
	})
	if err == nil {
		cache.UserCache.InvalidateScore(userId)
	}
	return err, userScore
}
