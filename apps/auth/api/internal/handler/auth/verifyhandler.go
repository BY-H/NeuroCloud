package auth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"neuro-cloud/apps/auth/api/internal/logic/auth"
	"neuro-cloud/apps/auth/api/internal/svc"
	"neuro-cloud/apps/auth/api/internal/types"
)

// 验证 Token
func VerifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerifyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewVerifyLogic(r.Context(), svcCtx)
		resp, err := l.Verify(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
