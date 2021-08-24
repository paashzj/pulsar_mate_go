package pulsar

import (
	"pulsar_mate_go/pkg/config"
)

func configFunction() error {
	if config.ClusterEnable {
		return configFunctionCluster()
	}
	return configFunctionStandalone()
}

func configFunctionCluster() error {
	// todo
	return nil
}

func configFunctionStandalone() error {
	// todo
	return nil
}
