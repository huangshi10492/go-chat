package handler

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-chat/app/connect/internal/svc"
	"go-chat/pkg/jwt"
	"net/http"
)

func WsConnectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("ws connect")
		type Req struct {
			Token string `form:"token"`
		}
		var req Req
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error(err)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		uid, err := jwt.GetId(svcCtx.Config.Auth.AccessSecret, req.Token)
		if err != nil {
			logx.Error(err)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		svcCtx.WebSocket.NewClient(w, r, uid)
	}
}
