package main

import (
	"gokyrie/cmd"
)

// @title go-web-api
// @version 1.0 版本
// @description 接口描述
func main() {

	defer cmd.Clean()
	cmd.Start()
	// ctx, cancelCtx := context.WithCancel(context.Background())
	// defer cancelCtx()

	// <-ctx.Done()
}
