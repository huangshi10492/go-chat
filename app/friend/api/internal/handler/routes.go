// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-chat/app/friend/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/friend/send",
				Handler: SendFriendHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/friend/agree",
				Handler: AgreeFriendHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/friend/modifyRemark",
				Handler: ModifyFriendRemarkHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/friend/delete",
				Handler: DeleteFriendHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/friend/get",
				Handler: GetFriendListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/friend/getRequset",
				Handler: GetFriendRequestListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
