// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"WMSS/product/api/internal/model"
	"context"

	"WMSS/product/api/internal/svc"
	"WMSS/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取产品详情
func NewGetProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductDetailLogic {
	return &GetProductDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductDetailLogic) GetProductDetail(req *types.GetProductDetailReq) (resp *types.Response, err error) {
	productId := req.ProductId

	productInfo, err := l.svcCtx.ProductModel.FindOne(l.ctx, productId)
	if err != nil {
		if err == model.ErrNotFound {
			l.Logger.Infof("Product not found: %s", productId)
			return &types.Response{
				Code: 404,
				Msg:  "产品不存在",
				Data: nil,
			}, nil
		}

		l.Logger.Errorf("GetProductDetailLogic FindOne err:%v", err)
		return nil, err
	}

	// 转换为前端需要的格式
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
		ProductSubType:    productInfo.ProductSubType.String, // 如果为 null 则为空字符串
		PurchaseFeeRate:   float32(productInfo.PurchaseFeeRate.Float64),
		RedemptionFeeRule: productInfo.RedemptionFeeRule.String,
		Description:       productInfo.Description.String,
		CreateTime:        productInfo.CreateTime.Time.Format("2006-01-02 15:04:05"),
		UpdateTime:        productInfo.UpdateTime.Time.Format("2006-01-02 15:04:05"),
	}

	return &types.Response{
		Code: 0, // 成功使用 0
		Msg:  "success",
		Data: product,
	}, nil
}
