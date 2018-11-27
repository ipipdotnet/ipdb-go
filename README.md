# ipdb-go
[![TravisCI Build Status](https://travis-ci.org/ipipdotnet/ipdb-go.svg?branch=master)](https://travis-ci.org/ipipdotnet/ipdb-go)
[![Coverage Status](https://coveralls.io/repos/github/ipipdotnet/ipdb-go/badge.svg?branch=master)](https://coveralls.io/github/ipipdotnet/ipdb-go?branch=master)
[![IPDB Database API Document](https://godoc.org/github.com/ipipdotnet/ipdb-go?status.svg)](https://godoc.org/github.com/ipipdotnet/ipdb-go)

IPIP.net officially supported IP database ipdb format parsing library

# Installing
<code>
    go get github.com/ipipdotnet/ipdb-go
</code>

# Code Example

## 支持IPDB格式地级市精度IP离线库(免费版，每周高级版，每日标准版，每日高级版，每日专业版，每日旗舰版)
<pre>
<code>
package main

import (
	"github.com/ipipdotnet/ipdb-go"
	"fmt"
	"log"
)

func main() {
	db, err := ipdb.NewCity("/path/to/city.ipv4.ipdb")
	if err != nil {
		log.Fatal(err)
	}

	db.Reload("/path/to/city.ipv4.ipdb") // 更新 ipdb 文件后可调用 Reload 方法重新加载内容

	fmt.Println(db.IsIPv4()) // check database support ip type
	fmt.Println(db.IsIPv6()) // check database support ip type
	fmt.Println(db.BuildTime()) // database build time
	fmt.Println(db.Languages()) // database support language
	fmt.Println(db.Fields()) // database support fields

	fmt.Println(db.FindInfo("2001:250:200::", "CN")) // return CityInfo
	fmt.Println(db.Find("1.1.1.1", "CN")) // return []string
	fmt.Println(db.FindMap("118.28.8.8", "CN")) // return map[string]string
	fmt.Println(db.FindInfo("127.0.0.1", "CN")) // return CityInfo

	fmt.Println()
}
</code>
</pre>
## 地级市精度库数据字段说明
<pre>
country_name : 国家名字 （每周高级版及其以上版本包含）
region_name  : 省名字   （每周高级版及其以上版本包含）
city_name    : 城市名字 （每周高级版及其以上版本包含）
owner_domain : 所有者   （每周高级版及其以上版本包含）
isp_domain  : 运营商 （每周高级版与每日高级版及其以上版本包含）
latitude  :  纬度   （每日标准版及其以上版本包含）
longitude : 经度    （每日标准版及其以上版本包含）
timezone : 时区     （每日标准版及其以上版本包含）
utc_offset : UTC时区    （每日标准版及其以上版本包含）
china_admin_code : 中国行政区划代码 （每日标准版及其以上版本包含）
idd_code : 国家电话号码前缀 （每日标准版及其以上版本包含）
country_code : 国家2位代码  （每日标准版及其以上版本包含）
continent_code : 大洲代码   （每日标准版及其以上版本包含）
idc : IDC |  VPN   （每日专业版及其以上版本包含）
base_station : 基站 | WIFI （每日专业版及其以上版本包含）
country_code3 : 国家3位代码 （每日专业版及其以上版本包含）
european_union : 是否为欧盟成员国： 1 | 0 （每日专业版及其以上版本包含）
currency_code : 当前国家货币代码    （每日旗舰版及其以上版本包含）
currency_name : 当前国家货币名称    （每日旗舰版及其以上版本包含）
anycast : ANYCAST       （每日旗舰版及其以上版本包含）
</pre>
## 适用于IPDB格式的中国地区 IPv4 区县库
<pre>
db, err := ipdb.NewDistrict("/path/to/quxian.ipdb")
if err != nil {
	log.Fatal(err)
}
fmt.Println(db.IsIPv4())    // check database support ip type
fmt.Println(db.IsIPv6())    // check database support ip type
fmt.Println(db.Languages()) // database support language
fmt.Println(db.Fields())    // database support fields

fmt.Println(db.Find("1.12.7.255", "CN"))
fmt.Println(db.FindMap("2001:250:200::", "CN"))
fmt.Println(db.FindInfo("1.12.7.255", "CN"))

fmt.Println()
</pre>

## 适用于IPDB格式的基站 IPv4 库
<pre>
db, err := ipdb.NewBaseStation("/path/to/station_ip.ipdb")
if err != nil {
	log.Fatal(err)
}

fmt.Println(db.FindMap("223.220.223.255", "CN"))
</pre>