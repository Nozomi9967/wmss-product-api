// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"WMSS/product/api/internal/model"
	"context"
	"database/sql"

	"WMSS/product/api/internal/svc"
	"WMSS/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新产品信息
func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProductLogic) UpdateProduct(req *types.UpdateProductReq) (resp *types.Response, err error) {
	// 1. 参数校验
	if req.ProductId == "" {
		return &types.Response{
			Code: 400,
			Msg:  "产品ID不能为空",
			Data: nil,
		}, nil
	}

	// 2. 查询产品是否存在
	existProduct, err := l.svcCtx.ProductModel.FindOne(l.ctx, req.ProductId)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.Response{
				Code: 404,
				Msg:  "产品不存在",
				Data: nil,
			}, nil
		}
		l.Logger.Errorf("UpdateProduct FindOne err:%v", err)
		return &types.Response{
			Code: 500,
			Msg:  "查询产品失败",
			Data: nil,
		}, nil
	}

	// 3. 构建更新数据（只更新非空字段）
	updateData := &model.ProductInfo{
		ProductId:         req.ProductId,
		ProductName:       existProduct.ProductName,
		ProductType:       existProduct.ProductType,
		ProductSubType:    existProduct.ProductSubType,
		RiskLevel:         existProduct.RiskLevel,
		ProductStatus:     existProduct.ProductStatus,
		Manager:           existProduct.Manager,
		Custodian:         existProduct.Custodian,
		PurchaseFeeRate:   existProduct.PurchaseFeeRate,
		RedemptionFeeRule: existProduct.RedemptionFeeRule,
		Description:       existProduct.Description,
		CreateBy:          existProduct.CreateBy,
	}

	// 更新传入的非空字段
	if req.ProductName != "" {
		updateData.ProductName = req.ProductName
	}
	if req.ProductType != "" {
		updateData.ProductType = req.ProductType
	}
	if req.ProductSubType != "" {
		updateData.ProductSubType = sql.NullString{
			String: req.ProductSubType,
			Valid:  true,
		}
	}
	if req.RiskLevel != "" {
		updateData.RiskLevel = req.RiskLevel
	}
	if req.ProductStatus != "" {
		updateData.ProductStatus = req.ProductStatus
	}
	if req.Manager != "" {
		updateData.Manager = req.Manager
	}
	if req.Custodian != "" {
		updateData.Custodian = req.Custodian
	}
	if req.PurchaseFeeRate > 0 {
		updateData.PurchaseFeeRate = sql.NullFloat64{
			Float64: float64(req.PurchaseFeeRate),
			Valid:   true,
		}
	}
	if req.RedemptionFeeRule != "" {
		updateData.RedemptionFeeRule = sql.NullString{
			String: req.RedemptionFeeRule,
			Valid:  true,
		}
	}
	if req.Description != "" {
		updateData.Description = sql.NullString{
			String: req.Description,
			Valid:  true,
		}
	}

	// 4. 执行更新
	err = l.svcCtx.ProductModel.Update(l.ctx, updateData)
	if err != nil {
		l.Logger.Errorf("UpdateProduct Update err:%v", err)
		return &types.Response{
			Code: 500,
			Msg:  "更新产品失败",
			Data: nil,
		}, nil
	}

	// 5. 返回成功
	return &types.Response{
		Code: 0,
		Msg:  "更新成功",
		Data: nil,
	}, nil
}
