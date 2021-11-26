package main

import (
	"github.com/anqiansong/ketty/console"
	"github.com/pkg/errors"
)

func main() {
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
