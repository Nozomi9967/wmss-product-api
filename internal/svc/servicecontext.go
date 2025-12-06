// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"WMSS/product/api/internal/config"
	"WMSS/product/api/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	ProductModel model.ProductInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)

	return &ServiceContext{
		Config:       c,
		ProductModel: model.NewProductInfoModel(conn),
	}
}
