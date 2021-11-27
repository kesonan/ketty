# Ketty
[![Go](https://github.com/anqiansong/ketty/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/anqiansong/ketty/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/anqiansong/ketty/branch/main/graph/badge.svg?token=KAMZX9OEYF)](https://codecov.io/gh/anqiansong/ketty)
[![Go Report Card](https://goreportcard.com/badge/github.com/anqiansong/ketty)](https://goreportcard.com/report/github.com/anqiansong/ketty)
[![License](https://img.shields.io/github/license/anqiansong/ketty)](https://github.com/anqiansong/ketty/blob/main/LICENSE)


[中文](README.md)|English

ketty is a Logger output with pretty and colorful log of Golang 。

## Installation

```bash
$ go install github.com/anqiansong/ketty@latest
```

## Quick Start

```go
func main(){
    console.Info(`
    {
        "name":"Hello Ketty",
        "description":"a color logger",
        "author":"anqiansong",
        "category":"console",
        "github":"https://github.com/anqiansong/ketty",
        "useage":[
            "info",
            "debug"
        ]
    }`)
    console.Debug("Hello Ketty")
    console.Warn("Hello Ketty")
    console.Error(errors.New("error test"))
}
```

## Output

### Terminal
![terminal](./resource/terminal.png)

### Idea
![idea1](./resource/idea1.png)
![idea1](./resource/idea2.png)

## Usage
### Use Directly
There are some default options if you use directly:
* the border style is `frame.WithLineStyle`
* colorful output is opened by default

```go
func main(){
    console.Info("Hello ketty, This is info log")
    console.Debug("Hello ketty, This debug log")
    console.Warn("Hello ketty, This warn log")
    console.Error(errors.New("Hello ketty,This is an error"))
}
```

### Initialization
```go
    // Use Plus-Style as border.
    plusStyle := text.WithPlusStyle()
    c := console.NewConsole(console.WithTextOption(plusStyle))
```

### Console Configure
```go
    c.DisableBorder() // disable border
    c.DisableColor() // disable colorful output
```

### Print
```go
    // Print info log
    c.Info("Hello Ketty, It's now %q", time.Now())
```

## Border
### Preset
* WithLineStyle Defaul
```text
[INFO] 2021-11-26 22:29:14.508 ┌────────────────────────────────────────────────────────────────────────────
[INFO] 2021-11-26 22:29:14.508 │  Hello Ketty, It's now "2021-11-26 22:29:14.508085 +0800 CST m=+0.000229190"
[INFO] 2021-11-26 22:29:14.508 └────────────────────────────────────────────────────────────────────────────
```
* WithDotStyle
```text
[INFO] 2021-11-26 22:30:22.913 .............................................................................
[INFO] 2021-11-26 22:30:22.913 .  Hello Ketty, It's now "2021-11-26 22:30:22.913678 +0800 CST m=+0.000199931"
[INFO] 2021-11-26 22:30:22.913 .............................................................................
```

* WithStarStyle
```text
[INFO] 2021-11-26 22:31:00.699 *****************************************************************************
[INFO] 2021-11-26 22:31:00.699 *  Hello Ketty, It's now "2021-11-26 22:31:00.699094 +0800 CST m=+0.000186578"
[INFO] 2021-11-26 22:31:00.699 *****************************************************************************
```

* WithPlusStyle
```text
[INFO] 2021-11-26 22:31:26.952 +----------------------------------------------------------------------------
[INFO] 2021-11-26 22:31:26.952 |  Hello Ketty, It's now "2021-11-26 22:31:26.952376 +0800 CST m=+0.000168647"
[INFO] 2021-11-26 22:31:26.952 +----------------------------------------------------------------------------
```

* WithFivePointedStarStyle
```text
[INFO] 2021-11-26 22:31:58.146 ★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★
[INFO] 2021-11-26 22:31:58.146 ★  Hello Ketty, It's now "2021-11-26 22:31:58.146534 +0800 CST m=+0.000171850"
[INFO] 2021-11-26 22:31:58.146 ★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★
```

* WithDoubleLine
```text
[INFO] 2021-11-26 22:32:21.574 ╔════════════════════════════════════════════════════════════════════════════
[INFO] 2021-11-26 22:32:21.574 ║  Hello Ketty, It's now "2021-11-26 22:32:21.573911 +0800 CST m=+0.000152572"
[INFO] 2021-11-26 22:32:21.574 ╚════════════════════════════════════════════════════════════════════════════
```

* DisableBorder
```text
[INFO] 2021-11-26 22:33:01.695   Hello Ketty, It's now "2021-11-26 22:33:01.695338 +0800 CST m=+0.000156150"
```

### Custom Style
* WithCommonBorder
```go
// The all direction use the same character
plusStyle := text.WithCommonBorder("x")
```
```text
[INFO] 2021-11-26 22:34:01.437 xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
[INFO] 2021-11-26 22:34:01.437 x  Hello Ketty, It's now "2021-11-26 22:34:01.437286 +0800 CST m=+0.000153825"
[INFO] 2021-11-26 22:34:01.437 xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

* WithBorder
```go
    // arg1: the character in the left-top
    // arg2: the character in the left-bottom
    // arg3: the character in horizontal
    // arg4: the character in vertical
    plusStyle := text.WithBorder("=","=","-","|")
    c := console.NewConsole(console.WithTextOption(plusStyle))
```
```text
[INFO] 2021-11-26 22:37:38.409 =----------------------------------------------------------------------------
[INFO] 2021-11-26 22:37:38.409 |  Hello Ketty, It's now "2021-11-26 22:37:38.408952 +0800 CST m=+0.000155037"
[INFO] 2021-11-26 22:37:38.409 =----------------------------------------------------------------------------
```

## Persistence
TODO

## Notes
The colorful output does not support on Windows.
