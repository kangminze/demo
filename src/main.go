// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

//import "demo/src/cmd"

import (
	_ "demo/docs"
	"demo/src/dao"
	"demo/src/log"
	"demo/src/router"
	"demo/src/work"
	"github.com/gin-gonic/gin"
	"github.com/sevenNt/wzap"
	"os"
	"os/signal"
	"syscall"
)

func start() {
	//初始化日志
	log.Init()
	//初始化数据库
	dao.Init()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGEMT, syscall.SIGKILL)

	r := gin.Default()
	router.Init(r)

	//init work
	go work.Start()

	//当有信号发出中断信息时close channel
	go func() {
		<-quit
		work.Stop()
		wzap.Info("close work job channel")
	}()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
	//cmd.Execute()
}

func main() {
	start()
}
