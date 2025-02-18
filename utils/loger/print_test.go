package loger

import "testing"

func TestPrint(t *testing.T) {
	s1 := Format("当前时间是：{}", Now())
	PrintWithColor(s1, "blue")

	name, code := "王天风", "毒蜂"
	s2 := Format("姓名：{}，代号：{}", name, code)
	PrintWithColor(s2, "red")
}
