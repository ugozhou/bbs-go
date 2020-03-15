package repositories

import (
	"bbs-go/model"
	"github.com/jinzhu/gorm"
	"github.com/mlogclub/simple"
)

var WaiterRepository = newWaiterRepository()

func newWaiterRepository() *waiterRepository {
	return &waiterRepository{}
}

type waiterRepository struct {
}

func (r *waiterRepository) Get(db *gorm.DB, id int64) *model.Waiter {
	ret := &model.Waiter{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *waiterRepository) Take(db *gorm.DB, where ...interface{}) *model.Waiter {
	ret := &model.Waiter{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *waiterRepository) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Waiter) {
	cnd.Find(db, &list)
	return
}

func (r *waiterRepository) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.Waiter {
	ret := &model.Waiter{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *waiterRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.Waiter, paging *simple.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *waiterRepository) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Waiter, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.Waiter{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *waiterRepository) Count(db *gorm.DB, cnd *simple.SqlCnd) int {
	return cnd.Count(db, &model.Waiter{})
}

func (r *waiterRepository) Create(db *gorm.DB, t *model.Waiter) (err error) {
	err = db.Create(t).Error
	return
}

func (r *waiterRepository) Update(db *gorm.DB, t *model.Waiter) (err error) {
	err = db.Save(t).Error
	return
}

func (r *waiterRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.Waiter{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *waiterRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.Waiter{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *waiterRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.Waiter{}, "id = ?", id)
}
