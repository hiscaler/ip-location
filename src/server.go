package main

import (
	"config"
	"encoding/json"
	"fmt"
	"io"
	"location"
	"log"
	"net/http"
	"response"
	"strconv"
	"strings"
)

const (
	FormatNormal = "normal"
	FormatJson   = "json"
	FormatJsonp  = "jsonp"
)

var cfg *config.Config

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
	cfg = config.NewConfig()
}

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
	ch := make(chan location.Location)

	ok := make(chan bool)
	locations := []string{
		"pconline",
		"ipip",
		"ipapi",
		"taobao",
	}

	for _, locationType := range locations {
		var base location.Location
		var err error
		switch locationType {
		case "pconline":
			go func() {
				l := location.PcOnlineLocation{}
				l.SetIp(ip)
				if base, err = l.Find(); err == nil {
					ok <- true
					ch <- base
				} else {
					log.Println("Error: ", err)
				}
			}()

		case "ipip":
			go func() {
				l := location.IpIpLocation{}
				l.SetIp(ip)
				if base, err = l.Find(); err == nil {
					ok <- true
					ch <- base
				} else {
					log.Println("Error: ", err)
				}
			}()

		case "ipapi":
			go func() {
				l := location.IpApiLocation{}
				l.SetIp(ip)
				if base, err = l.Find(); err == nil {
					ok <- true
					ch <- base
				} else {
					log.Println("Error: ", err)
				}
			}()

		case "taobao":
			go func() {
				l := location.TaoBaoLocation{}
				l.SetIp(ip)
				if base, err = l.Find(); err == nil {
					ok <- true
					ch <- base
				} else {
					log.Println("Error: ", err)
				}
			}()
		}

	}

	go func() {
		select {
		case ipLocation := <-ch:
			resp.Success = ipLocation.Success
			if ipLocation.Success {
				resp.Data.Name = ipLocation.Name
				resp.Data.Ip = ipLocation.Ip
				resp.Data.IspId = ipLocation.IspId
				resp.Data.IspName = ipLocation.IspName
				resp.Data.CountryId = ipLocation.CountryId
				resp.Data.CountryName = ipLocation.CountryName
				resp.Data.AreaId = ipLocation.AreaId
				resp.Data.AreaName = ipLocation.AreaName
				resp.Data.ProvinceId = ipLocation.ProvinceId
				resp.Data.ProvinceName = ipLocation.ProvinceName
				resp.Data.CityId = ipLocation.CityId
				resp.Data.CityName = ipLocation.CityName
				resp.Data.RegionId = ipLocation.RegionId
				resp.Data.RegionName = ipLocation.RegionName
				resp.Data.Address = ipLocation.Address
			}

			log.Println("Query IP: "+ip, fmt.Sprintf("%+v", resp))
		}
	}()
	<-ok

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
	port := strconv.Itoa(cfg.Port)
	go func() {
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatal("Start serve failed, ", err)
		}
	}()
	log.Printf("Server started successful, port is " + port + ", please access http://127.0.0.1:" + port + "/q?ip=you-ip-address to query you ip location information.")
	<-done
}
