package main

import (
	"configcenter/mockTest/models"
	"fmt"
)

/**
 * @Author: Tao Jun
 * @Since: 2022/7/13
 * @Desc: curl.go
**/

const (
	host = "172.22.50.191"
	//corePort = 50009
	apiPort = 8080
	//apiPort  = 31921
	hostPort = 60001
	//hostPort = 32357
	//host = "172.22.50.25"
	//corePort = 32001
)

var (
	header map[string]string = map[string]string{"HTTP_BLUEKING_SUPPLIER_ID": "0", "Content-Type": "application/json"}
)

func main() {
	cm := &models.CreateMany{
		Header: header,
	}
	cm.GetData(10000)
	//res := cm.CreateCoreMany(fmt.Sprintf("http://%s:%d/api/v3/createmany/model/datacenter/instance", host, corePort))
	res := cm.CreateApiHostMany(fmt.Sprintf("http://%s:%d/api/v3/createmany/datacenter", host, apiPort))
	//res := cm.CreateApiHostMany(fmt.Sprintf("http://%s:%d/host/v3/createmany/datacenter", host, hostPort))
	fmt.Println(string(res))
}
