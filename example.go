package main

import (
	"fmt"

	"github.com/465583030/bee/getproxy"
	"github.com/465583030/bee/goreq"
)

func main() {
	resp, body, err := goreq.New().Get("http://www.baidu.com/").End()
	//defer resp
	//getproxy.Get(body)
	fmt.Println(resp)
	fmt.Println(body)
	fmt.Println(err)

	peers := getproxy.Get()
	for _, v := range peers {
		//fmt.Printf("index:%d  value:%d\n", i, v)
		fmt.Println(v)
	}
	//fmt.Println()
}
