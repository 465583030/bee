package main

import (
	"fmt"

	"github.com/465583030/bee/getproxy"
	"github.com/465583030/bee/goreq"
)

func main() {
	resp, body, err := goreq.New().Get("http://www.baidu.com/").End()

	fmt.Println(resp)
	fmt.Println(body)
	fmt.Println(err)

	peers := getproxy.Get()
	for i, v := range peers {
		if checked := v.Check(); checked > 0 {
			fmt.Printf("%d : (ip:%s, port:%s, type:%s,status:%d)\n", i, v.Ip, v.Port, v.Proto, v.Status)
		}
	}
}
