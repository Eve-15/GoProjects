package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		themeCookie, err := r.Cookie("theme")
		if err == nil {
			// 从 Cookie 中获取主题值
			theme := themeCookie.Value
			// 根据主题值应用不同的样式或逻辑
			if theme == "dark" {
				// 应用暗色主题的样式
				fmt.Fprintln(w, "Using dark theme")
			} else {
				// 应用默认主题的样式
				fmt.Fprintln(w, "Using default theme")
			}
		} else {
			// 没有找到名为 "theme" 的 Cookie
			fmt.Fprintln(w, "Using default theme")
		}
	})

	http.HandleFunc("/select-theme", func(w http.ResponseWriter, r *http.Request) {
		// 获取用户选择的主题
		theme := r.FormValue("theme")

		// 创建一个新的 Cookie
		cookie := &http.Cookie{
			Name:  "theme",
			Value: theme,
			Path:  "/",
		}

		// 将 Cookie 添加到响应的 Header 中
		http.SetCookie(w, cookie)

		// 重定向到首页
		http.Redirect(w, r, "/", http.StatusFound)
	})

	http.ListenAndServe(":8080", nil)
}
