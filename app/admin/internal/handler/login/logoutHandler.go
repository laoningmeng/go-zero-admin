package login

import (
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/logic"
	"github.com/laoningmeng/go-zero-admin/app/admin/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Logout(r.Header.Get("authorization"))
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
