package main

import (
	"fmt"
	"sync"

	"go-learning.com/singleton/whsgt"
)

func main() {
	var wg *sync.WaitGroup
	wg = &sync.WaitGroup{}
	wg.Add(5)
	go addGoodsToWareHouse(wg)
	go addGoodsToWareHouse(wg)
	go addGoodsToWareHouse(wg)
	go releaseGoodsFromWareHouse(wg)
	go printCurrentConditionOfWareHouse(wg)

	wg.Wait()
}

func addGoodsToWareHouse(wg *sync.WaitGroup) {
	capacity, err := whsgt.GetInstance().AddGoods()
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("Goods is safe in warehouse, current warehouse %v\n", capacity)
	wg.Done()
}

func releaseGoodsFromWareHouse(wg *sync.WaitGroup) {
	capacity, err := whsgt.GetInstance().ReleaseGoods()
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("Goods is released from warehouse, current warehouse maxCapacity %v\n", capacity)
	wg.Done()
}

func printCurrentConditionOfWareHouse(wg *sync.WaitGroup) {
	currentCapacity, maxCapacity := whsgt.GetInstance().GetCurrentCondition()
	fmt.Printf("Current capacity %v, max capacity %v\n", currentCapacity, maxCapacity)
	wg.Done()
}
