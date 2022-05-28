package main

import (
	"flag"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
)

func main() {

	var url string
	flag.StringVar(&url, "url", "", "")
	flag.Parse()

	c := colly.NewCollector(
		colly.AllowURLRevisit(),
	)

	a := colly.NewCollector(
		colly.AllowURLRevisit(),
	)

	rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:7890")
	if err != nil {
		fmt.Println("代理错误~")
	}

	a.SetProxyFunc(rp)
	c.SetProxyFunc(rp)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36")
		r.Headers.Add("Sec-Ch-Viewport-Width", "1680")
		r.Headers.Add(`Sec-Ch-Ua`, `" Not A;Brand";v="99", "Chromium";v="101", "Google Chrome";v="101"`)
		r.Headers.Add("Sec-Ch-Ua-Platform", "macOS")
		r.Headers.Add("Sec-Ch-Ua-Arch", "x86")
		r.Headers.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		r.Headers.Add("X-Client-Data", "CKq1yQEIkbbJAQiktskBCMG2yQEIqZ3KAQjJ0soBCJXdygEIlqHLAQjb78sBCOaEzAEI2qnMAQiKq8wBCMKszAEI36zMAQikr8wBGKupygE=")
		r.Headers.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,ru;q=0.7")
	})

	a.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36")
		r.Headers.Add("Sec-Ch-Viewport-Width", "1680")
		r.Headers.Add(`Sec-Ch-Ua`, `" Not A;Brand";v="99", "Chromium";v="101", "Google Chrome";v="101"`)
		r.Headers.Add("Sec-Ch-Ua-Platform", "macOS")
		r.Headers.Add("Sec-Ch-Ua-Arch", "x86")
		r.Headers.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		r.Headers.Add("X-Client-Data", "CKq1yQEIkbbJAQiktskBCMG2yQEIqZ3KAQjJ0soBCJXdygEIlqHLAQjb78sBCOaEzAEI2qnMAQiKq8wBCMKszAEI36zMAQikr8wBGKupygE=")
		r.Headers.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,ru;q=0.7")
	})

	a.OnHTML("div.NJo7tc.Z26q7c.jGGQ5e > div > a", func(e *colly.HTMLElement) {
		fmt.Println(e.Attr("href"))
	})

	c.OnHTML(".AaVjTc .fl", func(e *colly.HTMLElement) {
		if e.Index == 0 {
			if a.Visit(url) != nil {
				return
			}
		} else {
			if a.Visit(e.Request.AbsoluteURL(e.Attr("href"))) != nil {
				return
			}
		}
	})

	if c.Visit(url) != nil {
		return
	}
}
