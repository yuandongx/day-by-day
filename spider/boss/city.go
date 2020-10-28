package boss

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type City struct {
	Code              int
	Name              string
	Tip               string
	SubLevelModelList []City
	FirstChar         string
	Pinyin            string
	Rank              float64
	Mark              float64
	PositionType      float64
	CityType          float64
}
type zpdata struct {
	HotCityList []City
	CityList    []City
}
type Cities struct {
	Code    int `json:"code"`
	Message string `json:"message"`
	ZpData  zpdata `json:"zpData"`
}

func fecth(usrl string) ([]byte, error) {
	res, err := http.Get(usrl)
	if err == nil {
		defer res.Body.Close()
  }
	return ioutil.ReadAll(res.Body)
}

//https://www.zhipin.com/wapi/zpCommon/data/city.json
func CityList(url string) Cities {
	cc := Cities{}
	if data, err := fecth(url); err == nil {
		json.Unmarshal(data, &cc)
	}
	return cc
}
