package bigbiy_www

import (
	"bigbiy_web/config"
	"bigbiy_web/models"
	"bigbiy_web/util"
	"net/http"
	"strconv"
)

//去往主页的方法
func Index_v2(w http.ResponseWriter, r *http.Request) {
	var data = make(map[string]interface{})
	//获取redis中的图片
	img := util.Get_redis("biying_img")
	data["img"] = img
	template_path := config.Template_path + "index_v2.html"
	util.Render_template(w, template_path, data)
}

//去往nvl的方法
func Nvl_v2(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1024 * 3)
	last_book_id := util.Get_argument(r, "last_book_id")
	var data = make(map[string]interface{})
	//获取redis中的图片
	img := util.Get_redis("biying_img")
	data["img"] = img
	//获取redis中的热搜词
	hot_words := util.Get_redis("hot_words")
	data["hot_words"] = hot_words
	// 获取推荐书本
	banners := Get_banner()
	data["banner"] = banners
	//查询分页展示的书籍
	var books []models.Banner_novel
	sql_str := "select id,name,book_img,author from book where id >? and has_chapter=1 limit 20;"
	rows, err := util.DB.Query(sql_str, last_book_id)
	util.CheckErr(err)
	var one_banner models.Banner_novel
	for rows.Next() {
		rows.Scan(&one_banner.Book_id, &one_banner.Name, &one_banner.Img, &one_banner.Author)
		Get_desc_by_book_id(&one_banner)
		books = append(books, one_banner)
	}
	data["books"] = books
	template_path := config.Template_path + "nvl_v2.html"
	util.Render_template(w, template_path, data)
}

func Get_desc_by_book_id(one_hot_novel *models.Banner_novel) {
	book_id := one_hot_novel.Book_id
	sql_str := "select chapter_text from chapter where book_id=" + strconv.Itoa(book_id) + " limit 1;"
	rows, err := util.DB.Query(sql_str)
	defer rows.Close()
	util.CheckErr(err)
	var text string
	for rows.Next() {
		rows.Scan(&text)
	}
	var text_list []string
	util.Json_to_object(text, &text_list)
	var desc string
	var i = 0
	for _, sentence := range text_list {
		desc = desc + sentence
		i = i + 1
		if i > 2 {
			break
		}
	}
	one_hot_novel.Desc = desc + "......"
}

// 根据书本id查询数据库中的书本信息
func Get_banner_by_id(novel_id int) models.Banner_novel {
	sql_str := "select id,name,book_img,author from book where id=" + strconv.Itoa(novel_id) + ";"
	rows, err := util.DB.Query(sql_str)
	defer rows.Close()
	util.CheckErr(err)
	var one_banner models.Banner_novel
	for rows.Next() {
		rows.Scan(&one_banner.Book_id, &one_banner.Name, &one_banner.Img, &one_banner.Author)
		Get_desc_by_book_id(&one_banner)
	}
	return one_banner
}

//获取banner数据的方法
func Get_banner() []models.Banner_novel {
	banner_id_list := util.Get_banner_novel_id()
	var banners []models.Banner_novel
	for _, novel_id := range banner_id_list {
		one_banner := Get_banner_by_id(novel_id)
		banners = append(banners, one_banner)
	}
	return banners
}
