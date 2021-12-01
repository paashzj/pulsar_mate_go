package path

import (
	"os"
	"path/filepath"
)

// pulsar
var (
	PulsarHome                     = os.Getenv("PULSAR_HOME")
	PulsarConfigDir                = filepath.FromSlash(PulsarHome + "/conf")
	PulsarConfig                   = filepath.FromSlash(PulsarConfigDir + "/broker.conf")
	PulsarOriginalConfig           = filepath.FromSlash(PulsarConfigDir + "/broker_original.conf")
	PulsarStandaloneConfig         = filepath.FromSlash(PulsarConfigDir + "/standalone.conf")
	PulsarStandaloneOriginalConfig = filepath.FromSlash(PulsarConfigDir + "/standalone_original.conf")
	PulsarFunctionConfig           = filepath.FromSlash(PulsarConfigDir + "/functions_worker.yml")
	PulsarFunctionOriginalConfig   = filepath.FromSlash(PulsarConfigDir + "/functions_worker_original.yml")
	PulsarClientConfig             = filepath.FromSlash(PulsarConfigDir + "/client.conf")
	PulsarClientOriginalConfig     = filepath.FromSlash(PulsarConfigDir + "/client_original.conf")
)

// mate
var (
	PulsarMatePath                    = filepath.FromSlash(PulsarHome + "/mate")
	PulsarScripts                     = filepath.FromSlash(PulsarMatePath + "/scripts")
	PulsarInitScript                  = filepath.FromSlash(PulsarScripts + "/init-pulsar.sh")
	PulsarStartClusterAllScript       = filepath.FromSlash(PulsarScripts + "/start-pulsar-cluster-all.sh")
	PulsarStartClusterBrokerScript    = filepath.FromSlash(PulsarScripts + "/start-pulsar-cluster-broker.sh")
	PulsarStartClusterFunctionScript  = filepath.FromSlash(PulsarScripts + "/start-pulsar-cluster-function.sh")
	PulsarStartStandaloneAllScript    = filepath.FromSlash(PulsarScripts + "/start-pulsar-standalone-all.sh")
	PulsarStartStandaloneBrokerScript = filepath.FromSlash(PulsarScripts + "/start-pulsar-standalone-broker.sh")
	PulsarCerts                       = filepath.FromSlash(PulsarMatePath + "/cert")
	ZkClientKeyCert                   = filepath.FromSlash(PulsarCerts + "/zk_client_key.jks")
	ZkClientTrustCert                 = filepath.FromSlash(PulsarCerts + "/zk_client_trust.jks")
	BkClientKeyCert                   = filepath.FromSlash(PulsarCerts + "/bk_client_key.jks")
	BkClientKeyPassword               = filepath.FromSlash(PulsarCerts + "/bk_client_key.passwd")
	BkClientTrustCert                 = filepath.FromSlash(PulsarCerts + "/bk_client_trust.jks")
	BkClientTrustPassword             = filepath.FromSlash(PulsarCerts + "/bk_client_trust.passwd")
	PulsarClientKeyCert               = filepath.FromSlash(PulsarCerts + "/pulsar_client_key.jks")
	PulsarClientTrustCert             = filepath.FromSlash(PulsarCerts + "/pulsar_client_trust.jks")
	PulsarServerKeyCert               = filepath.FromSlash(PulsarCerts + "/pulsar_server_key.jks")
	PulsarServerTrustCert             = filepath.FromSlash(PulsarCerts + "/pulsar_server_trust.jks")
	PulsarFunctionClientKeyCert       = filepath.FromSlash(PulsarCerts + "/pulsar_function_client_key.jks")
	PulsarFunctionClientTrustCert     = filepath.FromSlash(PulsarCerts + "/pulsar_function_client_trust.jks")
	PulsarFunctionServerKeyCert       = filepath.FromSlash(PulsarCerts + "/pulsar_function_server_key.jks")
	PulsarFunctionServerTrustCert     = filepath.FromSlash(PulsarCerts + "/pulsar_function_server_trust.jks")
)
