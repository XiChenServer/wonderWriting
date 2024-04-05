package user

import (
	"calligraphy/common/response"
	"net/http"

	"calligraphy/apps/app/api/internal/logic/user"
	"calligraphy/apps/app/api/internal/svc"
)

func UserModAvatarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserModAvatarLogic(r.Context(), svcCtx)
		resp, err := l.UserModAvatar(r)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
		response.HttpResult(r, w, resp, err)

	}
}
