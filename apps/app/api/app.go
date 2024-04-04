package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"calligraphy/apps/app/api/internal/config"
	"calligraphy/apps/app/api/internal/handler"
	"calligraphy/apps/app/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
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

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

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
