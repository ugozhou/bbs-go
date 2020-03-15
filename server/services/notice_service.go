package services

import (
	"bbs-go/model"
	"bbs-go/repositories"
	"github.com/mlogclub/simple"
)

var NoticeService = newNoticeService()

func newNoticeService() *noticeService {
	return &noticeService{}
}

type noticeService struct {
}

func (s *noticeService) Get(id int64) *model.Notice {
	return repositories.NoticeRepository.Get(simple.DB(), id)
}

func (s *noticeService) Take(where ...interface{}) *model.Notice {
	return repositories.NoticeRepository.Take(simple.DB(), where...)
}

func (s *noticeService) Find(cnd *simple.SqlCnd) []model.Notice {
	return repositories.NoticeRepository.Find(simple.DB(), cnd)
}

func (s *noticeService) FindOne(cnd *simple.SqlCnd) *model.Notice {
	return repositories.NoticeRepository.FindOne(simple.DB(), cnd)
}

func (s *noticeService) FindPageByParams(params *simple.QueryParams) (list []model.Notice, paging *simple.Paging) {
	return repositories.NoticeRepository.FindPageByParams(simple.DB(), params)
}

func (s *noticeService) FindPageByCnd(cnd *simple.SqlCnd) (list []model.Notice, paging *simple.Paging) {
	return repositories.NoticeRepository.FindPageByCnd(simple.DB(), cnd)
}

func (s *noticeService) Count(cnd *simple.SqlCnd) int {
	return repositories.NoticeRepository.Count(simple.DB(), cnd)
}

func (s *noticeService) Create(t *model.Notice) error {
	return repositories.NoticeRepository.Create(simple.DB(), t)
}

func (s *noticeService) Update(t *model.Notice) error {
	return repositories.NoticeRepository.Update(simple.DB(), t)
}

func (s *noticeService) Updates(id int64, columns map[string]interface{}) error {
	return repositories.NoticeRepository.Updates(simple.DB(), id, columns)
}

func (s *noticeService) UpdateColumn(id int64, name string, value interface{}) error {
	return repositories.NoticeRepository.UpdateColumn(simple.DB(), id, name, value)
}

func (s *noticeService) Delete(id int64) {
	repositories.NoticeRepository.Delete(simple.DB(), id)
}
