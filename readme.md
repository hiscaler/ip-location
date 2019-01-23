IP 地址信息查询
===============

## 配置
您可以打开 config/config.json 文件进行相应的配置
```json
{
  "port": 8080
}
```

## 启动
```bash
./server
```

## 查询参数
执行完毕后，您可以通过 `http://127.0.0.1:8080/q?ip=您要查询的 IP` 访问即可，默认情况下，返回的数据为 JSON 格式。如果您需要返回其他格式，请参考下表：

| 参数 | 作用 |
| --- | --- |
| _format | 返回格式，支持 normal、json、jsonp |
| name | 当 _format 为 normal 格式时返回的变量名称，默认为 `_ipLocation` |
| callback | 当 _format 为 jsonp 格式时返回的回调名称，默认为 `callback` |

### 查询示例
1. http://127.0.0.1:8080/q?ip=118.249.188.173&_format=json
```json
{
    "success": true,
    "data": {
        "ip": "118.249.188.173",
        "name": "TaoBao",
        "isp_id": "100017",
        "isp_name": "电信",
        "country_id": "cn",
        "country_name": "中国",
        "area_id": "",
        "area_name": "",
        "province_id": "430000",
        "province_name": "湖南",
        "city_id": "430100",
        "city_name": "长沙",
        "region_id": "",
        "region_name": "",
        "address": ""
    }
}
```

2. http://127.0.0.1:8080/q?ip=118.249.188.173&_format=jsonp&callback=customCallback
```javascript
customCallback({
    "ip": "118.249.188.173",
    "name": "TaoBao",
    "isp_id": "100017",
    "isp_name": "电信",
    "country_id": "cn",
    "country_name": "中国",
    "area_id": "",
    "area_name": "",
    "province_id": "430000",
    "province_name": "湖南",
    "city_id": "430100",
    "city_name": "长沙",
    "region_id": "",
    "region_name": "",
    "address": ""
});
```

3. http://127.0.0.1:8080/q?ip=118.249.188.173&_format=normal&name=_customerIpLocation
```javascript
var _customerIpLocation={"ip":"118.249.188.173","name":"TaoBao","isp_id":"100017","isp_name":"电信","country_id":"cn","country_name":"中国","area_id":"","area_name":"","province_id":"430000","province_name":"湖南","city_id":"430100","city_name":"长沙","region_id":"","region_name":"","address":""};
```