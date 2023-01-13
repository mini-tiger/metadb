package common

import (
	"github.com/go-resty/resty"
	"net/http"
)

/**
 * @Author: Tao Jun
 * @Since: 2023/1/3
 * @Desc: http_proxy.go
**/

var Client *resty.Client = resty.New()

type RestyProxyEntry struct {
	headerMap map[string]string
	header    http.Header
	url       string
	body      interface{}
	respBody  interface{}
}

func NewRP(headerMap map[string]string, header http.Header, url string, body interface{}, respBody interface{}) *RestyProxyEntry {
	return &RestyProxyEntry{
		headerMap: headerMap,
		header:    header,
		url:       url,
		body:      body,
		respBody:  respBody,
	}
}

func (rp *RestyProxyEntry) RequestData() (resp *resty.Response, err error) {

	r := Client.R()
	r.Header = rp.header
	//fmt.Printf("%+v\n", header)

	for k, v := range rp.headerMap {
		r.SetHeader(k, v)
	}

	return r.

		//SetHeader("HTTP_BLUEKING_SUPPLIER_ID", "0").
		//SetHeader("Content-Type", "application/json").
		SetBody(rp.body).
		SetResult(rp.respBody).
		Post(rp.url)

}
