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

type DeleteProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除产品（逻辑删除）
func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteProductLogic) DeleteProduct(req *types.DeleteProductReq) (resp *types.Response, err error) {
	// 1. 参数校验
	if req.ProductId == "" {
		return &types.Response{
			Code: 400,
			Msg:  "产品ID不能为空",
			Data: nil,
		}, nil
	}

	// 2. 检查产品是否存在
	_, err = l.svcCtx.ProductModel.FindOne(l.ctx, req.ProductId)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.Response{
				Code: 404,
				Msg:  "产品不存在",
				Data: nil,
			}, nil
		}
		l.Logger.Errorf("DeleteProduct FindOne err:%v", err)
		return &types.Response{
			Code: 500,
			Msg:  "查询产品失败",
			Data: nil,
		}, nil
	}

	// 3. 执行逻辑删除
	err = l.svcCtx.ProductModel.DeleteLogical(l.ctx, req.ProductId)
	if err != nil {
		l.Logger.Errorf("DeleteProduct Delete err:%v", err)
		return &types.Response{
			Code: 500,
			Msg:  "删除产品失败",
			Data: nil,
		}, nil
	}

	// 4. 返回成功
	return &types.Response{
		Code: 0,
		Msg:  "删除成功",
		Data: nil,
	}, nil
}
