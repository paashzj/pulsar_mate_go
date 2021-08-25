package pulsar

import (
	"github.com/paashzj/gutil"
	"go.uber.org/zap"
	"os"
	"pulsar_mate_go/pkg/config"
	"pulsar_mate_go/pkg/path"
	"pulsar_mate_go/pkg/util"
)

func StartInit() {
	stdout, stderr, err := gutil.CallScript(path.PulsarInitScript)
	util.Logger().Error("shell result ", zap.String("stdout", stdout), zap.String("stderr", stderr))
	if err != nil {
		util.Logger().Error("pulsar cluster init failed ", zap.Error(err))
	} else {
		os.Exit(0)
	}
}

func StartOther() {
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
}
