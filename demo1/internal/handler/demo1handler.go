package handler

import (
	"net/http"

	"demo1/internal/logic"
	"demo1/internal/svc"
	"demo1/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Demo1Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDemo1Logic(r.Context(), svcCtx)
		resp, err := l.Demo1(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
