package handlers

import (
	"bigbiy_web/handlers/bigbiy_www"
	"net/http"
)

func MyUrls() {
	http.HandleFunc("/", bigbiy_www.Index_v2)
	http.HandleFunc("/detail", bigbiy_www.Go_to_article_detail)
}
