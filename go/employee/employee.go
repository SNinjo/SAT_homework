package employee

import (
	"fmt"
	"multithreading/meat"
	"sync"
	"time"
)

type Employee struct {
	Id string
}

type Employees struct {
	All []Employee
}

func getCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (employee Employee) process(channelMeat chan meat.Meat, waitGroup *sync.WaitGroup) {
	for meat := range channelMeat {
		fmt.Printf("%s 在 %s 取得%s\n", employee.Id, getCurrentTime(), meat.GetName())
		time.Sleep(time.Duration(meat.GetProcessingSeconds()) * time.Second)
		fmt.Printf("%s 在 %s 處理完%s\n", employee.Id, getCurrentTime(), meat.GetName())
		waitGroup.Done()
	}
}

func (employees Employees) Process(meatList []meat.Meat) {
	channelMeat := make(chan meat.Meat, len(meatList))
	waitGroup := new(sync.WaitGroup)
	for _, meat := range meatList {
		waitGroup.Add(1)
		channelMeat <- meat
	}
	for _, employee := range employees.All {
		go employee.process(channelMeat, waitGroup)
	}
	waitGroup.Wait()
	close(channelMeat)
}
