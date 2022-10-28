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
	// RequestFormat is the format of the request.
	RequestFormat = "json"
	// RequestVersion is the version of the request.
	RequestVersion = "1.0"
	// RequestTimestamp is the timestamp of the request.
	RequestTimestamp = "Y-m-d H:i:s"
	// RequestSignMethod is the sign method of the request.
	RequestSignMethod = "md5"
)

const (
	// UnionOpenCategoryGoodsGet jd.union.open.category.goods.get
	// 商品类目查询接口
	// see https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.category.goods.get
	UnionOpenCategoryGoodsGet = "jd.union.open.category.goods.get"

	// UnionOpenGoodsJingFenQuery jd.union.open.goods.jingfen.query
	// 京东联盟精选优质商品查询接口
	// see https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.goods.jingfen.query
	UnionOpenGoodsJingFenQuery = "jd.union.open.goods.jingfen.query"

	// UnionOpenGoodsBigFieldQuery jd.union.open.goods.bigfield.query
	// 商品详情查询接口,大字段信息
	// see https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.goods.bigfield.query
	UnionOpenGoodsBigFieldQuery = "jd.union.open.goods.bigfield.query"

	// UnionOpenGoodsMaterialQuery jd.union.open.goods.material.query
	// 猜你喜欢商品推荐
	// See https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.goods.material.query
	UnionOpenGoodsMaterialQuery = "jd.union.open.goods.material.query"

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

// EliteID 分类ID类型
type EliteID uint

const (
	// GoodCoupon  1-好券商品
	GoodCoupon EliteID = 1 // 1-好券商品
	// SuperHypermarket 2-超级大卖场
	SuperHypermarket EliteID = 2 // 2-超级大卖场
	// NineDivision 10-9.9专区
	NineDivision EliteID = 10 // 10-9.9专区
	// HotSell 22-热销爆品
	HotSell EliteID = 22 // 22-热销爆品
	// Commend 23-为你推荐
	Commend EliteID = 23 // 23-为你推荐
	// DigitalHomeAppliance 24-数码家电
	DigitalHomeAppliance EliteID = 24 // 24-数码家电
	// SuperMarket 25-京东超市
	SuperMarket EliteID = 25 // 25-超市
	// MotherAndBabyToys 26-母婴玩具
	MotherAndBabyToys EliteID = 26 // 26-母婴玩具
	// FurnitureDaily 27-家具日用
	FurnitureDaily EliteID = 27 // 27-家具日用
	// BeautyMakeup  28-美妆穿搭
	BeautyMakeup EliteID = 28 // 28-美妆穿搭
	// HealthCare 29-医药保健
	HealthCare EliteID = 29 // 29-医药保健
	// BooksStationary 30-图书文具
	BooksStationary EliteID = 30 // 30-图书文具
	// TodayRecommend 31-今日必推
	TodayRecommend EliteID = 31 // 31-今日必推
	// BrandHQGoods 32-品牌好货
	BrandHQGoods EliteID = 32 // 32-品牌好货
	// SeckillGoods 33-秒杀商品
	SeckillGoods EliteID = 33 // 33-秒杀商品
	// PinGouGoods 34-拼购商品
	PinGouGoods EliteID = 34 // 34-拼购商品
	// HighIncome 40-高收益
	HighIncome EliteID = 40 // 40-高收益
	// SelfSupportHotSell 41-自营热销榜
	SelfSupportHotSell EliteID = 41 // 41-自营热卖榜
	// NewArrival 109-新品首发
	NewArrival EliteID = 109 // 109-新品首发
	// SelfSupport 110-自营
	SelfSupport EliteID = 110 // 110-自营
	// FirstPurchase 125-首购商品
	FirstPurchase EliteID = 125 // 125-首购商品
	// HighCommission 129-高佣榜单
	HighCommission EliteID = 129 // 129-高佣榜单
	// VideoGoods  130-视频商品
	VideoGoods EliteID = 130 // 130-视频商品
)

const (
	// SortNamePrice 排序字段(price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30DaysSku：sku维度30天引单量，comments：评论数，goodComments：好评数)
	SortNamePrice = "price" // 价格
	// SortNameCommissionShare 排序字段(price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30DaysSku：sku维度30天引单量，comments：评论数，goodComments：好评数)
	SortNameCommissionShare = "commissionShare" // 佣金比例
	// SortNameCommission 排序字段(price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30DaysSku：sku维度30天引单量，comments：评论数，goodComments：好评数)
	SortNameCommission = "commission" // 佣金
	// SortNameInOrderCount30DaysSku 排序字段(price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30DaysSku：sku维度30天引单量，comments：评论数，goodComments：好评数)
	SortNameInOrderCount30DaysSku = "inOrderCount30DaysSku" // sku维度30天引单量
	// SortNameComments 排序字段(price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30DaysSku：sku维度30天引单量，comments：评论数，goodComments：好评数)
	SortNameComments = "comments" // 评论数
	// SortNameGoodComments 排序字段(price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30DaysSku：sku维度30天引单量，comments：评论数，goodComments：好评数)
	SortNameGoodComments = "goodComments" // 好评数
)

const (
	// SortByDesc asc,desc升降序,默认降序
	SortByDesc = "desc" // 降序
	// SortByAsc asc,desc升降序,默认降序
	SortByAsc = "asc" // 升序
)

const (
	// CategoryInfo 查询域集合，不填写则查询全部，
	// 目目前支持：categoryInfo（类目信息）,imageInfo（图片信息）,baseBigFieldInfo（基础大字段信息）,bookBigFieldInfo（图书大字段信息）,videoBigFieldInfo（影音大字段信息）,detailImages（商详图）

	// CategoryInfo categoryInfo（类目信息）
	CategoryInfo = "categoryInfo" // 类目信息
	// ImageInfo imageInfo（图片信息）
	ImageInfo = "imageInfo" // 图片信息
	// BaseBigFieldInfo baseBigFieldInfo（基础大字段信息）
	BaseBigFieldInfo = "baseBigFieldInfo" // 基础大字段信息
	// BookBigFieldInfo bookBigFieldInfo（图书大字段信息）
	BookBigFieldInfo = "bookBigFieldInfo" // 图书大字段信息
	// VideoBigFieldInfo videoBigFieldInfo（影音大字段信息）
	VideoBigFieldInfo = "videoBigFieldInfo" // 影音大字段信息
	// DetailImages detailImages（商详图）
	DetailImages = "detailImages" // 商详图
)
