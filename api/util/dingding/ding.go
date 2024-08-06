package dingding

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"time"
)

type Ding struct {
	title   string
	content string
	method  string
	atPhone []string
	isAtAll bool
}

// Send 发送钉钉消息
func (d Ding) Send() error {
	var urlPath = "https://oapi.dingtalk.com/robot/send?access_token=4b4d943b0c96340ff0b773724d858f6489f54a370d366c43a926f870635356f1"
	if d.title == "" {
		return errors.New("请填写标题")
	}
	if d.content == "" {
		return errors.New("请填写内容")
	}
	jsonObj := gin.H{
		"msgtype": "text",
		"text":    map[string]interface{}{"content": d.title + "\n\n" + d.method + "\n\n" + d.content},
		"at":      gin.H{"atMobiles": d.atPhone, "isAtAll": d.isAtAll},
	}
	jsonByte, _ := json.Marshal(jsonObj)
	t, sign := d.sign()
	urlPath = fmt.Sprintf("%s&sign=%s&timestamp=%d", urlPath, sign, t)
	_, _ = http.Post(urlPath, "application/json", bytes.NewReader(jsonByte))
	return nil
}

func (d Ding) SetTitle(title string) Ding {
	d.title = os.Getenv("GIN_ENV") + ":" + title
	return d
}

func (d Ding) SetContent(content string) Ding {
	d.content = content
	d.method = getFuncName()
	return d
}

func getFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func (d Ding) At(names ...string) Ding {
	if len(names) == 0 {
		return d
	}
	var phoneMap = map[string]string{
		"dabao":     "18556970801",
		"paidaxing": "15321562967",
		"doufu":     "17326025615",
		"donghai":   "15737646132",
		"qingyang":  "18602576910",
		"daxiong":   "13018949510",
		"tufei":     "18682027992",
		"zhangyg":   "13757124196",
	}
	for _, name := range names {
		phone, _ := phoneMap[name]
		d.atPhone = append(d.atPhone, phone)
	}
	return d
}

func (d Ding) AtAll() Ding {
	d.isAtAll = true
	return d
}

func (d Ding) sign() (int64, string) {
	var t int64
	t = time.Now().Unix() * 1000
	secret := "SECff4b7a9b9550623f5739017e30bed389cf50957bdd3e5ad845b23eed36b84855"
	var key = fmt.Sprintf("%d\n%s", t, secret)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(key))
	hmacStr := h.Sum(nil)
	b := base64.StdEncoding.EncodeToString(hmacStr)
	urlStr := url.QueryEscape(b)
	return t, urlStr
}
