/*
 * Copyright union-jd-go Author(https://houseme.github.io/union-jd-go/). All Rights Reserved.
 *
 * This Source Code Form is subject to the terms of the MIT License.
 * If a copy of the MIT was not distributed with this file,
 * You can obtain one at https://github.com/houseme/union-jd-go.
 */

package test

import (
	"context"
	"reflect"
	"testing"

	"github.com/gogf/gf/v2/frame/g"

	jd "github.com/houseme/union-jd-go"
	"github.com/houseme/union-jd-go/config"
	"github.com/houseme/union-jd-go/goods"
	"github.com/houseme/union-jd-go/orders"
)

var (
	ctx       = context.Background()
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36"
	conf      *config.Config
	good      *goods.Goods
	order     *orders.Orders
	err       error
)

func init() {
	if conf, err = config.NewConfig(ctx, "1111", "1111", config.WithLogger(g.Log("union-jd-go")), config.WithUserAgent(userAgent)); err != nil {
		panic(err)
	}

	union := jd.NewUnionJD()
	if good, err = union.GetGoods(ctx, conf); err != nil {
		panic(err)
	}
	if order, err = union.GetOrders(ctx, conf); err != nil {
		panic(err)
	}
}

func TestNewUnionJD(t *testing.T) {
	tests := []struct {
		name string
		want *jd.UnionJD
	}{
		{
			name: "TestNewUnionJD",
			want: &jd.UnionJD{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jd.NewUnionJD(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUnionJD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionJD_GetGoods(t *testing.T) {
	type args struct {
		ctx context.Context
		c   *config.Config
	}
	tests := []struct {
		name    string
		args    args
		want    *goods.Goods
		wantErr bool
	}{
		{
			name: "TestUnionJD_GetGoods",
			args: args{
				ctx: ctx,
				c:   conf,
			},
			want:    good,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &jd.UnionJD{}
			got, err := u.GetGoods(tt.args.ctx, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGoods() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.GetConfig(), tt.want.GetConfig()) {
				t.Errorf("GetGoods() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionJD_GetOrders(t *testing.T) {
	type args struct {
		ctx context.Context
		c   *config.Config
	}
	tests := []struct {
		name    string
		args    args
		want    *orders.Orders
		wantErr bool
	}{
		{
			name: "TestUnionJD_GetOrders",
			args: args{
				ctx: ctx,
				c:   conf,
			},
			want:    order,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &jd.UnionJD{}
			got, err := u.GetOrders(tt.args.ctx, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.GetConfig(), tt.want.GetConfig()) {
				t.Errorf("GetOrders() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionJD_Version(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "TestUnionJD_Version",
			want: config.Version,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &jd.UnionJD{}
			if got := u.Version(); got != tt.want {
				t.Errorf("Version() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkUnionJD_Version(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u := &jd.UnionJD{}
		u.Version()
	}
}

func Test_Main(t *testing.T) {
	var x = 0.0
	var a = x / x
	var m = map[float64]int{a: 1}
	m[a] = 2
	for k := range m {
		delete(m, k)
	}
	println(len(m))
}
