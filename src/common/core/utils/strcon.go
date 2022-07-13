package utils

import (
	"encoding/json"
	"reflect"
	"strings"
)

/**
 * @Author: Tao Jun
 * @Since: 2022/7/7
 * @Desc: strcon.go
**/
func JsonToMap(s string) (resultMap map[string]interface{}, err error) {
	//marshal, _ := json.Marshal([]byte(s))
	resultMap = make(map[string]interface{}, 0)
	//before := time.Now()
	err = json.Unmarshal([]byte(s), &resultMap)
	return
	//fmt.Println(time.Since(before))
	//fmt.Printf("%+v", resultMap)
}

func StructToMapByReflect(s interface{}) map[string]interface{} {
	var resultMap = make(map[string]interface{}, 10)
	//before := time.Now()

	ty := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		resultMap[strings.ToLower(ty.Field(i).Name)] = v.Field(i).Interface()
	}
	//fmt.Println(time.Since(before))
	//fmt.Printf("%+v",resultMap)
	return resultMap
}
