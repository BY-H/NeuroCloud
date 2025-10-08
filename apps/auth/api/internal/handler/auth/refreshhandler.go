package auth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"neuro-cloud/apps/auth/api/internal/logic/auth"
	"neuro-cloud/apps/auth/api/internal/svc"
	"neuro-cloud/apps/auth/api/internal/types"
)

// 刷新 Token
func RefreshHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RefreshReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewRefreshLogic(r.Context(), svcCtx)
		resp, err := l.Refresh(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
