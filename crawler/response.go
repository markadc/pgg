package crawler

import (
	"encoding/json"
	"log"
	"pgg"
	"pgg/utils/py"
)

// Response 响应结构体
type Response struct {
	Content []byte
	Text    string
}

// JSON 获取JSON格式数据
func (r Response) JSON() (pgg.A, error) {
	return py.Loads(r.Text)
}

// JSTF JSON字符串格式化，优化打印显示
func (r Response) JSTF() string {
	jsonData, err := r.JSON()
	if err != nil {
		panic(err)
	}
	bs, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		log.Fatalf("JSON serialization failed: %v", err)
	}
	prettyStr := string(bs)
	return prettyStr
}
