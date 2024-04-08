package group

import (
	"calligraphy/common/response"
	"net/http"

	"calligraphy/apps/app/api/internal/logic/group"
	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CheckPunchCardHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckPunchCardModelRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := group.NewCheckPunchCardLogic(r.Context(), svcCtx)
		resp, err := l.CheckPunchCard(&req)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		response.HttpResult(r, w, resp, err)

	}
}
