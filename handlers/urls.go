package handlers

import (
	"bigbiy_web/handlers/bigbiy_www"
	"net/http"
)

func MyUrls() {
	http.HandleFunc("/", bigbiy_www.Index_v2)
	http.HandleFunc("/novel_v2", bigbiy_www.Nvl_v2)
}
