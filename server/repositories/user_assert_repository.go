package repositories

import (
	"bbs-go/model"
	"github.com/jinzhu/gorm"
	"github.com/mlogclub/simple"
)

var UserAssertRepository = newUserAssertRepository()

func newUserAssertRepository() *userAssertRepository {
	return &userAssertRepository{}
}

type userAssertRepository struct {
}

func (r *userAssertRepository) Get(db *gorm.DB, id int64) *model.UserAssert {
	ret := &model.UserAssert{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userAssertRepository) Take(db *gorm.DB, where ...interface{}) *model.UserAssert {
	ret := &model.UserAssert{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userAssertRepository) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.UserAssert) {
	cnd.Find(db, &list)
	return
}

func (r *userAssertRepository) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.UserAssert {
	ret := &model.UserAssert{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *userAssertRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.UserAssert, paging *simple.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *userAssertRepository) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.UserAssert, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.UserAssert{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *userAssertRepository) Count(db *gorm.DB, cnd *simple.SqlCnd) int {
	return cnd.Count(db, &model.UserAssert{})
}

func (r *userAssertRepository) Create(db *gorm.DB, t *model.UserAssert) (err error) {
	err = db.Create(t).Error
	return
}

func (r *userAssertRepository) Update(db *gorm.DB, t *model.UserAssert) (err error) {
	err = db.Save(t).Error
	return
}

func (r *userAssertRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.UserAssert{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *userAssertRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.UserAssert{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *userAssertRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.UserAssert{}, "id = ?", id)
}
