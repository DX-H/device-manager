package handler

import (
	"net/http"

	"finishy1995/device-manager/enhanced/http/internal/logic"
	"finishy1995/device-manager/enhanced/http/internal/svc"
	"finishy1995/device-manager/enhanced/http/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetMetadataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetMetadataRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetMetadataLogic(r.Context(), svcCtx)
		resp, err := l.GetMetadata(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
