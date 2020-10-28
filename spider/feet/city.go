package feet

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type City struct {
	code              string
	name              string
	tip               string
	subLevelModelList []City
	firstChar         string
	pinyin            string
	rank              int
	mark              int
	positionType      int
	cityType          int
}
type zpdata struct {
	hotCityList []City
}
type Cities struct {
	code    int
	message string
	zpData  zpdata
}

func fecth(usrl string) []byte {
	res, err := http.Get(usrl)
	if err == nil {
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err == nil {
			return body
		}
	}
	return nil
}

//https://www.zhipin.com/wapi/zpCommon/data/city.json
func (c *Cities) CityList(url string) error {
	if data := fecth(url); data != nil {
		return json.Unmarshal(data, c)
	}
	return nil
}
func (c City) Name() string {
	return c.name
}
func (c City) Code() int {
	return c.code
}
func (c City) Rank() int {
	return c.rank
}
func (c City) SubCities() []City {
	return c.subLevelModelList
}
func (c *Cities) GetCityByname(name string) City {
	return nil
}
