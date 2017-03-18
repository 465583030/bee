// example
package main

import (
	"fmt"

	"github.com/465583030/bee/fetcher"
)

func main() {
	resp, body, err := fetcher.NewFetcher("douban.com").Get("/")
	if err != nil {
		panic(err)
	}
	fmt.Printf("status:%d", resp.StatusCode)
	fmt.Println("body:\n%v", string(body))
}
