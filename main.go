package main

import (
	"fmt"
	"gocn/db"
	"gocn/dingding"
	"gocn/splider"
	"net/http"
	"time"
)

func rss(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, dingding.Rss)
}

func main() {
	go db.Run()
	go dingding.Send()

	splider.Run()

	// 当使用定时任务启动时，使用这里，等待输入写入文件完成
	time.Sleep(time.Minute * 1)

	fmt.Println("serve 8080")
	http.Handle("/rss/", http.HandlerFunc(rss))
	_ = http.ListenAndServe(":8080", nil)
}
