package main

import (
	"awesomeProject1/config"
	"awesomeProject1/lib/logger"
	"fmt"
	"net/http"
)

func main() {
	config.Ip = "0.0.0.0"

	http.HandleFunc("/", hello)
	server := &http.Server{
		Addr: ":8888",
	}
	fmt.Println("server startup...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello 路由器没有路!"))
	logger.Logger.Info(" 这是一个测试")
}
