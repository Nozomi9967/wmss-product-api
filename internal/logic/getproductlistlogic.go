// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"database/sql"

	"WMSS/product/api/internal/svc"
	"WMSS/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页查询产品列表
func NewGetProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductListLogic {
	return &GetProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductListLogic) GetProductList(req *types.ProductListReq) (resp *types.Response, err error) {
	// 1. 参数校验和默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	if req.PageSize > 100 {
		req.PageSize = 100
	}

	// 2. 调用 Model 层查询
	products, total, err := l.svcCtx.ProductModel.FindList(l.ctx, req)
	if err != nil {
		l.Logger.Errorf("GetProductList FindList err:%v", err)
		return &types.Response{
			Code: 500,
			Msg:  "查询失败",
			Data: nil,
		}, nil
	}

	// 3. 转换数据
	list := make([]*types.ProductInfo, 0, len(products))
	for _, productInfo := range products {
		product := &types.ProductInfo{
			ProductId:     productInfo.ProductId,
			ProductName:   productInfo.ProductName,
			ProductType:   productInfo.ProductType,
			RiskLevel:     productInfo.RiskLevel,
			ProductStatus: productInfo.ProductStatus,
			Manager:       productInfo.Manager,
			Custodian:     productInfo.Custodian,
			CreateBy:      int64(productInfo.CreateBy),

			// 处理可空字段
			ProductSubType:    productInfo.ProductSubType.String,
			PurchaseFeeRate:   float32(productInfo.PurchaseFeeRate.Float64),
			RedemptionFeeRule: productInfo.RedemptionFeeRule.String,
			Description:       productInfo.Description.String,
			CreateTime:        formatNullTime(productInfo.CreateTime),
			UpdateTime:        formatNullTime(productInfo.UpdateTime),
		}
		list = append(list, product)
	}

	// 4. 返回结果
	return &types.Response{
		Code: 0,
		Msg:  "success",
		Data: map[string]interface{}{
			"total":       total,
			"list":        list,
			"page":        req.Page,
			"page_size":   req.PageSize,
			"total_pages": (total + int64(req.PageSize) - 1) / int64(req.PageSize),
		},
	}, nil
}

// 辅助函数：格式化 NullTime
func formatNullTime(nt sql.NullTime) string {
	if nt.Valid {
		return nt.Time.Format("2006-01-02 15:04:05")
	}
	return ""
}
