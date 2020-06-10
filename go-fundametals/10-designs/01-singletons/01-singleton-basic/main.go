package main

import (
	"fmt"

	"go-learning.com/singleton/whsgt"
)

func main() {
	addGoodsToWareHouse()
	addGoodsToWareHouse()
	addGoodsToWareHouse()
	releaseGoodsFromWareHouse()
	printCurrentConditionOfWareHouse()
}

func addGoodsToWareHouse() {
	capacity, err := whsgt.GetInstance().AddGoods()
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("Goods is safe in warehouse, current warehouse maxCapacity %v\n", capacity)
}

func releaseGoodsFromWareHouse() {
	capacity, err := whsgt.GetInstance().ReleaseGoods()
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("Goods is released from warehouse, current warehouse maxCapacity %v\n", capacity)
}

func printCurrentConditionOfWareHouse() {
	currentCapacity, maxCapacity := whsgt.GetInstance().GetCurrentCondition()
	fmt.Printf("Current capacity %v, max capacity %v\n", currentCapacity, maxCapacity)
}
