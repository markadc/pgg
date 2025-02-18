package loger

import (
	"fmt"
	"strings"
)

// Format 大括号 {} 表示占位
func Format(template string, values ...interface{}) string {
	for _, value := range values {
		template = strings.Replace(template, "{}", fmt.Sprintf("%v", value), 1)
	}
	return template
}

func Print(f string, values ...interface{}) {
	s := Format(f, values...)
	fmt.Println(s)
}

// PrintWithColor 带颜色的输出
func PrintWithColor(content string, color string) {
	colors := map[string]string{
		"black":         "30",
		"red":           "31",
		"green":         "32",
		"yellow":        "33",
		"blue":          "34",
		"magenta":       "35",
		"cyan":          "36",
		"white":         "37",
		"gray":          "90",
		"light_red":     "91",
		"light_green":   "92",
		"light_yellow":  "93",
		"light_blue":    "94",
		"light_magenta": "95",
		"light_cyan":    "96",
		"light_white":   "97",
	}
	colorCode, exists := colors[color]
	if !exists {
		fmt.Println(content)
	} else {
		fmt.Printf("\033[%sm%s\033[0m\n", colorCode, content)
	}
}
