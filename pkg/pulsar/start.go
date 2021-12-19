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
	util.Logger().Info("shell result ", zap.String("stdout", stdout), zap.String("stderr", stderr))
	if err != nil {
		util.Logger().Error("pulsar cluster init failed ", zap.Error(err))
	} else {
		os.Exit(0)
	}
}

func StartOther() {
	if config.ClusterEnable {
		if config.Role == "all" {
			startScript(path.PulsarStartClusterAllScript)
		} else if config.Role == "broker" {
			startScript(path.PulsarStartClusterBrokerScript)
		} else if config.Role == "function" {
			startScript(path.PulsarStartClusterFunctionScript)
		} else if config.Role == "proxy" {
			startScript(path.PulsarStartClusterProxyScript)
		}
	} else {
		if config.Role == "all" {
			startScript(path.PulsarStartStandaloneAllScript)
		} else if config.Role == "broker" {
			startScript(path.PulsarStartStandaloneBrokerScript)
		}
	}
}

func startScript(script string) {
	stdout, stderr, err := gutil.CallScript(script)
	util.Logger().Info("shell result ", zap.String("stdout", stdout), zap.String("stderr", stderr))
	if err != nil {
		util.Logger().Error("start pulsar server failed ", zap.Error(err))
		os.Exit(1)
	}
}
