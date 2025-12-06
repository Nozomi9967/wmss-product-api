// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"WMSS/product/api/internal/types"
	"net/http"

	"WMSS/product/api/internal/logic"
	"WMSS/product/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 删除产品（逻辑删除）
func DeleteProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteProductReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewDeleteProductLogic(r.Context(), svcCtx)
		resp, err := l.DeleteProduct(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
