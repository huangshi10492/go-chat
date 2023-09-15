package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-chat/app/friend/api/internal/logic"
	"go-chat/app/friend/api/internal/svc"
	"go-chat/app/friend/api/internal/types"
)

func ModifyFriendRemarkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ModifyFriendRemarkReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewModifyFriendRemarkLogic(r.Context(), svcCtx)
		err := l.ModifyFriendRemark(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
