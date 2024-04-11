package community

import (
	"calligraphy/apps/app/api/internal/types"
	"calligraphy/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"calligraphy/apps/app/api/internal/logic/community"
	"calligraphy/apps/app/api/internal/svc"
)

func LookAllPostsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LookAllPostsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := community.NewLookAllPostsLogic(r.Context(), svcCtx)

		resp, err := l.LookAllPosts(&req)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		response.HttpResult(r, w, resp, err)

	}
}
