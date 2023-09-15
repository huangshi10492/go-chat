// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-chat/app/message/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/message/sendMessage",
				Handler: SendMessageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/message/read",
				Handler: ReadHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/message/withdraw",
				Handler: WithdrawHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
