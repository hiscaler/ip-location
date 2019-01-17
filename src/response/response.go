package response

type Response struct {
	Success bool `json:"success"`
	Data    struct {
		Ip           string `json:"ip"`
		IspId        string `json:"isp_id"`
		IspName      string `json:"isp_name"`
		CountryId    string `json:"country_id"`
		CountryName  string `json:"country_name"` //  国家
		AreaId       string `json:"area_id"`
		AreaName     string `json:"area_name"` // 地区
		ProvinceId   string `json:"province_id"`
		ProvinceName string `json:"province_name"` // 省
		CityId       string `json:"city_id"`
		CityName     string `json:"city_name"` // 市
		RegionId     string `json:"region_id"`
		RegionName   string `json:"region_name"` // 区
		Address      string `json:"address"`     // 详细地址
	} `json:"data"`
}
