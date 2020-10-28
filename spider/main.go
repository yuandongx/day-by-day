package main

import (
	"fmt"
	"go.tour/spider/boss"
)

//https://www.zhipin.com/wapi/zpCommon/data/position.json
//https://www.zhipin.com/wapi/zpCommon/data/city.json
func main() {

	allC := boss.CityList("https://www.zhipin.com/wapi/zpCommon/data/city.json")
	for _, v := range allC.ZpData.HotCityList {
		fmt.Println(v.Name, v.SubLevelModelList)
	}
		for _, vv := range allC.ZpData.CityList {
		fmt.Println(vv.Name, vv.SubLevelModelList)
	}
}
