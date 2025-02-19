package response

import (
	"calligraphy/common/xerr"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

type Response struct {
	Code    uint32      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// http返回
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	if err == nil {

		//成功返回
		r := &Response{
			Code:    200,
			Message: "success",
			Data:    resp,
		}
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		errcode := uint32(500)
		//errmsg := "服务器错误"
		fmt.Println(err.Error())
		causeErr := errors.Cause(err)                // err类型
		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
			//自定义CodeError

			errcode = e.GetErrCode()
			//errmsg = e.GetErrMsg()
		} else {

			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				errcode = grpcCode
				//errmsg = gstatus.Message()
			}
		}

		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
		if err.Error() == "rpc error: code = Unknown desc = already checked in today" {
			httpx.WriteJson(w, http.StatusOK, &Response{
				Code:    200,
				Message: "今日已经打开",
				Data:    resp,
			})
			return
		}
		if err.Error() == "rpc error: code = Unknown desc = user already liked this post" {
			httpx.WriteJson(w, http.StatusOK, &Response{
				Code:    200,
				Message: "已经点赞了",
				Data:    resp,
			})
			return
		}
		if err.Error() == "rpc error: code = Unknown desc = user is already followed" {
			httpx.WriteJson(w, http.StatusOK, &Response{
				Code:    200,
				Message: "已经关注了",
				Data:    resp,
			})
			return
		}

		if err.Error() == "rpc error: code = Unknown desc = post already collected by user" {
			httpx.WriteJson(w, http.StatusOK, &Response{
				Code:    200,
				Message: "已经收藏了",
				Data:    resp,
			})
			return
		}

		httpx.WriteJson(w, http.StatusBadRequest, &Response{
			Code:    errcode,
			Message: err.Error(),
			Data:    resp,
		})
	}
}
