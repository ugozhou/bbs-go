package repositories

import (
	"bbs-go/model"
	"github.com/jinzhu/gorm"
	"github.com/mlogclub/simple"
)

var UserAssertLogRepository = newUserAssertLogRepository()

func newUserAssertLogRepository() *userAssertLogRepository {
	return &userAssertLogRepository{}
}

type userAssertLogRepository struct {
}

func (r *userAssertLogRepository) Get(db *gorm.DB, id int64) *model.UserAssertLog {
	ret := &model.UserAssertLog{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userAssertLogRepository) Take(db *gorm.DB, where ...interface{}) *model.UserAssertLog {
	ret := &model.UserAssertLog{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userAssertLogRepository) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.UserAssertLog) {
	cnd.Find(db, &list)
	return
}

func (r *userAssertLogRepository) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.UserAssertLog {
	ret := &model.UserAssertLog{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *userAssertLogRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.UserAssertLog, paging *simple.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *userAssertLogRepository) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.UserAssertLog, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.UserAssertLog{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *userAssertLogRepository) Count(db *gorm.DB, cnd *simple.SqlCnd) int {
	return cnd.Count(db, &model.UserAssertLog{})
}

func (r *userAssertLogRepository) Create(db *gorm.DB, t *model.UserAssertLog) (err error) {
	err = db.Create(t).Error
	return
}

func (r *userAssertLogRepository) Update(db *gorm.DB, t *model.UserAssertLog) (err error) {
	err = db.Save(t).Error
	return
}

func (r *userAssertLogRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.UserAssertLog{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *userAssertLogRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.UserAssertLog{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *userAssertLogRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.UserAssertLog{}, "id = ?", id)
}
