package logMiddleWare

import (
	"bytes"
	"configcenter/src/common/blog"
	"io/ioutil"
	"net/http"
)

/**
 * @Author: Tao Jun
 * @Since: 2022/7/11
 * @Desc: log.go
**/

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			blog.Infof("Access: [%s] method: [%s] params %v", r.RequestURI, r.Method, r.URL.Query())
		} else {
			//blog.Infof("Access: [%s] method: [%s] ", r.RequestURI, r.Method)
			data, _ := ioutil.ReadAll(r.Body) //获取post的数据
			r.Body = ioutil.NopCloser(bytes.NewBuffer(data))
			//fmt.Println(string(con))
			blog.Infof("Access: [%s] method: [%s] params %v", r.RequestURI, r.Method, string(data))
		}
		next.ServeHTTP(w, r)
	})

}
