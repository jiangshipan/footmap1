package main

import (
	"fmt"
	"footmap/app/controller"
	"footmap/app/dao"
	"footmap/app/models"
	"footmap/app/utils"
	"net/http"
	"time"
)


func main() {
	server := http.Server{
		Addr:"127.0.0.1:8080",
	}
	record := make(chan models.Record, 5)
	go getRecord(record)
	go doRecord(record)
	http.HandleFunc("/user/",controller.Verify)
	http.HandleFunc("/record/",controller.Verify)

	err := server.ListenAndServe()
	utils.CheckError(err)
}

func getRecord (in chan models.Record) {
	fmt.Println("获取任务线程开启!")
	//每五分钟向通道中加入事件
	for {
		records := dao.GetRecord()
		for i := range records {
			fmt.Println(records[i].TargetTime)
			in <- records[i]
		}
		fmt.Println("阻塞中")
		time.Sleep(5 * time.Minute)
	}
}
func doRecord (in chan models.Record) {
	fmt.Println("执行任务线程开启!")
	//执行record
	for {
		select {
		case record := <- in :
			fmt.Printf("执行%s任务了!", record.Title)
		default:
			fmt.Println("阻塞中")
			time.Sleep(30 * time.Second)
		}
	}
}