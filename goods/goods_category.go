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

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/gtrace"

	"github.com/houseme/union-jd-go/config"
	"github.com/houseme/union-jd-go/entity"
	"github.com/houseme/union-jd-go/pkg"
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
		util = pkg.NewUtil(s.config.Logger())
		str  string
	)

	if str, err = gjson.New(req).ToJsonString(); err != nil {
		err = gerror.Wrap(err, "json to string err")
		return
	}
	s.config.Logger().Info(ctx, "query category params:", str)
	var reqe *entity.Request
	if reqe, err = util.NewRequest(ctx, s.config, config.UnionOpenCategoryGoodsGet, str); err != nil {
		err = gerror.Wrap(err, "new request err")
		return
	}
	var response *gclient.Response
	if response, err = g.Client().SetAgent(s.config.UserAgent()).Timeout(s.config.Timeout()).Get(ctx, s.config.ServerURL(), reqe); err != nil {
		err = gerror.Wrap(err, "request err")
		return
	}
	defer func() {
		if err := response.Close(); err != nil {
			s.config.Logger().Error(ctx, "response close err: ", err)
		}
	}()

	content := response.ReadAllString()
	s.config.Logger().Info(ctx, "request result content: ", content)
	if err = gjson.New(content).Scan(&res); err != nil {
		err = gerror.Wrap(err, "response read err")
		return
	}
	return
}
