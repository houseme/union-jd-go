/*
 * Copyright union-jd-go Author(https://houseme.github.io/union-jd-go/). All Rights Reserved.
 *
 * This Source Code Form is subject to the terms of the MIT License.
 * If a copy of the MIT was not distributed with this file,
 * You can obtain one at https://github.com/houseme/union-jd-go.
 */

package entity

// OpenCategoryGoodsGetResponse .
type OpenCategoryGoodsGetResponse struct {
	JDUnionOpenCategoryGoodsGetResponse *JDUnionOpenCategoryGoodsGetResponse `json:"jd_union_open_category_goods_get_responce"`
}

// JDUnionOpenCategoryGoodsGetResponse jd union
type JDUnionOpenCategoryGoodsGetResponse struct {
	Code          uint   `json:"code,string"`
	ErrorMessage  string `json:"errorMessage"`
	ErrorSolution string `json:"errorSolution"`
	GetResult     string `json:"getResult"`
}

// CategoryResp get result
type CategoryResp struct {
	Data      []*CategoryItem `json:"data"`
	Message   string          `json:"message"`
	RequestID string          `json:"requestId"`
	Code      uint            `json:"code"`
}

// CategoryItem category item response
type CategoryItem struct {
	Name     string `json:"name"`
	Grade    uint   `json:"grade"`
	ID       uint   `json:"id"`
	ParentID uint   `json:"parentId"`
}

// OpenCategoryGoodsGetRequest .
type OpenCategoryGoodsGetRequest struct {
	CategoryReq *CategoryReq `json:"req"`
}

// CategoryReq .
type CategoryReq struct {
	ParentID uint `json:"parentId,string"`
	Grade    uint `json:"grade,string"`
}
