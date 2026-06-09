package check_time

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

type ExcelToMain struct {
}

func (t *ExcelToMain) CheckToTime() bool {

	currentTime := time.Now()
	targetTime := time.Date(2024, 4, 15, 0, 0, 0, 0, time.UTC)

	if currentTime.After(targetTime) {
		//fmt.Println("当前时间大于指定日期")
		return true
	} else {
		//fmt.Println("当前时间小于或等于指定日期")
		return false
	}
}
func (t *ExcelToMain) Hang() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	//ServerClose()
	fmt.Printf("kill process exit ------- signal:[%v]", s)
	//log4.Info("kill process exit ------- signal:[%v]", s)

}
