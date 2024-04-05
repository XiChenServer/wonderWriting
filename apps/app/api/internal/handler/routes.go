// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	community "calligraphy/apps/app/api/internal/handler/community"
	user "calligraphy/apps/app/api/internal/handler/user"
	"calligraphy/apps/app/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/basic/getemailverification",
				Handler: user.GetEmailVerificationHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: user.UserRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.UserLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/forgetpwd",
				Handler: user.UserForgetPwdHandler(serverCtx),
			},
		},
		rest.WithPrefix("/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/userinfo",
				Handler: user.UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/modpwd",
				Handler: user.UserModPwdHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/modavatar",
				Handler: user.UserModAvatarHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/modbackground",
				Handler: user.UserModBackgroundHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/modinfo",
				Handler: user.UserModInfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/look/allposts",
				Handler: community.LookAllPostsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/community"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create/post",
				Handler: community.UsercretePostHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/community"),
	)
}
