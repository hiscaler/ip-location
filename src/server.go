package main

import (
	"encoding/json"
	"fmt"
	"io"
	"location"
	"log"
	"net/http"
	"response"
	"strings"
)

const (
	FormatNormal = "normal"
	FormatJson   = "json"
	FormatJsonp  = "jsonp"
)

func QueryIpInformation(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	format := strings.ToLower(q.Get("_format"))
	contentType := ""
	switch format {
	case FormatJson:
		contentType = "application/json; charset=UTF-8"

	default:
		contentType = "application/javascript; charset=UTF-8"
	}
	w.Header().Set("Content-Type", contentType)
	resp := response.Response{}
	ip := strings.TrimSpace(q.Get("ip"))
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

	data := ""
	switch format {
	case FormatNormal:
		b, _ := json.Marshal(resp.Data)
		variableName := strings.TrimSpace(q.Get("name"))
		if variableName == "" {
			variableName = "_ipLocation"
		}
		data = fmt.Sprintf("var %s=%s;", variableName, string(b))

	case FormatJsonp:
		b, _ := json.Marshal(resp.Data)
		callback := strings.TrimSpace(q.Get("callback"))
		if callback == "" {
			callback = "callback"
		}
		data = fmt.Sprintf("%s(%s);", callback, string(b))

	default:
		b, _ := json.Marshal(resp)
		data = string(b)
	}

	io.WriteString(w, data)
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
