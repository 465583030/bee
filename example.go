package main

import (
	"github.com/465583030/bee/getproxy"
	"github.com/465583030/bee/goreq"
)

func main() {
	resp, body, err := goreq.New().Get("http://cn-proxy.com/").end
	//defer resp
	getproxy.Parse_cn_proxy(body)
}
