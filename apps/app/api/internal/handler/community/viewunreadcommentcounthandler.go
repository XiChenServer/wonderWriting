package community

import (
	"calligraphy/common/response"
	"net/http"

	"calligraphy/apps/app/api/internal/logic/community"
	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ViewUnreadCommentCountHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ViewUnreadCommentCountRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := community.NewViewUnreadCommentCountLogic(r.Context(), svcCtx)
		resp, err := l.ViewUnreadCommentCount(&req)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		response.HttpResult(r, w, resp, err)

	}
}
