package pulsar

import (
	"github.com/paashzj/gutil"
	"pulsar_mate_go/pkg/config"
	"pulsar_mate_go/pkg/path"
	"pulsar_mate_go/pkg/util"
)

func configBrokerStandaloneWithFunction() error {
	util.Logger().Info("begin to generate or refresh broker standalone config")
	configProp, err := gutil.ConfigPropFromFile(path.PulsarStandaloneOriginalConfig)
	if err != nil {
		return err
	}
	configBrokerWithFunctionCommon(configProp)
	return configProp.Write(path.PulsarStandaloneConfig)
}

func configBrokerClusterWithFunction() error {
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
	configBrokerWithFunctionCommon(configProp)
	return configProp.Write(path.PulsarConfig)
}

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

func configBrokerWithFunctionCommon(prop *gutil.ConfigProperties) {
	configBrokerCommon(prop)
}

func configBrokerCommon(prop *gutil.ConfigProperties) {
	prop.Set("advertisedAddress", config.PulsarAdvertisedAddress)
	prop.SetBool("authenticationEnabled", config.PulsarAuthenticationEnabled)
	prop.Set("authenticationProviders", config.PulsarAuthenticationProviders)
	prop.SetBool("authorizationEnabled", config.PulsarAuthorizationEnabled)
	prop.Set("tokenSecretKey", config.PulsarTokenSecretKey)
	prop.Set("superUserRoles", config.PulsarSuperUserRoles)
	prop.Set("brokerClientAuthenticationPlugin", config.PulsarClientAuthPlugin)
	prop.SetBool("allowAutoTopicCreation", config.PulsarAllowAutoTopicCreation)
	prop.SetBool("brokerDeleteInactivePartitionedTopicMetadataEnabled", config.PulsarBrokerDeleteInactivePartitionedTopicMetadataEnabled)
	prop.SetBool("tlsAllowInsecureConnection", config.PulsarTlsAllowInsecureConnection)
	if config.BkTlsEnable {
		prop.Set("bookkeeperTLSKeyFilePath", path.BkClientKeyCert)
		prop.Set("bookkeeperTLSKeyStorePasswordPath", path.BkClientKeyPassword)
		prop.Set("bookkeeperTLSTrustCertsFilePath", path.BkClientTrustCert)
		prop.Set("bookkeeperTLSTrustStorePasswordPath", path.BkClientTrustPassword)
	}
	if config.PulsarTlsEnable {
		// client
		prop.SetBool("brokerClientTlsEnabledWithKeyStore", true)
		prop.Set("brokerClientTlsTrustStore", path.PulsarClientTrustCert)
		prop.Set("brokerClientTlsTrustStorePassword", "pulsar_client_pwd")
		// server
		prop.SetInt("brokerServicePortTls", 6651)
		prop.SetInt("webServicePortTls", 8081)
		prop.Set("tlsKeyStore", path.PulsarServerKeyCert)
		prop.Set("tlsKeyStorePassword", "pulsar_server_pwd")
		prop.Set("tlsTrustStore", path.PulsarServerTrustCert)
		prop.Set("tlsTrustStorePassword", "pulsar_server_pwd")
	}
}
