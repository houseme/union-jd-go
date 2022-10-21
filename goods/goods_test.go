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
	"reflect"
	"testing"

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
	if c, err = config.NewConfig(ctx, "26fff63a7f71792e826efb045f2eb31f", "cd8eb3c6c6ed4731b6545db44585476e"); err != nil {
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
			wantRes: nil,
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
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("QueryCate() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
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
