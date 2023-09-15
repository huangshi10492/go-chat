package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-chat/app/friend/api/internal/logic"
	"go-chat/app/friend/api/internal/svc"
)

func GetFriendRequestListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetFriendRequestListLogic(r.Context(), svcCtx)
		resp, err := l.GetFriendRequestList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
