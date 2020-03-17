package services

import (
	"bbs-go/model"
	"bbs-go/repositories"
	"github.com/mlogclub/simple"
)

var BannerService = newBannerService()

func newBannerService() *bannerService {
	return &bannerService{}
}

type bannerService struct {
}

func (s *bannerService) Get(id int64) *model.Banner {
	return repositories.BannerRepository.Get(simple.DB(), id)
}

func (s *bannerService) Take(where ...interface{}) *model.Banner {
	return repositories.BannerRepository.Take(simple.DB(), where...)
}

func (s *bannerService) Find(cnd *simple.SqlCnd) []model.Banner {
	return repositories.BannerRepository.Find(simple.DB(), cnd)
}

func (s *bannerService) FindOne(cnd *simple.SqlCnd) *model.Banner {
	return repositories.BannerRepository.FindOne(simple.DB(), cnd)
}

func (s *bannerService) FindPageByParams(params *simple.QueryParams) (list []model.Banner, paging *simple.Paging) {
	return repositories.BannerRepository.FindPageByParams(simple.DB(), params)
}

func (s *bannerService) FindPageByCnd(cnd *simple.SqlCnd) (list []model.Banner, paging *simple.Paging) {
	return repositories.BannerRepository.FindPageByCnd(simple.DB(), cnd)
}

func (s *bannerService) Count(cnd *simple.SqlCnd) int {
	return repositories.BannerRepository.Count(simple.DB(), cnd)
}

func (s *bannerService) Create(t *model.Banner) error {
	return repositories.BannerRepository.Create(simple.DB(), t)
}

func (s *bannerService) Update(t *model.Banner) error {
	return repositories.BannerRepository.Update(simple.DB(), t)
}

func (s *bannerService) Updates(id int64, columns map[string]interface{}) error {
	return repositories.BannerRepository.Updates(simple.DB(), id, columns)
}

func (s *bannerService) UpdateColumn(id int64, name string, value interface{}) error {
	return repositories.BannerRepository.UpdateColumn(simple.DB(), id, name, value)
}

func (s *bannerService) Delete(id int64) {
	repositories.BannerRepository.Delete(simple.DB(), id)
}
