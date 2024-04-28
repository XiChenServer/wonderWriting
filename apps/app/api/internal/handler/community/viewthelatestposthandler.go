package community

import (
	"net/http"

	"calligraphy/apps/app/api/internal/logic/community"
	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"api_v2/common/response"
)

func ViewTheLatestPostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ViewTheLatestPostRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := community.NewViewTheLatestPostLogic(r.Context(), svcCtx)
		resp, err := l.ViewTheLatestPost(&req)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		response.Response(r, w, resp, err)

	}
}
