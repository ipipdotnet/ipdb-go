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

## 支持IPDB格式
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
## 数据字段说明
<pre>
country_name : 国家名字 
region_name  : 省名字 
city_name    : 城市名字 
owner_domain : 所有者  
isp_domain  : 运营商 
latitude  :  纬度  
longitude : 经度  
timezone : 时区    
utc_offset : UTC时区   
china_admin_code : 中国行政区划代码 
idd_code : 国家电话号码前缀 
country_code : 国家2位代码
continent_code : 大洲代码   
idc : IDC |  VPN  
base_station : 基站 | WIFI 
country_code3 : 国家3位代码 
european_union : 是否为欧盟成员国： 1 | 0 
currency_code : 当前国家货币代码  
currency_name : 当前国家货币名称  
anycast : ANYCAST   
</pre>
