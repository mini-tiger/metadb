package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

/**
 * @Author: Tao Jun
 * @Since: 2022/7/13
 * @Desc: http.go
**/
const timeout = -1

func ReqHttp(url string, reqData interface{}, header map[string]string, method string) (error, []byte) {
	client := http.Client{
		Timeout: time.Second * timeout, // Set 10ms timeout.
	}
	//data := make(map[string]interface{})
	//data["name"] = "zhaofan"
	//data["age"] = "23"
	bytesData, _ := json.Marshal(reqData)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	req.Method = method
	for key, value := range header {
		req.Header.Add(key, value)
	}

	resp, _ := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}
	//fmt.Println(string(body))
	return nil, body
}
