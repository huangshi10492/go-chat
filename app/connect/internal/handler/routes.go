package handler

import (
	"github.com/zeromicro/go-zero/rest"
	"go-chat/app/connect/internal/svc"
	"net/http"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/connect/ws",
				Handler: WsConnectHandler(serverCtx),
			},
		},
	)
}
