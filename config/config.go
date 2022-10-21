/*
 * Copyright union-jd-go Author(https://houseme.github.io/union-jd-go/). All Rights Reserved.
 *
 * This Source Code Form is subject to the terms of the MIT License.
 * If a copy of the MIT was not distributed with this file,
 * You can obtain one at https://github.com/houseme/union-jd-go.
 */

package config

import (
	"context"
	"errors"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
)

// Config this is config for JD
type Config struct {
	ctx       context.Context
	appKey    string
	appSecret string
	serverURL string
	opt       options
}

type options struct {
	// AccessToken is for identify with JD
	AccessToken string
	// Timeout request timeout, default is `5 * time.Second`.
	Timeout time.Duration
	// Logger logger, nil by default
	Logger glog.ILogger
	// UserAgent http request User-Agent header
	UserAgent string
}

// Option The option is a polaris option.
type Option func(o *options)

// WithAccessToken set access token at runtime
func WithAccessToken(accessToke string) Option {
	return func(o *options) {
		o.AccessToken = accessToke
	}
}

// WithTimeout set timeout
func WithTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.Timeout = timeout
	}
}

// WithLogger set logger
func WithLogger(l glog.ILogger) Option {
	return func(o *options) {
		o.Logger = l
	}
}

func WithUserAgent(userAgent string) Option {
	return func(o *options) {
		o.UserAgent = userAgent
	}
}

// NewConfig init JD config
func NewConfig(ctx context.Context, appKey, appSecret string, opts ...Option) (*Config, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-union-config-newConfig")
	defer span.End()

	if len(gstr.Trim(appKey)) == 0 {
		return nil, errors.New("appKey is empty")
	}

	if len(gstr.Trim(appSecret)) == 0 {
		return nil, errors.New("appSecret is empty")
	}

	op := options{
		Timeout:     Timeout,
		Logger:      g.Log(defaultLogger),
		UserAgent:   userAgent,
		AccessToken: "",
	}

	for _, option := range opts {
		option(&op)
	}

	return &Config{
		ctx:       ctx,
		appKey:    appKey,
		appSecret: appSecret,
		serverURL: serverURL,
		opt:       op,
	}, nil
}

// AppKey return appKey
func (c *Config) AppKey() string {
	return c.appKey
}

// AppSecret return appKey
func (c *Config) AppSecret() string {
	return c.appSecret
}

// AccessToken return accessToken
func (c *Config) AccessToken() string {
	return c.opt.AccessToken
}

// Logger return logger
func (c *Config) Logger() glog.ILogger {
	return c.opt.Logger
}

// ServerURL return server url
func (c *Config) ServerURL() string {
	return c.serverURL
}

// Timeout return timeout
func (c *Config) Timeout() time.Duration {
	return c.opt.Timeout
}

// UserAgent return header user-agent
func (c *Config) UserAgent() string {
	return c.opt.UserAgent
}
