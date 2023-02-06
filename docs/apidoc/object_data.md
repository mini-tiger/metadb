#### 1. 已有模型插入单条数据
###### api接口格式：{{url}}8090/api/v3/create/instance/object/{objid}

##### api测试接口： http://{{url:port}}/api/v3/create/instance/object/datacenter
##### 方法: POST 
###### 测试接口备注：k8s namespace:cmdbv3      模型:datacenter 

##### 请求参数:

```

headers:
    HTTP_BLUEKING_SUPPLIER_ID: 0
    BK_User : admin
    Content-Type: application/json

body:
{
    "bk_inst_name": "北京亦庄联通数据中心", # metadb平台唯一ID
    "areaCode": "100001",
    "areaName": "华北区",
    "datacenterCode": "S003",
    "datacenterName": "北京亦庄联通数据中心",
    "datacenterTypeCode": "06",
    "datacenterTypeName": "自营",  
    "jfgsCode": "11",
    "jfgsName": "SBG2",
    "unifiedDatacenterCode": "S003",
    "usageStateNo": "1",
    "createdate": "2022-06-19"
}

```
##### 返回参数:
```
response:
{
    "result": true,
    "bk_error_code": 0,   # 0为正常  其它为失败
    "bk_error_msg": "success", # success 为正常
    "permission": null,
    "data": {
        "areaCode": "100001",
        "areaName": "华北区",
        "bk_inst_id": 6,
        "bk_inst_name": "1",
        "bk_obj_id": "datacenter",
        "bk_supplier_account": "0",
        "create_time": "2022-06-21T10:37:31.689+08:00",
        "createdate": "2022-06-19",
        "datacenterCode": "S003",
        "datacenterName": "1",
        "datacenterTypeCode": "06",
        "datacenterTypeName": "自营",
        "jfgsCode": "11",
        "jfgsName": "SBG2",
        "last_time": "2022-06-21T10:37:31.689+08:00",
        "unifiedDatacenterCode": "S003",
        "usageStateNo": "1"
    }
}


```
#### 2. 已有模型批量查找数据
###### api接口格式：{{url}}8090/api/v3/find/instassociation/object/{objid}

##### api测试接口：http://{{url:port}}/api/v3/find/instassociation/object/datacenter
##### 方法: POST 
###### 测试接口备注：k8s namespace:cmdbv3      模型:datacenter 

##### 请求参数:

```

headers:
    HTTP_BLUEKING_SUPPLIER_ID: 0
    BK_User : admin
    Content-Type: application/json

body:
{
    "condition": {
        "datacenter": [
            {
                "field": "datacenterCode",
                "operator": "$in",
                "value": [
                    "S003",
                    "S004"
                ]
            }
        ]
    },
    "fields": {},
    "page": {
        "start": 0,
        "limit": 20,
        "sort": "bk_inst_id"
    }
}
```

| 字段           | 含义                 | 备注               |
| ---            | ---                  |------------------|
| condition         |  查询参数      | 必须,以mongo查询参数为基础 |
| fields  |  筛选列败 | 非必须              |
| page   |  分页 默认10      | 非必须              |




##### 返回参数:
```

response:
{
    "result": true,
    "bk_error_code": 0,
    "bk_error_msg": "success",
    "permission": null,
    "data": {
        "count": 1,
        "info": [
            {
                "areaCode": "100001",
                "areaName": "华北区",
                "bk_inst_id": 1,
                "bk_inst_name": "北京亦庄联通数据中心",
                "bk_obj_id": "datacenter",
                "bk_supplier_account": "0",
                "create_time": "2022-06-21T02:46:44.95Z",
                "createdate": "2022-06-19",
                "datacenterCode": "S003",
                "datacenterName": "北京亦庄联通数据中心",
                "datacenterTypeCode": "06",
                "datacenterTypeName": "自营",
                "jfgsCode": "11",
                "jfgsName": "SBG2",
                "last_time": "2022-06-21T02:46:44.95Z",
                "unifiedDatacenterCode": "S003",
                "usageStateNo": "1"
            }
        ]
    }
}


```

| 字段           | 含义                 | 备注 |
| ---            | ---                  | ---  |
| created         |  创建后数据的id      |      |
| bk_error_code  |  0为正常  其它为失败 |      |
| bk_error_msg   |  success 为正常      |      |
| data.count   | 统计行数      |      |





#### 3. 已有模型更新单条数据
###### api接口格式：{{url}}8090/api/v3/update/instance/object/{objid}/inst/{bk_inst_id}

##### api测试接口：http://{{url:port}}/api/v3/update/instance/object/datacenter/inst/1
##### 方法: PUT
###### 测试接口备注：k8s namespace:cmdbv3      模型:datacenter 


##### 请求参数:

###### 注: **url 中 {bk_inst_id} 通过 [查找数据接口](#2-已有模型批量查找数据) 获取 字段 bk_inst_id**

```

headers:
    HTTP_BLUEKING_SUPPLIER_ID: 0
    BK_User : admin
    Content-Type: application/json

body:
{
    "areaCode": "100001",
    "areaName": "华北区",
    "bk_inst_name": "北京亦庄联通数据中心1",
    "bk_obj_id": "datacenter",
    "createdate": "2022-06-19",
    "datacenterCode": "S003",
    "datacenterName": "北京亦庄联通数据中心",
    "datacenterTypeCode": "06",
    "datacenterTypeName": "自营",
    "jfgsCode": "11",
    "jfgsName": "SBG2",
    "unifiedDatacenterCode": "S003",
    "usageStateNo": "1"
}
```

##### 请求参数:

| 字段   | 含义      | 备注 |
|------|---------| ---  |
| body | 模型字段map |  不必须全部字段    |





##### 返回参数:
```

response:
{
    "result": true,
    "bk_error_code": 0,
    "bk_error_msg": "success",
    "permission": null,
    "data": null
}


```

| 字段           | 含义                 | 备注 |
| ---            | ---                  | ---  |
| bk_error_code  |  0为正常  其它为失败 |      |
| bk_error_msg   |  success 为正常      |      |



#### 4. 已有模型删除数据
###### api接口格式：{{url}}8090/api/v3/deletemany/instance/object/{objid}

##### api测试接口：：http://{{url:port}}/api/v3/deletemany/instance/object/datacenter
##### 方法: DELETE
###### 测试接口备注：k8s namespace:cmdbv3      模型:datacenter 


##### 请求参数:

###### 注: **url 中 {inst_ids} 通过 [查找数据接口](#2-已有模型批量查找数据) 获取 字段 bk_inst_id **



```

headers:
    HTTP_BLUEKING_SUPPLIER_ID: 0
    BK_User : admin
    Content-Type: application/json

body:
    {
        "delete":{
            "inst_ids":[42871,32111]
        }
    }
```



##### 返回参数:
```

response:
{
    "result": true,
    "bk_error_code": 0,
    "bk_error_msg": "success",
    "permission": null,
    "data": null
}

```


| 字段           | 含义                 | 备注 |
| ---            | ---                  | ---  |
| bk_error_code  |  0为正常  其它为失败 |      |
| bk_error_msg   |  success 为正常      |      |

