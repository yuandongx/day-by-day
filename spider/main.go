package main

import (
	"fmt"
	"go-tour/spider/feet"
)

//https://www.zhipin.com/wapi/zpCommon/data/position.json
//https://www.zhipin.com/wapi/zpCommon/data/city.json
func main() {
	text := feet.Fecth("https://www.zhipin.com/wapi/zpCommon/data/city.json")
	fmt.Println(text)
}
