package main

import (
	"math/rand"
	"multithreading/employee"
	"multithreading/meat"
)

func main() {
	meatList := []meat.Meat{}
	for time := 0; time < 10; time++ {
		meatList = append(meatList, meat.Beef{})
	}
	for time := 0; time < 7; time++ {
		meatList = append(meatList, meat.Pork{})
	}
	for time := 0; time < 5; time++ {
		meatList = append(meatList, meat.Chicken{})
	}
	// 透過打亂肉品的順序來達到隨機選取的效果
	rand.Shuffle(len(meatList), func(i, j int) {
		meatList[i], meatList[j] = meatList[j], meatList[i]
	})

	// 創建處理肉品的員工並賦予 ID
	employees := employee.Employees{
		All: []employee.Employee{{Id: "A"}, {Id: "B"}, {Id: "C"}, {Id: "D"}, {Id: "E"}},
	}
	employees.Process(meatList)
}
