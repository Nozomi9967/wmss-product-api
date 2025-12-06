package logic

import (
	"WMSS/product/api/internal/model"
	"WMSS/product/api/internal/svc"
	"WMSS/product/api/internal/types"
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateProductLogic) CreateProduct(req *types.CreateProductReq) (resp types.Response, err error) {
	// 生成唯一ID，可用 UUID
	productId := strings.ReplaceAll(uuid.New().String(), "-", "")

	// 构建要插入的结构体
	newProduct := &model.ProductInfo{
		ProductId:         productId,
		ProductName:       req.ProductName,
		ProductType:       req.ProductType,
		ProductSubType:    sql.NullString{String: req.ProductSubType, Valid: req.ProductSubType != ""},
		RiskLevel:         req.RiskLevel,
		ProductStatus:     req.ProductStatus,
		Manager:           req.Manager,
		Custodian:         req.Custodian,
		PurchaseFeeRate:   sql.NullFloat64{Float64: float64(req.PurchaseFeeRate), Valid: req.PurchaseFeeRate != 0},
		RedemptionFeeRule: sql.NullString{String: req.RedemptionFeeRule, Valid: req.RedemptionFeeRule != ""},
		Description:       sql.NullString{String: req.Description, Valid: req.Description != ""},
		CreateBy:          uint64(1), //TODO
		CreateTime:        sql.NullTime{Time: time.Now(), Valid: true},
		UpdateTime:        sql.NullTime{Time: time.Now(), Valid: true},
	}

	// 插入数据库
	_, err = l.svcCtx.ProductModel.Insert(l.ctx, newProduct)
	if err != nil {
		l.Logger.Errorf("插入产品失败: %v", err)
		return types.Response{
			Code: 500,
			Msg:  "创建产品失败",
		}, err
	}

	return types.Response{
		Code: 200,
		Msg:  "创建产品成功",
	}, nil
}
