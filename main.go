package main

import (
	"fmt"
	"github.com/CodeAdventurers/ShareDex-BE/initialize"
)

func main() {
	//1. 初始化logger
	initialize.InitLogger()
	//2. 初始化路由
	Router := initialize.Routers()
	go func() {
		if err := Router.Run(fmt.Sprintf(":%d", 8888)); err != nil {
			fmt.Printf("startup service failed, err:%v\n", err)
		}
	}()
}
