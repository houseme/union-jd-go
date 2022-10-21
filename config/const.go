/*
 * Copyright union-jd-go Author(https://houseme.github.io/union-jd-go/). All Rights Reserved.
 *
 * This Source Code Form is subject to the terms of the MIT License.
 * If a copy of the MIT was not distributed with this file,
 * You can obtain one at https://github.com/houseme/union-jd-go.
 */

package config

import (
	"time"
)

const (
	// Version golang sdk version
	Version = "0.0.1"

	// ServerURL this is JD open server url;
	serverURL = "https://api.jd.com/routerjson"

	// userAgent name of jd user agent
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36"

	// Timeout request timeout
	Timeout = 5 * time.Second

	// defaultLogger is the default logger for the JD open.
	defaultLogger = "union-jd-go"
)

const (
	RequestFormat     = "json"
	RequestVersion    = "1.0"
	RequestTimestamp  = "Y-m-d H:i:s"
	RequestSignMethod = "md5"
)

const (
	// UnionOpenCategoryGoodsGet jd.union.open.category.goods.get
	// 商品类目查询接口
	// see https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.category.goods.get
	UnionOpenCategoryGoodsGet = "jd.union.open.category.goods.get"

	// DefaultParentID 父类目id 一级类目id为0
	// see https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.category.goods.get
	DefaultParentID uint = 0

	// GradeZero 类目级别(类目级别 0，1，2 代表一、二、三级类目)
	// see https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.category.goods.get
	GradeZero uint = 0
	// GradeOne 类目级别(类目级别 0，1，2 代表一、二、三级类目)
	// see https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.category.goods.get
	GradeOne uint = 1
	// GradeTwo 类目级别(类目级别 0，1，2 代表一、二、三级类目)
	// see https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.category.goods.get
	GradeTwo uint = 2
)
