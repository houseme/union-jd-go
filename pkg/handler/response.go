/*
 * Copyright union-jd-go Author(https://houseme.github.io/union-jd-go/). All Rights Reserved.
 *
 * This Source Code Form is subject to the terms of the MIT License.
 * If a copy of the MIT was not distributed with this file,
 * You can obtain one at https://github.com/houseme/union-jd-go.
 */

package handler

import (
	"github.com/houseme/union-jd-go/entity"
)

// UnionResponse is a response struct.
type UnionResponse struct {
	OpenCategoryGoodsGetResponse               *entity.OpenCategoryGoodsGetResponse               `json:"open_category_goods_get_response,omitempty"`
	UnionOpenGoodsJingFenQueryResponseTopLevel *entity.UnionOpenGoodsJingFenQueryResponseTopLevel `json:"union_open_goods_jingfen_query_response,omitempty"`
}
