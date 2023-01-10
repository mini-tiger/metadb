### 1.订阅事件接口

- API: POST /api/v3/event/subscribe/{bk_supplier_account}/{bk_biz_id}
- API 名称: subscribe_event
- 功能说明：
	- 中文：事件订阅
	- English：subscribe event

- input body:

``` json
{
  "subscription_name":"mysubscribe",
  "system_name":"SystemName",
  "callback_url":"http://127.0.0.1:8080/callback",
  "confirm_mode":"httpstatus",
  "confirm_pattern":"200",
  "subscription_form":"modulecreate,moduleupdate,moduledelete",
  "timeout":10
}
```

- input 字段说明

| 字段                | 类型   | 是否必须 | 默认值 | 说明                                             | Description                                                                       |
| ------------------- | ------ | -------- | ------ | ------------------------------------------------ | --------------------------------------------------------------------------------- |
| bk_biz_id           | int    | 是       | 无     | 业务id（目前为0）                                           | business id                                                                       |
| bk_supplier_account | string | 是       | 无     | 开发商账号  （目前为0）                                        | supplier account code                                                             |
| subscription_name   | string | 是       | 无     | 订阅的名字                                       | the subscription name                                                             |
| system_name         | string | 是       | 无     | 订阅事件的系统的名字                             | the subscriber's name                                                             |
| callback_url        | string | 是       | 无     | 回调函数                                         | the callbacks of the subscribers                                                  |
| confirm_mode        | string | 是       | 无     | 事件发送成功校验模式,可选 1-httpstatus,2-regular | confirm success mode of send to callback success, could be 1-httpstatus,2-regular |
| confirm_pattern     | string | 是       | 无     | callback的httpstatus或正则                       | the correct return httpstatus or regular                                          |
| subscription_form   | string | 是       | 无     | 订阅的事件,以逗号分隔                            | subcription event names, should split by comma                                    |
| timeout             | int    | 是       | 无     | 发送事件超时时间                                 | time out when send event message to callback                                      |


- output:

```
{
    "result":true,
    "bk_error_code":0,
    "bk_error_msg":"",
    "data":{
        "subscription_id": 1
    }
}
```

- output 字段说明

| 字段          | 类型   | 说明                                    | Description                                                |
| ------------- | ------ | --------------------------------------- | ---------------------------------------------------------- |
| result        | bool   | ture：成功，false：失败                 | true:success, false: failure                               |
| bk_error_code | int    | 错误编码。 0表示success，>0表示失败错误 | error code. 0 represent success, >0 represent failure code |
| bk_error_msg  | string | 请求失败返回的错误信息                  | error message from failed request                          |
| data          | object | 操作结果                                | the result                                                 |


data 字段说明

| 名称            | 类型 | 说明             | Description                    |
| --------------- | ---- | ---------------- | ------------------------------ |
| subscription_id | int  | 新增订阅的订阅ID | the id of the new subscription |



### 2. 推送数据格式

触发url 请求数据:

```
{
    "event_id":32467,
    "event_type":"instdata",
    "action":"delete",
    "action_time":"2023-01-09 12:08:55",
    "obj_type":"water_cooled_chiller",
    "data":[
        {
            "cur_data":null,
            "pre_data":{
                "SN":"",
                "bk_inst_id":30230,
                "bk_inst_name":"2",
                "bk_obj_id":"water_cooled_chiller",
                "bk_supplier_account":"0",
                "create_time":{
                    "$date":"2023-01-09T04:07:31.085Z"
                },
                "last_time":{
                    "$date":"2023-01-09T04:07:31.085Z"
                },
                "manufacturer":"",
                "model":"",
                "rated_power":1,
                "rated_temperature_of_chilled_water":null,
                "refrigerating_capacity":null,
                "system_design_pressure":null
            }
        }
    ],
    "bk_supplier_account":"0",
    "cursor":"MQ04DTYzYmI5MzAzNTE5MTk5NDA4NzgxOTk5OQ1kZWxldGUNMTY3MzIzNzMzNQ0y",
    "update_fields":[

    ],
    "deleted_fields":[

    ],
    "distribution_id":16,
    "subscription_id":1
}
```

- 触发url 字段说明

| 字段          | 类型   | 说明                                    | Description                                                |
| ------------- | ------ | --------------------------------------- | ---------------------------------------------------------- |
| event_type        | string   | 实例数据                 | true:success, false: failure                               |
|action | string    | 操作数据的动作  | create,update,delete |
|obj_type  | string | 模型名称                  |                           |
| data          | object | 数据变化对象，数据内容为模型定义的字段与数据值                                | cur_data 当前数据,pre_data 修改前数据       
|bk_supplier_account | string    | 数据账户  | 默认0|
|cursor | string    | 数据变化ID   | 后台使用 |
|distribution_id | int    | 分配事件ID  | 递增|
|subscription_id | int    | 订阅事件ID  |  |