package handler

import (
	"calligraphy/common/response"
	"net/http"

	"calligraphy/apps/app/api/internal/logic"
	"calligraphy/apps/app/api/internal/svc"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		response.HttpResult(r, w, resp, err)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		//response.Response(r, w, resp, err)

	}
}
