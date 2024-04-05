package user

import (
	"net/http"

	"api_v2/common/response"
	"calligraphy/apps/app/api/internal/logic/user"
	"calligraphy/apps/app/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserModBackgroundHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserModBackgroundLogic(r.Context(), svcCtx)
		resp, err := l.UserModBackground()
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		response.Response(r, w, resp, err)

	}
}
