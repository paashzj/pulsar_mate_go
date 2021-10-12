package pulsar

import (
	"github.com/paashzj/gutil"
	"pulsar_mate_go/pkg/config"
	"pulsar_mate_go/pkg/path"
	"pulsar_mate_go/pkg/util"
)

func configBrokerStandalone() error {
	util.Logger().Info("begin to generate or refresh broker standalone config")
	configProp, err := gutil.ConfigPropFromFile(path.PulsarStandaloneOriginalConfig)
	if err != nil {
		return err
	}
	configBrokerCommon(configProp)
	return configProp.Write(path.PulsarStandaloneConfig)
}

func configBrokerCluster() error {
	configProp, err := gutil.ConfigPropFromFile(path.PulsarOriginalConfig)
	if err != nil {
		return err
	}
	ipv4Addr, err := GetInterfaceIpv4Addr("eth0")
	if err != nil {
		return err
	}
	configProp.Set("advertisedAddress", ipv4Addr)
	configProp.Set("zookeeperServers", config.ZkAddress)
	configProp.Set("configurationStoreServers", config.ZkAddress)
	configProp.Set("clusterName", config.ClusterName)
	configProp.Set("allowAutoTopicCreationType", "partitioned")
	configBrokerCommon(configProp)
	return configProp.Write(path.PulsarConfig)
}

func configBrokerCommon(prop *gutil.ConfigProperties) {
	prop.SetBool("authenticationEnabled", config.PulsarAuthenticationEnabled)
	prop.Set("authenticationProviders", config.PulsarAuthenticationProviders)
	prop.SetBool("authorizationEnabled", config.PulsarAuthorizationEnabled)
	prop.Set("tokenSecretKey", config.PulsarTokenSecretKey)
	prop.Set("superUserRoles", config.PulsarSuperUserRoles)
	prop.Set("brokerClientAuthenticationPlugin", config.PulsarClientAuthPlugin)
	prop.SetBool("allowAutoTopicCreation", config.PulsarAllowAutoTopicCreation)
	prop.SetBool("brokerDeleteInactivePartitionedTopicMetadataEnabled", config.PulsarBrokerDeleteInactivePartitionedTopicMetadataEnabled)
}
