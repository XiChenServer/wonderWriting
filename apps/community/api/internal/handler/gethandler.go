package handler

import (
	"net/http"

	"api_v2/common/response"
	"calligraphy/apps/community/api/internal/logic"
	"calligraphy/apps/community/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetLogic(r.Context(), svcCtx)
		resp, err := l.Get()
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		response.Response(r, w, resp, err)

	}
}
