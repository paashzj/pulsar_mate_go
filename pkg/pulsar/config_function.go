package pulsar

import (
	"github.com/paashzj/gutil"
	"pulsar_mate_go/pkg/config"
	"pulsar_mate_go/pkg/path"
)

func configFunction() error {
	configYaml, err := gutil.ConfigYamlFromFile(path.PulsarFunctionOriginalConfig)
	if err != nil {
		return err
	}
	if config.PulsarTlsEnable {
		configYaml.SetBool("useTls", true)
		configYaml.Set("pulsarServiceUrl", "pulsar+ssl://"+config.PulsarFunctionBrokerServerHost+":6651")
		configYaml.Set("pulsarWebServiceUrl", "https://"+config.PulsarFunctionBrokerWebServiceHost+":8081")
	} else {
		configYaml.Set("pulsarServiceUrl", "pulsar://"+config.PulsarFunctionBrokerServerHost+":6650")
		configYaml.Set("pulsarWebServiceUrl", "http://"+config.PulsarFunctionBrokerWebServiceHost+":8080")
	}
	return configYaml.Write(path.PulsarFunctionConfig)
}
