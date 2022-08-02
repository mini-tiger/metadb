package models

import (
	"configcenter/mockTest/utils"
	"log"

	"fmt"
	"time"
)

/**
 * @Author: Tao Jun
 * @Since: 2022/7/14
 * @Desc: createMany.go
**/
type CreateMany struct {
	Data   []*Instance
	Header map[string]string
}

func (c *CreateMany) GetData(num int) {
	c.Data = GetMockInstance(num)
}

func (c *CreateMany) CreateCoreMany(url string) []byte {
	start := time.Now()
	defer func() {
		if err := recover(); err != nil {
			log.Printf("recover request Time:%f err:%v\n", time.Now().Sub(start).Seconds(), err)
		}
	}()

	fmt.Println(url)

	//fmt.Println(data)
	reqParams := map[string]interface{}{"datas": c.Data}
	_, byteStr := utils.ReqHttp(url, reqParams, c.Header, "POST")
	log.Printf("request Time:%f\n", time.Now().Sub(start).Seconds())
	return byteStr

}

func (c *CreateMany) CreateApiHostMany(url string) []byte {
	start := time.Now()
	defer func() {
		if err := recover(); err != nil {
			log.Printf("recover request Time:%f err:%v\n", time.Now().Sub(start).Seconds(), err)
		}
	}()

	fmt.Println(url)

	//fmt.Println(data)
	reqParams := map[string]interface{}{"data": c.Data}
	_, byteStr := utils.ReqHttp(url, reqParams, c.Header, "POST")
	log.Printf("request Time:%f\n", time.Now().Sub(start).Seconds())
	//fmt.Println(string(res))
	return byteStr
}
