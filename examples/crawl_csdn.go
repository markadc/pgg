package main

import (
	"encoding/json"
	"fmt"
	"pgg"
	"pgg/requests"
	"pgg/utils/loger"
)

// 定义与 JSON 结构匹配的结构体
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	List  []Item `json:"list"`
	Total *int   `json:"total"`
}

type Item struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func gotoMarkadc() {
	url := "https://blog.csdn.net/community/home-api/v1/get-business-list?page=1&size=100&businessType=lately&noMore=false&username=MarkAdc"
	headers := pgg.S{
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36 Edg/133.0.0.0",
		"Referer":    "https://blog.csdn.net/MarkAdc",
	}
	resp, err := requests.Get(url, headers, nil)
	if err != nil {
		fmt.Println(err)
	}
	var response Response
	err = json.Unmarshal(resp, &response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.Data.List, "\n")
	for _, item := range response.Data.List {
		fmt.Printf("Title: %s\nURL: %s\n\n", item.Title, item.URL)
	}
	if len(response.Data.List) > 0 {
		loger.Success("一共搜索到了 {} 条文章\n", len(response.Data.List))
	} else {
		loger.Warning("没有搜索到文章")
	}
}

func main() {
	gotoMarkadc()
}
