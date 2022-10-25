/*
 * Copyright union-jd-go Author(https://houseme.github.io/union-jd-go/). All Rights Reserved.
 *
 * This Source Code Form is subject to the terms of the MIT License.
 * If a copy of the MIT was not distributed with this file,
 * You can obtain one at https://github.com/houseme/union-jd-go.
 */

package pkg

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"reflect"
	"sort"
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/houseme/union-jd-go/config"
	"github.com/houseme/union-jd-go/entity"
	"github.com/houseme/union-jd-go/pkg/handler"
)

type util struct {
	logger glog.ILogger
}

// NewUtil helper util
func NewUtil(log glog.ILogger) *util {
	return &util{
		logger: log,
	}
}

// NewRequest new request
func (l *util) NewRequest(ctx context.Context, c *config.Config, method, params string) (req *entity.Request, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-union-jd-util-NewRequest")
	defer span.End()

	l.logger.Debug(ctx, "new request start params:", params)
	req = &entity.Request{
		Method:       method,
		Version:      config.RequestVersion,
		Timestamp:    gtime.Now().Format(config.RequestTimestamp),
		Format:       config.RequestFormat,
		SignMethod:   config.RequestSignMethod,
		AppKey:       c.AppKey(),
		BuyParamJSON: params,
	}
	if req.Sign, err = l.MakeSign(ctx, c.AppSecret(), *req); err != nil {
		return nil, gerror.Wrap(err, "make sign failed")
	}
	l.logger.Debug(ctx, "new request end result:", req)
	return req, nil
}

// MakeSign make sign
func (l *util) MakeSign(ctx context.Context, appSecret string, req entity.Request) (string, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-union-jd-util-MakeSign")
	defer span.End()

	hash := md5.New()
	if _, err := hash.Write([]byte(appSecret + l.ConcatenateSignSource(ctx, req) + appSecret)); err != nil {
		err = gerror.Wrap(err, "write string failed")
		return "", err
	}
	return strings.ToUpper(hex.EncodeToString(hash.Sum(nil))), nil
}

// ConcatenateSignSource get sign url 排序并拼接签名的内容信息
func (l *util) ConcatenateSignSource(ctx context.Context, data interface{}) string {
	ctx, span := gtrace.NewSpan(ctx, "tracing-union-jd-util-ConcatenateSignSource")
	defer span.End()

	var (
		tt     = reflect.TypeOf(data)
		v      = reflect.ValueOf(data)
		count  = v.NumField()
		keys   = make([]string, count)
		params = make(map[string]string)
	)

	l.logger.Debug(ctx, "helper ConcatenateSignSource tt", tt, " v", v)

	for i := 0; i < count; i++ {
		if v.Field(i).CanInterface() { // 判断是否为可导出字段
			l.logger.Printf(ctx, "%s %s = %v -tag:%s", tt.Field(i).Name, tt.Field(i).Type, v.Field(i).Interface(), tt.Field(i).Tag)
			tagName := tt.Field(i).Tag.Get("json")
			keys[i] = tagName
			params[tagName] = gconv.String(v.Field(i).Interface())
		}
	}
	// sort params
	sort.Strings(keys)
	var buf bytes.Buffer
	for i := range keys {
		if params[keys[i]] == "" || keys[i] == "sign" {
			continue
		}
		if _, err := buf.WriteString(keys[i] + params[keys[i]]); err != nil {
			l.logger.Error(ctx, "helper ConcatenateSignSource buf.WriteString", err)
			return ""
		}
	}
	l.logger.Debug(ctx, "helper ConcatenateSignSource string start:", buf.String())
	return buf.String()
}

// Handler request handler
func (l *util) Handler(ctx context.Context, req *handler.UnionRequest) (resp *handler.UnionResponse, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-union-jd-util-Handler")
	defer span.End()

	l.logger.Debug(ctx, "handler start params:", req)
	var str string
	if str, err = l.ParamsToString(ctx, req); err != nil {
		err = gerror.Wrap(err, "json to string err")
		return
	}
	l.logger.Debug(ctx, "query category params:", str)
	var reqe *entity.Request
	if reqe, err = l.NewRequest(ctx, req.GetConfig(), req.GetMethod(), str); err != nil {
		err = gerror.Wrap(err, "new request err")
		return
	}

	var response *gclient.Response
	if response, err = g.Client().SetAgent(req.GetConfig().UserAgent()).Timeout(req.GetConfig().Timeout()).Get(ctx, req.GetConfig().ServerURL(), reqe); err != nil {
		err = gerror.Wrap(err, "request err")
		return
	}
	defer func() {
		if err := response.Close(); err != nil {
			l.logger.Error(ctx, "response close err: ", err)
		}
	}()

	var content = response.ReadAllString()
	l.logger.Debug(ctx, "request result content: ", content)
	if resp, err = l.ResponseToStruct(ctx, content, req.GetMethod()); err != nil {
		err = gerror.Wrap(err, "response to struct err")
		return
	}

	l.logger.Debug(ctx, "handler end result:", resp)
	return
}

// ResponseToStruct response to struct
func (l *util) ResponseToStruct(ctx context.Context, response, method string) (resp *handler.UnionResponse, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-union-jd-util-ResponseToStruct")
	defer span.End()

	l.logger.Debug(ctx, "response to struct start params:", response, " method:", method)
	switch method {
	case config.UnionOpenCategoryGoodsGet:
		resp, err = l.openCategoryGoodsGetResponse(ctx, response)
	case config.UnionOpenGoodsJingFenQuery:
		resp, err = l.unionOpenGoodsJingFenQueryResponse(ctx, response)
	default:
		resp = nil
		err = gerror.New("method not found")
	}
	l.logger.Debug(ctx, "response to struct end result:", resp)
	return
}

// openCategoryGoodsGetResponse open category goods get response
func (l *util) openCategoryGoodsGetResponse(ctx context.Context, response string) (resp *handler.UnionResponse, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-union-jd-util-openCategoryGoodsGetResponse")
	defer span.End()

	l.logger.Debug(ctx, "open category goods get response start params:", response)
	var result *entity.OpenCategoryGoodsGetResponse
	if err = gjson.New(response).Scan(&result); err != nil {
		err = gerror.Wrap(err, "response read err")
		return
	}

	if result == nil {
		err = gerror.New("response is nil")
		return
	}

	resp = &handler.UnionResponse{
		OpenCategoryGoodsGetResponse: result,
	}
	l.logger.Debug(ctx, "open category goods get response end result:", resp)
	return
}

// unionOpenGoodsJingFenQueryResponse union open goods jing fen query response
func (l *util) unionOpenGoodsJingFenQueryResponse(ctx context.Context, response string) (resp *handler.UnionResponse, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-union-jd-util-unionOpenGoodsJingFenQueryResponse")
	defer span.End()

	l.logger.Debug(ctx, "union open goods jing fen query response start params:", response)
	var result *entity.UnionOpenGoodsJingFenQueryResponseTopLevel
	if err = gjson.New(response).Scan(&result); err != nil {
		err = gerror.Wrap(err, "response read err")
		return
	}

	if result == nil {
		err = gerror.New("response is nil")
		return
	}

	resp = &handler.UnionResponse{
		UnionOpenGoodsJingFenQueryResponseTopLevel: result,
	}
	l.logger.Debug(ctx, "union open goods jing fen query response end result:", resp)
	return
}

// ParamsToString params to string
func (l *util) ParamsToString(ctx context.Context, req *handler.UnionRequest) (str string, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-union-jd-util-ParamsStr")
	defer span.End()

	l.logger.Debug(ctx, "params str start params:", req)
	var params interface{}
	switch req.GetMethod() {
	case config.UnionOpenCategoryGoodsGet:
		params = req.GetOpenCategoryGoodsGetRequest()
	case config.UnionOpenGoodsJingFenQuery:
		params = req.GetUnionOpenGoodsJingFenQueryRequest()
	default:
		params = nil
	}

	if params == nil {
		return "", gerror.New("params is nil")
	}

	if str, err = gjson.New(params).ToJsonString(); err != nil {
		err = gerror.Wrap(err, "json to string err")
		return
	}
	l.logger.Debug(ctx, "params str end result:", str)
	return str, nil
}
