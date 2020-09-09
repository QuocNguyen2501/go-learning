package whsgt

import (
	"fmt"
	"sync"
)

type WareHouse interface {
	AddGoods() (int, error)
	ReleaseGoods() (int, error)
	GetCurrentCondition() (currentCapacity, maxCapacity int)
}

type warehouseImpl struct {
	maxCapacity, currentCapacity int
}

func (w *warehouseImpl) AddGoods() (int, error) {
	if w.maxCapacity == w.currentCapacity {
		return w.maxCapacity, fmt.Errorf("Capacity of warehouse is full")
	}
	w.currentCapacity++
	return w.currentCapacity, nil
}

func (w *warehouseImpl) ReleaseGoods() (int, error) {
	if w.currentCapacity == 0 {
		return w.currentCapacity, fmt.Errorf("CurrentCapacity of warehouse is 0")
	}
	w.currentCapacity--
	return w.currentCapacity, nil
}

func (w *warehouseImpl) GetCurrentCondition() (currentCapacity, maxCapacity int) {
	return w.currentCapacity, w.maxCapacity
}

var instance *warehouseImpl
var once sync.Once

func GetInstance() WareHouse {
	if instance == nil {
		once.Do(func() {
			if instance == nil {
				instance = &warehouseImpl{
					maxCapacity: 10,
				}
			}
		})
	}
	return instance
}
