package pulsar

import (
	"github.com/paashzj/gutil"
	"pulsar_mate_go/pkg/config"
	"pulsar_mate_go/pkg/path"
)

func configProxy() error {
	configProp, err := gutil.ConfigPropFromFile(path.PulsarProxyOriginalConfig)
	if err != nil {
		return err
	}
	if config.PulsarTlsEnable {
		configProp.SetBool("useTls", true)
		configProp.Set("pulsarServiceUrl", "pulsar+ssl://"+config.PulsarFunctionBrokerServerHost+":6651")
		configProp.Set("pulsarWebServiceUrl", "https://"+config.PulsarFunctionBrokerWebServiceHost+":8081")
	} else {
		configProp.Set("pulsarServiceUrl", "pulsar://"+config.PulsarFunctionBrokerServerHost+":6650")
		configProp.Set("pulsarWebServiceUrl", "http://"+config.PulsarFunctionBrokerWebServiceHost+":8080")
	}
	configProp.Set("zookeeperServers", config.ZkAddress)
	configProp.Set("configurationStoreServers", config.ZkAddress)
	configProp.Set("clusterName", config.ClusterName)
	configProp.Set("advertisedAddress", config.PulsarAdvertisedAddress)
	return configProp.Write(path.PulsarProxyConfig)
}
