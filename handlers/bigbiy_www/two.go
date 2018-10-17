package bigbiy_www

import (
	"bigbiy_web/config"
	"bigbiy_web/util"
	"fmt"
	"net/http"
)

func Index_v2(w http.ResponseWriter, r *http.Request) {
	var data = make(map[string]interface{})
	//获取redis中的图片
	img := util.Get_redis("biying_img")
	data["img"] = img
	fmt.Println(img)
	template_path := config.Template_path + "index_v2.html"
	util.Render_template(w, template_path, data)
}
