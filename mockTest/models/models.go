package models

import (
	"fmt"
	"strconv"
	"time"
)

/**
 * @Author: Tao Jun
 * @Since: 2022/7/13
 * @Desc: models.go
**/
type Instance struct {
	AreaCode              string    `json:"areaCode"`
	AreaName              string    `json:"areaName"`
	BkInstName            string    `json:"bk_inst_name"`
	Createdate            string    `json:"createdate"`
	DatacenterCode        string    `json:"datacenterCode"`
	DatacenterName        string    `json:"datacenterName"`
	DatacenterTypeCode    string    `json:"datacenterTypeCode"`
	DatacenterTypeName    string    `json:"datacenterTypeName"`
	JfgsCode              string    `json:"jfgsCode"`
	JfgsName              string    `json:"jfgsName"`
	LastTime              time.Time `json:"last_time"`
	UnifiedDatacenterCode string    `json:"unifiedDatacenterCode"`
	UsageStateNo          string    `json:"usageStateNo"`
}

func GetMockInstance(num int) []*Instance {
	mockData := make([]*Instance, 0, num)
	for i := 0; i < num; i++ {
		mockData = append(mockData, &Instance{
			AreaCode:              "10001",
			AreaName:              "华北区",
			BkInstName:            strconv.Itoa(i),
			Createdate:            "",
			DatacenterCode:        strconv.Itoa(i),
			DatacenterName:        fmt.Sprintf("%s%d", "北京亦庄联通数据中心", i),
			DatacenterTypeCode:    "",
			DatacenterTypeName:    "",
			JfgsCode:              "",
			JfgsName:              "",
			LastTime:              time.Time{},
			UnifiedDatacenterCode: "",
			UsageStateNo:          "",
		})
	}
	return mockData
}
