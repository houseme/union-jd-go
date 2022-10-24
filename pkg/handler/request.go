/*
 * Copyright union-jd-go Author(https://houseme.github.io/union-jd-go/). All Rights Reserved.
 *
 * This Source Code Form is subject to the terms of the MIT License.
 * If a copy of the MIT was not distributed with this file,
 * You can obtain one at https://github.com/houseme/union-jd-go.
 */

package handler

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	"github.com/houseme/union-jd-go/config"
	"github.com/houseme/union-jd-go/entity"
)

type options struct {
	Method                            string `json:"method" dc:"API接口名称"`
	OpenCategoryGoodsGetRequest       entity.OpenCategoryGoodsGetRequest
	UnionOpenGoodsJingFenQueryRequest *entity.UnionOpenGoodsJingFenQueryRequest
}

// Option The option is a polaris option.
type Option func(o *options)

// UnionRequest is a request struct.
type UnionRequest struct {
	ctx    context.Context
	method string
	config *config.Config

	openCategoryGoodsGetRequest       *entity.OpenCategoryGoodsGetRequest       // 京东商品分类接口
	unionOpenGoodsJingFenQueryRequest *entity.UnionOpenGoodsJingFenQueryRequest // 京粉精选商品查询接口
}

// NewUnionRequest implements the use case for JD goods.
func NewUnionRequest(ctx context.Context, conf *config.Config, opts ...Option) *UnionRequest {
	ctx, span := gtrace.NewSpan(ctx, "tracing-union-handler-NewUnionRequest")
	defer span.End()

	op := options{}
	for _, opt := range opts {
		opt(&op)
	}
	return &UnionRequest{
		ctx:                               ctx,
		config:                            conf,
		method:                            op.Method,
		openCategoryGoodsGetRequest:       &op.OpenCategoryGoodsGetRequest,
		unionOpenGoodsJingFenQueryRequest: op.UnionOpenGoodsJingFenQueryRequest,
	}
}

// GetConfig .
func (s *UnionRequest) GetConfig() *config.Config {
	return s.config
}

// GetMethod .
func (s *UnionRequest) GetMethod() string {
	return s.method
}

// GetOpenCategoryGoodsGetRequest .
func (s *UnionRequest) GetOpenCategoryGoodsGetRequest() *entity.OpenCategoryGoodsGetRequest {
	return s.openCategoryGoodsGetRequest
}

// GetUnionOpenGoodsJingFenQueryRequest .
func (s *UnionRequest) GetUnionOpenGoodsJingFenQueryRequest() *entity.UnionOpenGoodsJingFenQueryRequest {
	return s.unionOpenGoodsJingFenQueryRequest
}

// WithMethod .
func WithMethod(method string) Option {
	return func(o *options) {
		o.Method = method
	}
}

// WithOpenCategoryGoodsGetRequest .
func WithOpenCategoryGoodsGetRequest(req *entity.OpenCategoryGoodsGetRequest) Option {
	return func(o *options) {
		o.OpenCategoryGoodsGetRequest = *req
	}
}

// WithUnionOpenGoodsJingFenQueryRequest .
func WithUnionOpenGoodsJingFenQueryRequest(req *entity.UnionOpenGoodsJingFenQueryRequest) Option {
	return func(o *options) {
		o.UnionOpenGoodsJingFenQueryRequest = req
	}
}

// UnionResponse is a response struct.
type UnionResponse struct {
	OpenCategoryGoodsGetResponse       *entity.OpenCategoryGoodsGetResponse       `json:"open_category_goods_get_response,omitempty"`
	UnionOpenGoodsJingFenQueryResponse *entity.UnionOpenGoodsJingFenQueryResponse `json:"union_open_goods_jingfen_query_response,omitempty"`
}
