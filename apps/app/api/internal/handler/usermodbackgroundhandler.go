package handler

import (
	"calligraphy/common/response"
	"net/http"

	"calligraphy/apps/app/api/internal/logic"
	"calligraphy/apps/app/api/internal/svc"
)

func UserModBackgroundHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserModBackgroundLogic(r.Context(), svcCtx)
		resp, err := l.UserModBackground(r)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		response.HttpResult(r, w, resp, err)

	}
}
