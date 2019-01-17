package main

import (
	"encoding/json"
	"fmt"
	"io"
	"location"
	"log"
	"net/http"
	"response"
)

func QueryIpInformation(w http.ResponseWriter, req *http.Request) {
	resp := response.Response{}
	ip := req.URL.Query().Get("ip")
	ipLocation := location.PcOnlineLocation{}
	ipLocation.SetIp(ip)
	if v, err := ipLocation.Find(); err == nil {
		resp.Success = v.Success
		if v.Success {
			resp.Data.Ip = v.Ip
			resp.Data.IspId = v.IspId
			resp.Data.IspName = v.IspName
			resp.Data.CountryId = v.CountryId
			resp.Data.CountryName = v.CountryName
			resp.Data.AreaId = v.AreaId
			resp.Data.AreaName = v.AreaName
			resp.Data.ProvinceId = v.ProvinceId
			resp.Data.ProvinceName = v.ProvinceName
			resp.Data.CityId = v.CityId
			resp.Data.CityName = v.CityName
			resp.Data.RegionId = v.RegionId
			resp.Data.RegionName = v.RegionName
			resp.Data.Address = v.Address
		}
	} else {
		log.Println("Error: ", err)
	}
	log.Println("Query IP: "+ip, fmt.Sprintf("%+v", resp))

	b, _ := json.Marshal(resp)
	io.WriteString(w, string(b))
}

// IP 地址信息查询
func main() {
	log.Println("Begin start server")
	http.HandleFunc("/q", QueryIpInformation)
	done := make(chan bool)
	port := "8765"
	go func() {
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatal("Start serve failed, ", err)
		}
	}()
	log.Printf("Server started successful, port is " + port)
	<-done
}
