package loger

import (
	"time"
)

// Now 返回当前时间
func Now() string {
	currTime := time.Now().Format("2006-01-02 15:04:05")
	return currTime
}

// SuppTime 补充时间
func SuppTime(s string, values ...interface{}) string {
	return Format("{}  {}", Now(), Format(s, values...))
}

func Debug(s string, values ...interface{}) {
	PrintWithColor(SuppTime(s, values...), "blue")
}

func Info(s string, values ...interface{}) {
	PrintWithColor(SuppTime(s, values...), "")
}

func Warning(s string, values ...interface{}) {
	PrintWithColor(SuppTime(s, values...), "yellow")
}

func Error(s string, values ...interface{}) {
	PrintWithColor(SuppTime(s, values...), "red")
}

func Success(s string, values ...interface{}) {
	PrintWithColor(SuppTime(s, values...), "green")
}

type Loger struct{}

func (l Loger) Blue(s string, values ...interface{}) {
	PrintWithColor(Format(s, values...), "blue")
}

func (l Loger) Default(s string, values ...interface{}) {
	Print(s, values...)
}

func (l Loger) Yellow(s string, values ...interface{}) {
	PrintWithColor(Format(s, values...), "yellow")
}

func (l Loger) Red(s string, values ...interface{}) {
	PrintWithColor(Format(s, values...), "red")
}

func (l Loger) Green(s string, values ...interface{}) {
	PrintWithColor(Format(s, values...), "green")
}
