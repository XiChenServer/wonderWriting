package main

import (
	"calligraphy/apps/app/api/internal/config"
	"calligraphy/apps/app/api/internal/handler"
	"calligraphy/apps/app/api/internal/svc"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

var configFile = flag.String("f", "etc/app.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(JwtUnauthorizedResult))
	//server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()
	//svctx := context.Background()

	//for _, mq := range mqs.Consumers(c, svctx, ctx) {
	//	serviceGroup.Add(mq)
	//}

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	//serviceGroup.Start()
	server.Start()

}

//
//func main() {
//	flag.Parse()
//
//	var c config.Config
//	conf.MustLoad(*configFile, &c)
//
//	server := rest.MustNewServer(c.RestConf)
//	defer server.Stop()
//
//	svcCtx := svc.NewServiceContext(c)
//	ctx := context.Background()
//	serviceGroup := service.NewServiceGroup()
//	defer serviceGroup.Stop()
//
//	for _, mq := range mqs.Consumers(c, ctx, svcCtx) {
//		serviceGroup.Add(mq)
//	}
//	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
//	server.Start()
//	serviceGroup.Start()
//}

//jwt认证失败返回给调用者

type jwtInfo struct {
}

func JwtUnauthorizedResult(w http.ResponseWriter, r *http.Request, err error) {
	//httpx.WriteJson(w, http.StatusUnauthorized, "jwt鉴权失败:"+err.Error())
	response := map[string]interface{}{
		"code":    http.StatusUnauthorized,
		"message": "jwt鉴权失败:" + err.Error(),
		"data":    jwtInfo{},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	_, _ = w.Write(jsonResponse)
}
