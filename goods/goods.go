/*
 * Copyright union-jd-go Author(https://houseme.github.io/union-jd-go/). All Rights Reserved.
 *
 * This Source Code Form is subject to the terms of the MIT License.
 * If a copy of the MIT was not distributed with this file,
 * You can obtain one at https://github.com/houseme/union-jd-go.
 */

// Package goods implements the use case for JD goods.
package goods

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/gtrace"

	"github.com/houseme/union-jd-go/config"
	"github.com/houseme/union-jd-go/entity"
)

// Goods .
type Goods struct {
	ctx    context.Context
	config *config.Config
}

// NewGoods implements the use case for JD goods.
func NewGoods(ctx context.Context, c *config.Config) (*Goods, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-union-jd-goods-NewGoods")
	defer span.End()

	if c == nil {
		return nil, gerror.New("JD configuration required")
	}

	return &Goods{
		ctx:    ctx,
		config: c,
	}, nil
}

// GetConfig .
func (s *Goods) GetConfig() *config.Config {
	return s.config
}

// QueryGoods jingfen goods query
// 京粉商品查询
func (s *Goods) QueryGoods(req *entity.UnionOpenGoodsJingFenQueryRequest) (resp *entity.JDUnionOpenGoodsBigFieldQueryTopLevel, err error) {
	ctx, span := gtrace.NewSpan(s.ctx, "tracing-union-jd-goods-QueryGoods")
	defer span.End()

	s.config.Logger().Info(ctx, "Query Goods list start params:", req)
	if req == nil {
		err = gerror.New("query Goods list request required")
		return
	}

	return
}

// NewJFGoodsReq 创建商品查询请求
func (s *Goods) NewJFGoodsReq(eliteID config.EliteID, pageIndex, pageSize uint, sortName, sort, pid, fields string) *entity.JFGoodsReq {
	return &entity.JFGoodsReq{
		EliteID:   eliteID,
		PageIndex: pageIndex,
		PageSize:  pageSize,
		SortName:  sortName,
		Sort:      sort,
		Pid:       pid,
		Fields:    fields,
	}
}

// NewUnionOpenGoodsJingFenQueryRequest 创建京粉精选商品查询接口请求
func (s *Goods) NewUnionOpenGoodsJingFenQueryRequest(goodsReq *entity.JFGoodsReq) *entity.UnionOpenGoodsJingFenQueryRequest {
	return &entity.UnionOpenGoodsJingFenQueryRequest{
		GoodsReq: goodsReq,
	}
}
