package main

import (
	"go.uber.org/zap"
	"os"
	"os/signal"
	"pulsar_mate_go/pkg/config"
	"pulsar_mate_go/pkg/pulsar"
	"pulsar_mate_go/pkg/util"
)

func main() {
	util.Logger().Debug("this is a debug msg")
	util.Logger().Info("this is a info msg")
	util.Logger().Error("this is a error msg")
	err := pulsar.Config()
	if err != nil {
		util.Logger().Error("generate config failed ", zap.Error(err))
	}
	if config.ClusterInit {
		pulsar.StartInit()
	} else {
		pulsar.StartOther()
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			return
		}
	}
}
