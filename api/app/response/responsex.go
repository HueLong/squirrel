package response

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

const (
	codeSuccess = 0
	codeFail    = 1

	httpCodeSuccess     = 200
	httpCodeErrorParams = 400

	httpMsgSuccess = "请求成功"
)

type (
	Response struct {
		ErrCode int         `json:"errcode"`
		ErrMsg  string      `json:"errmsg"`
		Data    interface{} `json:"data"`
	}

	ResponseDataList struct {
		ErrCode int          `json:"errcode"`
		ErrMsg  string       `json:"errmsg"`
		Data    ResponseList `json:"data"`
	}

	ResponseDataListWithCount struct {
		ErrCode int                   `json:"errcode"`
		ErrMsg  string                `json:"errmsg"`
		Data    ResponseListWithCount `json:"data"`
	}

	ResponseList struct {
		List interface{} `json:"list"`
	}

	ResponseListWithCount struct {
		Count int         `json:"count"`
		List  interface{} `json:"list"`
	}
)

func initResponse(ec int, es string) Response {
	return Response{
		ErrCode: ec,
		ErrMsg:  es,
	}
}

func initListResponse(ec int, es string) ResponseDataList {
	return ResponseDataList{
		ErrCode: ec,
		ErrMsg:  es,
	}
}

func initListCountResponse(ec int, es string) ResponseDataListWithCount {
	return ResponseDataListWithCount{
		ErrCode: ec,
		ErrMsg:  es,
	}
}

// =================== 常规返回 =====================
func HttpSuccess(ctx *gin.Context, data interface{}) {
	r := initResponse(codeSuccess, "请求成功")
	if data == nil {
		data = struct{}{}
	}
	r.Data = data
	ctx.Render(httpCodeSuccess, render.JSON{Data: r})
}

func HttpFail(ctx *gin.Context, data interface{}, err error) {
	r := initResponse(codeFail, err.Error())
	r.Data = data
	ctx.Render(httpCodeSuccess, render.JSON{Data: r})
}

// =================== 列表返回 =====================
func HttpSuccessList(ctx *gin.Context, data interface{}) {
	r := initListResponse(codeSuccess, "请求成功")
	if data == nil {
		data = [][]int8{}
	}
	r.Data.List = data
	ctx.Render(httpCodeSuccess, render.JSON{Data: r})
}

func HttpSuccessListWithCount(ctx *gin.Context, data interface{}, count int) {
	r := initListCountResponse(codeSuccess, "请求成功")
	if data == nil {
		data = [][]int8{}
	}
	r.Data.List = data
	r.Data.Count = count
	ctx.Render(httpCodeSuccess, render.JSON{Data: r})
}

func HttpFailList(ctx *gin.Context, data interface{}, err error) {
	r := initListResponse(codeFail, err.Error())
	if data == nil {
		data = [][]int8{}
	}
	r.Data.List = data
	ctx.Render(httpCodeSuccess, render.JSON{Data: r})
}

func HttpFailListWithCount(ctx *gin.Context, data interface{}, count int, err error) {
	r := initListCountResponse(codeFail, err.Error())
	if data == nil {
		data = [][]int8{}
	}
	r.Data.List = data
	r.Data.Count = count
	ctx.Render(httpCodeSuccess, render.JSON{Data: r})
}
