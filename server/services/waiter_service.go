package services

import (
	"bbs-go/model"
	"bbs-go/repositories"
	"github.com/mlogclub/simple"
)

var WaiterService = newWaiterService()

func newWaiterService() *waiterService {
	return &waiterService{}
}

type waiterService struct {
}

func (s *waiterService) Get(id int64) *model.Waiter {
	return repositories.WaiterRepository.Get(simple.DB(), id)
}

func (s *waiterService) Take(where ...interface{}) *model.Waiter {
	return repositories.WaiterRepository.Take(simple.DB(), where...)
}

func (s *waiterService) Find(cnd *simple.SqlCnd) []model.Waiter {
	return repositories.WaiterRepository.Find(simple.DB(), cnd)
}

func (s *waiterService) FindOne(cnd *simple.SqlCnd) *model.Waiter {
	return repositories.WaiterRepository.FindOne(simple.DB(), cnd)
}

func (s *waiterService) FindPageByParams(params *simple.QueryParams) (list []model.Waiter, paging *simple.Paging) {
	return repositories.WaiterRepository.FindPageByParams(simple.DB(), params)
}

func (s *waiterService) FindPageByCnd(cnd *simple.SqlCnd) (list []model.Waiter, paging *simple.Paging) {
	return repositories.WaiterRepository.FindPageByCnd(simple.DB(), cnd)
}

func (s *waiterService) Count(cnd *simple.SqlCnd) int {
	return repositories.WaiterRepository.Count(simple.DB(), cnd)
}

func (s *waiterService) Create(t *model.Waiter) error {
	return repositories.WaiterRepository.Create(simple.DB(), t)
}

func (s *waiterService) Update(t *model.Waiter) error {
	return repositories.WaiterRepository.Update(simple.DB(), t)
}

func (s *waiterService) Updates(id int64, columns map[string]interface{}) error {
	return repositories.WaiterRepository.Updates(simple.DB(), id, columns)
}

func (s *waiterService) UpdateColumn(id int64, name string, value interface{}) error {
	return repositories.WaiterRepository.UpdateColumn(simple.DB(), id, name, value)
}

func (s *waiterService) Delete(id int64) {
	repositories.WaiterRepository.Delete(simple.DB(), id)
}
