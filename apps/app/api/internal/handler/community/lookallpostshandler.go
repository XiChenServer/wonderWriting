package community

import (
	"calligraphy/common/response"
	"net/http"

	"calligraphy/apps/app/api/internal/logic/community"
	"calligraphy/apps/app/api/internal/svc"
)

func LookAllPostsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := community.NewLookAllPostsLogic(r.Context(), svcCtx)
		resp, err := l.LookAllPosts(r)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		response.HttpResult(r, w, resp, err)

	}
}
