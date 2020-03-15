package services

import (
	"bbs-go/model"
	"bbs-go/repositories"
	"github.com/mlogclub/simple"
)

var UserAssertLogService = newUserAssertLogService()

func newUserAssertLogService() *userAssertLogService {
	return &userAssertLogService{}
}

type userAssertLogService struct {
}

func (s *userAssertLogService) Get(id int64) *model.UserAssertLog {
	return repositories.UserAssertLogRepository.Get(simple.DB(), id)
}

func (s *userAssertLogService) Take(where ...interface{}) *model.UserAssertLog {
	return repositories.UserAssertLogRepository.Take(simple.DB(), where...)
}

func (s *userAssertLogService) Find(cnd *simple.SqlCnd) []model.UserAssertLog {
	return repositories.UserAssertLogRepository.Find(simple.DB(), cnd)
}

func (s *userAssertLogService) FindOne(cnd *simple.SqlCnd) *model.UserAssertLog {
	return repositories.UserAssertLogRepository.FindOne(simple.DB(), cnd)
}

func (s *userAssertLogService) FindPageByParams(params *simple.QueryParams) (list []model.UserAssertLog, paging *simple.Paging) {
	return repositories.UserAssertLogRepository.FindPageByParams(simple.DB(), params)
}

func (s *userAssertLogService) FindPageByCnd(cnd *simple.SqlCnd) (list []model.UserAssertLog, paging *simple.Paging) {
	return repositories.UserAssertLogRepository.FindPageByCnd(simple.DB(), cnd)
}

func (s *userAssertLogService) Count(cnd *simple.SqlCnd) int {
	return repositories.UserAssertLogRepository.Count(simple.DB(), cnd)
}

func (s *userAssertLogService) Create(t *model.UserAssertLog) error {
	return repositories.UserAssertLogRepository.Create(simple.DB(), t)
}

func (s *userAssertLogService) Update(t *model.UserAssertLog) error {
	return repositories.UserAssertLogRepository.Update(simple.DB(), t)
}

func (s *userAssertLogService) Updates(id int64, columns map[string]interface{}) error {
	return repositories.UserAssertLogRepository.Updates(simple.DB(), id, columns)
}

func (s *userAssertLogService) UpdateColumn(id int64, name string, value interface{}) error {
	return repositories.UserAssertLogRepository.UpdateColumn(simple.DB(), id, name, value)
}

func (s *userAssertLogService) Delete(id int64) {
	repositories.UserAssertLogRepository.Delete(simple.DB(), id)
}
