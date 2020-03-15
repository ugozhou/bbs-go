package repositories

import (
	"bbs-go/model"
	"github.com/jinzhu/gorm"
	"github.com/mlogclub/simple"
)

var NoticeRepository = newNoticeRepository()

func newNoticeRepository() *noticeRepository {
	return &noticeRepository{}
}

type noticeRepository struct {
}

func (r *noticeRepository) Get(db *gorm.DB, id int64) *model.Notice {
	ret := &model.Notice{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *noticeRepository) Take(db *gorm.DB, where ...interface{}) *model.Notice {
	ret := &model.Notice{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *noticeRepository) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Notice) {
	cnd.Find(db, &list)
	return
}

func (r *noticeRepository) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.Notice {
	ret := &model.Notice{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *noticeRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.Notice, paging *simple.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *noticeRepository) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Notice, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.Notice{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *noticeRepository) Count(db *gorm.DB, cnd *simple.SqlCnd) int {
	return cnd.Count(db, &model.Notice{})
}

func (r *noticeRepository) Create(db *gorm.DB, t *model.Notice) (err error) {
	err = db.Create(t).Error
	return
}

func (r *noticeRepository) Update(db *gorm.DB, t *model.Notice) (err error) {
	err = db.Save(t).Error
	return
}

func (r *noticeRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.Notice{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *noticeRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.Notice{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *noticeRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.Notice{}, "id = ?", id)
}
