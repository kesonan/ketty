# Ketty

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
TODO
