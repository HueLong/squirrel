package helper

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func GetImageLink(src string) string {
	if src == "" {
		return src
	}
	//统一小写处理
	src = strings.ToLower(src)
	//处理url
	urlParse, err := url.Parse(src)
	if err != nil {
		fmt.Println(fmt.Sprintf("GetImageLink报错：%s", err.Error()))
		return src
	}
	queryString := urlParse.RawQuery
	path := urlParse.Path
	//清除后缀
	reg, _ := regexp.Compile("![0-9a-zA-Z_!]+")
	path = reg.ReplaceAllString(path, "")
	if strings.Index(path, "/") != 0 {
		path = "/" + path
	}
	//判断是否已经有HTTP/HTTPS前缀
	path = os.Getenv("CDN_URL") + path
	if len(queryString) > 0 {
		path = path + "?" + queryString
	}
	return path
}
