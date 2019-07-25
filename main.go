// Project  : mysql2mongo
// package :  main
// file    :  main.go
// mysql2mongo for mysql sync.
// Copyright (c) 2018-2019
// vixtel.com All rights reserved.
package main

import (
	"os"
	"os/signal"
	"syscall"

	"./pkg/logging"

	"./handle"
)

func main() {
	logging.Info("start mysql2mongo")
	handle.Setup()
	handle.Work()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Kill, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sc
	logging.Info("stop mysql2mongo")
}
