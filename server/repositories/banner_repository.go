package repositories

import (
	"bbs-go/model"
	"github.com/jinzhu/gorm"
	"github.com/mlogclub/simple"
)

var BannerRepository = newBannerRepository()

func newBannerRepository() *bannerRepository {
	return &bannerRepository{}
}

type bannerRepository struct {
}

func (r *bannerRepository) Get(db *gorm.DB, id int64) *model.Banner {
	ret := &model.Banner{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *bannerRepository) Take(db *gorm.DB, where ...interface{}) *model.Banner {
	ret := &model.Banner{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *bannerRepository) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Banner) {
	cnd.Find(db, &list)
	return
}

func (r *bannerRepository) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.Banner {
	ret := &model.Banner{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *bannerRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.Banner, paging *simple.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *bannerRepository) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Banner, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.Banner{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *bannerRepository) Count(db *gorm.DB, cnd *simple.SqlCnd) int {
	return cnd.Count(db, &model.Banner{})
}

func (r *bannerRepository) Create(db *gorm.DB, t *model.Banner) (err error) {
	err = db.Create(t).Error
	return
}

func (r *bannerRepository) Update(db *gorm.DB, t *model.Banner) (err error) {
	err = db.Save(t).Error
	return
}

func (r *bannerRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.Banner{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *bannerRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.Banner{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *bannerRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.Banner{}, "id = ?", id)
}
