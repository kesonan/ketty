package main

import (
	"time"

	"github.com/anqiansong/ketty/console"
	"github.com/anqiansong/ketty/text"
)

func main() {
	// 快速开始
	// 	console.Info(`
	// {
	//     "name":"Hello Ketty",
	//     "description":"a color logger",
	//     "author":"anqiansong",
	//     "category":"console",
	//     "github":"https://github.com/anqiansong/ketty",
	//     "useage":[
	//         "info",
	//         "debug"
	//     ]
	// }`)
	// 	console.Debug("Hello Ketty")
	// 	console.Warn("Hello Ketty")
	// 	console.Error(errors.New("error test"))

	// 直接使用
	// console.Info("Hello ketty, This is info log")
	// console.Debug("Hello ketty, This debug log")
	// console.Warn("Hello ketty, This warn log")
	// console.Error(errors.New("Hello ketty,This is an error"))

	// 初始化
	plusStyle := text.WithBorder("=","=","-","|")
	c := console.NewConsole(console.WithTextOption(plusStyle))
	// c.DisableBorder()
	// c.DisableColor()
	c.Info("Hello Ketty, It's now %q", time.Now())
}
