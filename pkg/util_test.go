/*
 * Copyright union-jd-go Author(https://houseme.github.io/union-jd-go/). All Rights Reserved.
 *
 * This Source Code Form is subject to the terms of the MIT License.
 * If a copy of the MIT was not distributed with this file,
 * You can obtain one at https://github.com/houseme/union-jd-go.
 */

package pkg

import (
	"context"
	"reflect"
	"testing"

	"github.com/gogf/gf/v2/os/glog"

	"github.com/houseme/union-jd-go/config"
	"github.com/houseme/union-jd-go/entity"
)

func TestNewUtil(t *testing.T) {
	type args struct {
		log glog.ILogger
	}
	tests := []struct {
		name string
		args args
		want *util
	}{
		{
			name: "TestNewUtil",
			args: args{log: glog.New()},
			want: &util{logger: glog.New()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUtil(tt.args.log); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUtil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_util_ConcatenateSignSource(t *testing.T) {
	type fields struct {
		logger glog.ILogger
	}
	type args struct {
		ctx  context.Context
		data interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "Test_util_ConcatenateSignSource",
			fields: fields{logger: glog.New()},
			args: args{ctx: context.Background(),
				data: struct {
					StringParam string `json:"stringParam"`
					IntParam    int64  `json:"intParam"`
					SortParam   int64  `json:"sortParam"`
				}{
					StringParam: "value",
					IntParam:    123456,
					SortParam:   1,
				},
			},
			want: "intParam123456sortParam1stringParamvalue",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &util{
				logger: tt.fields.logger,
			}
			if got := l.ConcatenateSignSource(tt.args.ctx, tt.args.data); got != tt.want {
				t.Errorf("ConcatenateSignSource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_util_MakeSign(t *testing.T) {
	type fields struct {
		logger glog.ILogger
	}
	type args struct {
		ctx       context.Context
		appSecret string
		req       entity.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "Test_util_MakeSign",
			fields: fields{logger: glog.New()},
			args: args{
				ctx:       context.Background(),
				appSecret: "MYAX8WegUVlOVSzEyxh7w-RyeT8WmfTAEHv-S0S48FFqeEJ6_3U6ci45OTQQOfx4xAz_JsyHjomONDR_TIKGv5l5X5aDATZpnLHKN",
				req: entity.Request{
					Method: config.UnionOpenCategoryGoodsGet,
				},
			},
			want:    "05DC543401CBE7A05B577A01BAF1863F",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &util{
				logger: tt.fields.logger,
			}
			got, err := l.MakeSign(tt.args.ctx, tt.args.appSecret, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeSign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MakeSign() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_util_NewRequest(t *testing.T) {
	type fields struct {
		logger glog.ILogger
	}
	type args struct {
		ctx    context.Context
		c      *config.Config
		method string
		params string
	}

	conf, _ := config.NewConfig(context.Background(), "A123456", "S12345678")

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantReq *entity.Request
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Test_util_NewRequest",
			fields: fields{logger: glog.New()},
			args: args{
				ctx:    context.Background(),
				c:      conf,
				method: config.UnionOpenCategoryGoodsGet,
				params: "JZK27bviLKM",
			},
			wantReq: &entity.Request{
				Method:       config.UnionOpenCategoryGoodsGet,
				AppKey:       conf.AppKey(),
				Format:       config.RequestFormat,
				Timestamp:    "",
				BuyParamJSON: "JZK27bviLKM",
				SignMethod:   config.RequestSignMethod,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &util{
				logger: tt.fields.logger,
			}
			gotReq, err := l.NewRequest(tt.args.ctx, tt.args.c, tt.args.method, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotReq.AppKey, tt.wantReq.AppKey) {
				t.Errorf("NewRequest() gotReq = %v, want %v", gotReq, tt.wantReq)
			}

			if !reflect.DeepEqual(gotReq.Format, tt.wantReq.Format) {
				t.Errorf("NewRequest() gotReq = %v, want %v", gotReq, tt.wantReq)
			}

			if !reflect.DeepEqual(gotReq.SignMethod, tt.wantReq.SignMethod) {
				t.Errorf("NewRequest() gotReq = %v, want %v", gotReq, tt.wantReq)
			}
		})
	}
}
