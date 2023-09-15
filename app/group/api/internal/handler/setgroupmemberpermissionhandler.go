package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-chat/app/group/api/internal/logic"
	"go-chat/app/group/api/internal/svc"
	"go-chat/app/group/api/internal/types"
)

func SetGroupMemberPermissionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetGroupMemberPermissionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSetGroupMemberPermissionLogic(r.Context(), svcCtx)
		err := l.SetGroupMemberPermission(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
