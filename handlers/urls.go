package handlers

import (
	"net/http"
	"bigbiy_web/handlers/bigbiy_www"
)

func MyUrls() {
	// 显示所有信息
	http.HandleFunc("/", bigbiy_www.Show_all_message)
	http.HandleFunc("/detail", bigbiy_www.Go_to_article_detail)
}
