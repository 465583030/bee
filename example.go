package main

import (
	"465583030/getproxy"
	"465583030/goreq"
)

func main() {
	resp, body, err := goreq.New().Get("http://cn-proxy.com/").end
	//defer resp
	getproxy.Parse_cn_proxy(body)
}
