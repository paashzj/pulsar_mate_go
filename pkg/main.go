package main

import (
	"github.com/paashzj/gutil"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"pulsar_mate_go/pkg/config"
	"pulsar_mate_go/pkg/path"
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
		stdout, stderr, err := gutil.CallScript(path.PulsarInitScript)
		util.Logger().Error("shell result ", zap.String("stdout", stdout), zap.String("stderr", stderr))
		if err != nil {
			util.Logger().Error("pulsar server init failed ", zap.Error(err))
		} else {
			os.Exit(0)
		}
	}
	if config.Function {
		stdout, stderr, err := gutil.CallScript(path.PulsarStartFunctionScript)
		util.Logger().Error("shell result ", zap.String("stdout", stdout), zap.String("stderr", stderr))
		if err != nil {
			util.Logger().Error("start pulsar function server failed ", zap.Error(err))
			os.Exit(1)
		}
	}
	if config.ClusterEnable {
		stdout, stderr, err := gutil.CallScript(path.PulsarStartScript)
		util.Logger().Error("shell result ", zap.String("stdout", stdout), zap.String("stderr", stderr))
		if err != nil {
			util.Logger().Error("start pulsar server failed ", zap.Error(err))
			os.Exit(1)
		}
	} else {
		stdout, stderr, err := gutil.CallScript(path.PulsarStartStandaloneScript)
		util.Logger().Error("shell result ", zap.String("stdout", stdout), zap.String("stderr", stderr))
		if err != nil {
			util.Logger().Error("start pulsar server failed ", zap.Error(err))
			os.Exit(1)
		}
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
