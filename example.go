package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/465583030/bee/useragent"
	"github.com/PuerkitoBio/goquery"
)

//"github.com/465583030/bee/getproxy"
//"github.com/465583030/bee/goreq"

func main() {
	/*gr := goreq.New()
	resp, body, err := gr.SetHeader("User-Agent", useragent.GetRandomUserAgent()).Get("http://www.baidu.com/").End()
	fmt.Println(resp.Header)
	fmt.Println(body)
	fmt.Println(err)*/

	/*request := gorequest.New()
	resp, body, _ := request.Get("http://movie.douban.com/subject/2035218/?from=tag_all").End()
	defer resp.Body.Close()
	fmt.Println(body)*/
	/*peers := getproxy.Get()
	for i, v := range peers {
		if checked := v.Check(); checked > 0 {
			fmt.Printf("%d : (ip:%s, port:%s, type:%s, status:%d)\n", i, v.Ip, v.Port, v.Proto, v.Status)
		}
	}*/
	douban()
	xicidaili()
}

func douban() {
	resp, err := Get("http://movie.douban.com/subject/25850640/")
	CheckError(err)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	pipe := PipeItem{}
	err = json.Unmarshal([]byte(`
		{
			"type": "map",
			"selector": "",
			"subitem": [
						{
							"type": "text",
							"selector": "title",
							"name": "name",
							"filter": "trim(\n)|replace((豆瓣))|trim( )"
						},
						{
							"type": "attr[data-type]",
							"selector": "#content .gtleft a.bn-sharing",
							"name": "fenlei"
						},
						{
							"type": "attr[data-pic]",
							"selector": "#content .gtleft a.bn-sharing",
							"name": "thumbnail"
						},
						{
							"type": "text-array",
							"selector": "#info span.attrs a[rel=v\\:directedBy]",
							"name": "direct"
						},
						{
							"type": "text-array",
							"selector": "#info span a[rel=v\\:starring]",
							"name": "starring"
						},
						{
							"type": "text-array",
							"selector": "#info span[property=v\\:genre]",
							"name": "type"
						},
						{
							"type": "attr-array[src]",
							"selector": "#related-pic .related-pic-bd a:not(.related-pic-video) img",
							"name": "imgs",
							"filter": "join($)|replace(albumicon,photo)|split($)"
						},
						{
							"type": "text-array",
							"selector": "#info span[property=v\\:initialReleaseDate]",
							"name": "releasetime"
						},
						{
							"type": "text",
							"selector": "#info span[property=v\\:runtime]",
							"name": "longtime"
						},
						{
							"type": "text",
							"selector": "regexp:<span class=\"pl\">制片国家/地区:</span> ([\\w\\W]+?)<br/>",
							"name": "country",
							"filter": "split(/)|trim( )"
						},
						{
							"type": "text",
							"selector": "regexp:<span class=\"pl\">语言:</span> ([\\w\\W]+?)<br/>",
							"name": "language",
							"filter": "split(/)|trim( )"
						},
						{
							"type": "text",
							"selector": "regexp:<span class=\"pl\">集数:</span> (\\d+)<br/>",
							"name": "episode",
							"filter": "intval()"
						},
						{
							"type": "text",
							"selector": "regexp:<span class=\"pl\">又名:</span> ([\\w\\W]+?)<br/>",
							"name": "alias",
							"filter": "split(/)|trim( )"
						},
				    		{
				    			"type": "text",
							"selector": "#link-report span.hidden, #link-report span[property=v\\:summary]|last",
				    			"name": "brief",
				    			"filter": "trim(\n )|split(\n)|trim( )|wraphtml(p)|join()"
				    		},
						{
							"type": "text",
							"selector": "#interest_sectl .rating_num",
							"name": "score",
							"filter": "floatval()"
						},
						{
							"type": "text",
							"selector": "#content h1 span.year",
							"name": "year",
							"filter": "replace(()|replace())|intval()"
						},
						{
							"type": "text",
							"selector": "#comments-section > .mod-hd h2 a",
							"name": "comment",
							"filter": "replace(全部)|replace(条)|trim( )|intval()"
						}
			]
		}
	`), &pipe)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(pipe.PipeBytes(body, "html"))
}

func xicidaili() {
	proxys := make([]string, 0)
	resp, err := Get("http://www.xicidaili.com/nn/")
	CheckError(err)
	defer resp.Body.Close()
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	//fmt.Println(doc.Html())
	doc.Find("#ip_list tbody tr").Each(func(i int, s *goquery.Selection) {
		ip := s.Find("td").Eq(1).Text()
		port := s.Find("td").Eq(2).Text()
		if ip != "" && port != "" {
			//proxy[i] = fmt.Sprintf("http://%s:%s", ip, port)
			proxys = append(proxys, fmt.Sprintf("http://%s:%s", ip, port))
		}
	})
	for i, v := range proxys {
		fmt.Printf("%d: %v\n", i, v)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Add("Accept-Language", "en-US,en;q=0.8")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("User-Agent", useragent.GetRandomUserAgent())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return resp, nil
}
