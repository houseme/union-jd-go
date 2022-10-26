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
	"github.com/houseme/union-jd-go/pkg"
	"github.com/houseme/union-jd-go/pkg/handler"
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
func (s *Goods) QueryGoods(req *entity.UnionOpenGoodsJingFenQueryRequest) (res *entity.UnionOpenGoodsJingFenQueryResponseTopLevel, err error) {
	ctx, span := gtrace.NewSpan(s.ctx, "tracing-union-jd-goods-QueryGoods")
	defer span.End()

	s.config.Logger().Info(ctx, "Query Goods list start params:", req)
	if req == nil {
		err = gerror.New("query Goods list request required")
		return
	}
	var (
		util    = pkg.NewUtil(s.config.Logger())
		request = handler.NewUnionRequest(ctx, s.config, handler.WithMethod(config.UnionOpenGoodsJingFenQuery), handler.WithUnionOpenGoodsJingFenQueryRequest(req))
		resp    *handler.UnionResponse
	)

	if resp, err = util.Handler(ctx, request); err != nil {
		err = gerror.New("query category request failed")
		return
	}
	if resp == nil {
		err = gerror.New("query goods list response is nil")
		return
	}

	if resp.UnionOpenGoodsJingFenQueryResponseTopLevel == nil {
		err = gerror.New("query goods list response is nil")
		return
	}
	s.config.Logger().Debug(ctx, "query goods list response:", resp)
	res = resp.UnionOpenGoodsJingFenQueryResponseTopLevel
	s.config.Logger().Debug(ctx, "query goods list response:", res)
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

// QueryGoodsDetail goods detail query
// 京粉商品详情查询
func (s *Goods) QueryGoodsDetail(req *entity.UnionOpenGoodsBigFieldQueryRequest) (res *entity.UnionOpenGoodsBigFieldQueryResponseTopLevel, err error) {
	ctx, span := gtrace.NewSpan(s.ctx, "tracing-union-jd-goods-QueryGoodsDetail")
	defer span.End()

	s.config.Logger().Info(ctx, "Query Goods detail start params:", req)
	if req == nil {
		err = gerror.New("query Goods detail request required")
		return
	}
	var (
		util    = pkg.NewUtil(s.config.Logger())
		request = handler.NewUnionRequest(ctx, s.config, handler.WithMethod(config.UnionOpenGoodsBigFieldQuery), handler.WithUnionOpenGoodsBigFieldQueryRequest(req))
		resp    *handler.UnionResponse
	)

	if resp, err = util.Handler(ctx, request); err != nil {
		err = gerror.New("query category request failed")
		return
	}
	if resp == nil {
		err = gerror.New("query goods list response is nil")
		return
	}

	if resp.UnionOpenGoodsBigFieldQueryResponseTopLevel == nil {
		err = gerror.New("query goods list response is nil")
		return
	}
	s.config.Logger().Debug(ctx, "query goods list response:", resp)
	res = resp.UnionOpenGoodsBigFieldQueryResponseTopLevel
	s.config.Logger().Debug(ctx, "query goods list response:", res)
	return
}
