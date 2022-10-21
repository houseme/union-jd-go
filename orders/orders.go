/*
 * Copyright union-jd-go Author(https://houseme.github.io/union-jd-go/). All Rights Reserved.
 *
 * This Source Code Form is subject to the terms of the MIT License.
 * If a copy of the MIT was not distributed with this file,
 * You can obtain one at https://github.com/houseme/union-jd-go.
 */

// Package orders is an internal service.
package orders

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/gtrace"

	"github.com/houseme/union-jd-go/config"
)

// Orders is a jd service.
type Orders struct {
	ctx    context.Context
	config *config.Config
}

// NewOrders returns an orders service.
func NewOrders(ctx context.Context, c *config.Config) (*Orders, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-union-jd-orders-NewOrders")
	defer span.End()

	if c == nil {
		return nil, gerror.New("JD configuration required")
	}

	return &Orders{
		ctx:    ctx,
		config: c,
	}, nil
}

// GetConfig .
func (o *Orders) GetConfig() *config.Config {
	return o.config
}
