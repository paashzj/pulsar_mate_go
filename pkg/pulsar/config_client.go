package pulsar

import (
	"github.com/paashzj/gutil"
	"pulsar_mate_go/pkg/config"
	"pulsar_mate_go/pkg/path"
	"pulsar_mate_go/pkg/util"
)

func configClient() error {
	util.Logger().Info("begin to generate or refresh client config")
	configProp, err := gutil.ConfigPropFromFile(path.PulsarClientOriginalConfig)
	if err != nil {
		return err
	}
	if config.PulsarTlsEnable {
		configProp.SetBool("useKeyStoreTls", true)
		configProp.Set("tlsTrustStorePath", path.PulsarServerTrustCert)
		configProp.Set("tlsTrustStorePassword", "pulsar_client_pwd")
	}
	configProp.Set("authPlugin", config.PulsarClientAuthPlugin)
	return configProp.Write(path.PulsarClientConfig)
}
