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
	"fmt"
	"reflect"
	"testing"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/glog"

	"github.com/houseme/union-jd-go/config"
	"github.com/houseme/union-jd-go/entity"
)

var (
	ctx  context.Context
	c    *config.Config
	good *Goods
	err  error
)

func init() {
	ctx = context.Background()
	logger := glog.New()
	logger.SetLevel(glog.LEVEL_INFO)
	if c, err = config.NewConfig(ctx, "26fff63a7f71792e826efb045f2eb31f", "af7cf2e9e2a843178da51", config.WithLogger(logger)); err != nil {
		panic(err)
	}
	if good, err = NewGoods(ctx, c); err != nil {
		panic(err)
	}
}

func TestGoods_GetConfig(t *testing.T) {
	type fields struct {
		ctx    context.Context
		config *config.Config
	}
	tests := []struct {
		name   string
		fields fields
		want   *config.Config
	}{
		// TODO: Add test cases.
		{
			name: "TestGoods_GetConfig",
			fields: fields{
				ctx:    ctx,
				config: c,
			},
			want: c,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Goods{
				ctx:    tt.fields.ctx,
				config: tt.fields.config,
			}
			if got := s.GetConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoods_QueryCate(t *testing.T) {
	type fields struct {
		ctx    context.Context
		config *config.Config
	}
	type args struct {
		ctx context.Context
		req *entity.OpenCategoryGoodsGetRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *entity.OpenCategoryGoodsGetResponse
		wantErr bool
	}{
		{
			name: "TestGoods_QueryCate",
			fields: fields{
				ctx:    ctx,
				config: c,
			},
			args: args{
				ctx: ctx,
				req: &entity.OpenCategoryGoodsGetRequest{
					CategoryReq: &entity.CategoryReq{
						ParentID: 0,
						Grade:    0,
					},
				},
			},
			wantRes: &entity.OpenCategoryGoodsGetResponse{
				JDUnionOpenCategoryGoodsGetResponse: &entity.JDUnionOpenCategoryGoodsGetResponse{
					Code: 0,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Goods{
				ctx:    tt.fields.ctx,
				config: tt.fields.config,
			}
			gotRes, err := s.QueryCate(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryCate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes.JDUnionOpenCategoryGoodsGetResponse.Code, tt.wantRes.JDUnionOpenCategoryGoodsGetResponse.Code) {
				t.Errorf("QueryCate() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			var list *entity.CategoryResp
			if err = gjson.New(gotRes.JDUnionOpenCategoryGoodsGetResponse.GetResult).Scan(&list); err != nil {
				t.Errorf("QueryCate() error = %v", err)
			}
			fmt.Println(list, "len:", len(list.Data))
		})
	}
}

func TestNewGoods(t *testing.T) {
	type args struct {
		ctx context.Context
		c   *config.Config
	}
	tests := []struct {
		name    string
		args    args
		want    *Goods
		wantErr bool
	}{
		{
			name: "TestNewGoods",
			args: args{
				ctx: ctx,
				c:   c,
			},
			want:    good,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGoods(tt.args.ctx, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGoods() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.config, tt.want.config) {
				t.Errorf("NewGoods() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoods_QueryGoods(t *testing.T) {
	type fields struct {
		ctx    context.Context
		config *config.Config
	}
	type args struct {
		req *entity.UnionOpenGoodsJingFenQueryRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *entity.UnionOpenGoodsJingFenQueryResponseTopLevel
		wantErr bool
	}{
		{
			name: "TestGoods_QueryGoods",
			fields: fields{
				ctx:    ctx,
				config: c,
			},
			args: args{
				req: &entity.UnionOpenGoodsJingFenQueryRequest{
					GoodsReq: &entity.JFGoodsReq{
						EliteID:   1,
						PageSize:  2,
						PageIndex: 1,
						SortName:  "inOrderCount30DaysSku",
						Sort:      "desc",
						Fields:    "",
						Pid:       "",
					},
				},
			},
			wantRes: &entity.UnionOpenGoodsJingFenQueryResponseTopLevel{
				UnionOpenGoodsJingFenQueryResponse: &entity.UnionOpenGoodsJingFenQueryResponse{
					Code: 0,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Goods{
				ctx:    tt.fields.ctx,
				config: tt.fields.config,
			}
			gotRes, err := s.QueryGoods(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryGoods() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes.UnionOpenGoodsJingFenQueryResponse.Code, tt.wantRes.UnionOpenGoodsJingFenQueryResponse.Code) {
				t.Errorf("QueryGoods() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			t.Log(gotRes.UnionOpenGoodsJingFenQueryResponse.QueryResult)
			var result *entity.UnionOpenGoodsJingFenQueryResponseResult
			if err = gjson.New(gotRes.UnionOpenGoodsJingFenQueryResponse.QueryResult).Scan(&result); err != nil {
				t.Errorf("QueryCate() error = %v", err)
			}
			fmt.Println("===============================")
			fmt.Println(result, "len:", len(result.Data), "item:", result.Data[0])
		})
	}
}

func TestGoods_QueryGoodsDetail(t *testing.T) {
	type fields struct {
		ctx    context.Context
		config *config.Config
	}
	type args struct {
		req *entity.UnionOpenGoodsBigFieldQueryRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *entity.UnionOpenGoodsBigFieldQueryResponseTopLevel
		wantErr bool
	}{
		{
			name: "TestGoods_QueryGoodsDetail",
			fields: fields{
				ctx:    ctx,
				config: c,
			},
			args: args{
				req: &entity.UnionOpenGoodsBigFieldQueryRequest{
					BigFieldGoodsReq: &entity.BigFieldGoodsReq{
						SkuIds: []int64{10051671714670, 10035657279233},
					},
				},
			},
			wantRes: &entity.UnionOpenGoodsBigFieldQueryResponseTopLevel{
				UnionOpenGoodsBigFieldQueryResponse: &entity.UnionOpenGoodsBigFieldQueryResponse{
					Code: 0,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Goods{
				ctx:    tt.fields.ctx,
				config: tt.fields.config,
			}
			gotRes, err := s.QueryGoodsDetail(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryGoodsDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes.UnionOpenGoodsBigFieldQueryResponse.Code, tt.wantRes.UnionOpenGoodsBigFieldQueryResponse.Code) {
				t.Errorf("QueryGoodsDetail() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			t.Log(gotRes.UnionOpenGoodsBigFieldQueryResponse.QueryResult)
			var result *entity.UnionOpenGoodsBigFieldQueryResult
			if err = gjson.New(gotRes.UnionOpenGoodsBigFieldQueryResponse.QueryResult).Scan(&result); err != nil {
				t.Errorf("QueryCate() error = %v", err)
			}
			fmt.Println("===============================")
			fmt.Println(result, "len:", len(result.Data), "item:", result.Data[0])
		})
	}
}
