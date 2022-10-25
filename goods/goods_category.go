/*
 * Copyright union-jd-go Author(https://houseme.github.io/union-jd-go/). All Rights Reserved.
 *
 * This Source Code Form is subject to the terms of the MIT License.
 * If a copy of the MIT was not distributed with this file,
 * You can obtain one at https://github.com/houseme/union-jd-go.
 */

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

// QueryCate query category
func (s *Goods) QueryCate(ctx context.Context, req *entity.OpenCategoryGoodsGetRequest) (res *entity.OpenCategoryGoodsGetResponse, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-union-jd-goods-QueryCate")
	defer span.End()

	s.config.Logger().Info(ctx, "query category start params:", req)

	if req == nil {
		err = gerror.New("query category request required")
		return
	}

	var (
		util    = pkg.NewUtil(s.config.Logger())
		request = handler.NewUnionRequest(ctx, s.config, handler.WithMethod(config.UnionOpenCategoryGoodsGet), handler.WithOpenCategoryGoodsGetRequest(req))
		resp    *handler.UnionResponse
	)

	if resp, err = util.Handler(ctx, request); err != nil {
		err = gerror.New("query category request failed")
		return
	}
	if resp == nil {
		err = gerror.New("query category response is nil")
		return
	}

	if resp.OpenCategoryGoodsGetResponse == nil {
		err = gerror.New("query category response is nil")
		return
	}
	s.config.Logger().Debug(ctx, "query category response:", resp)
	res = resp.OpenCategoryGoodsGetResponse
	s.config.Logger().Debug(ctx, "query category response:", res)
	return
}
