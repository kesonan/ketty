/*
 * MIT License
 *
 * Copyright (c) 2021 anqiansong
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package main

import (
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
	plusStyle := text.WithBorder("=", "=", "-", "|")
	c := console.NewConsole(console.WithTextOption(plusStyle))
	// c.DisableBorder()
	// c.DisableColor()
	c.Info("Hello Ketty")
}
