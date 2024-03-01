package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // 设置路由
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil) // 监听端口
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Client!") // 发送响应内容
}
