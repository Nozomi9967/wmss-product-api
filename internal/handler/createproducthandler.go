// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"WMSS/product/api/internal/logic"
	"net/http"

	"WMSS/product/api/internal/svc"
	"WMSS/product/api/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var validate = validator.New()

// 创建新产品
// 正确示例：解析并校验
func CreateProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateProductReq
		// 解析参数并执行 binding 标签校验
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err) // 校验失败时返回错误
			return
		}
		// 校验参数
		if err := validate.Struct(&req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 执行业务逻辑
		l := logic.NewCreateProductLogic(r.Context(), svcCtx)
		resp, err := l.CreateProduct(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
