// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"github.com/Nozomi9967/wmss-product-api/internal/logic"
	"github.com/Nozomi9967/wmss-product-api/internal/svc"
	"github.com/Nozomi9967/wmss-product-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 分页查询产品列表
func GetProductListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetProductListLogic(r.Context(), svcCtx)
		resp, err := l.GetProductList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
