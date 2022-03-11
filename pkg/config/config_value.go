package config

import "github.com/paashzj/gutil"

// pulsar config
var (
	ClusterEnable                                             bool
	ClusterInit                                               bool
	Role                                                      string
	ZkAddress                                                 string
	ClusterName                                               string
	Spec                                                      string
	PulsarAdvertisedAddress                                   string
	PulsarAuthenticationEnabled                               bool
	PulsarAuthenticationProviders                             string
	PulsarAuthorizationEnabled                                bool
	PulsarTokenSecretKey                                      string
	PulsarSuperUserRoles                                      string
	PulsarClientAuthPlugin                                    string
	PulsarAllowAutoTopicCreation                              bool
	PulsarBrokerDeleteInactivePartitionedTopicMetadataEnabled bool
	PulsarBrokerEntryMetadataInterceptors                     string
	PulsarBrokerDeleteInactiveTopicsEnabled                   bool
	PulsarTlsAllowInsecureConnection                          bool
	PulsarExposingBrokerEntryMetadataToClientEnabled          bool
	PulsarFunctionBrokerServerHost                            string
	PulsarFunctionBrokerWebServiceHost                        string
	ZkTlsEnable                                               bool
	BkTlsEnable                                               bool
	PulsarTlsEnable                                           bool
	PulsarFunctionTlsEnable                                   bool
)

func init() {
	ClusterEnable = gutil.GetEnvBool("CLUSTER_ENABLE", false)
	ClusterInit = gutil.GetEnvBool("CLUSTER_INIT", false)
	Role = gutil.GetEnvStr("PULSAR_ROLE", "all")
	ZkAddress = gutil.GetEnvStr("ZK_ADDR", "")
	ClusterName = gutil.GetEnvStr("CLUSTER_NAME", "pulsar")
	Spec = gutil.GetEnvStr("SPEC", "SMALL")
	PulsarAdvertisedAddress = gutil.GetEnvStr("PULSAR_ADVERTISED_ADDRESS", "")
	PulsarAuthenticationEnabled = gutil.GetEnvBool("PULSAR_AUTHENTICATION_ENABLED", false)
	PulsarAuthenticationProviders = gutil.GetEnvStr("PULSAR_AUTHENTICATION_PROVIDERS", "")
	PulsarAuthorizationEnabled = gutil.GetEnvBool("PULSAR_AUTHORIZATION_ENABLED", false)
	PulsarTokenSecretKey = gutil.GetEnvStr("PULSAR_TOKEN_SECRET_KEY", "")
	PulsarSuperUserRoles = gutil.GetEnvStr("PULSAR_SUPER_USER_ROLES", "")
	PulsarClientAuthPlugin = gutil.GetEnvStr("PULSAR_CLIENT_AUTH_PLUGIN", "")
	PulsarAllowAutoTopicCreation = gutil.GetEnvBool("PULSAR_ALLOW_AUTO_TOPIC_CREATION", true)
	PulsarBrokerDeleteInactivePartitionedTopicMetadataEnabled = gutil.GetEnvBool("PULSAR_BROKER_DELETE_INACTIVE_PARTITIONED_TOPIC_METADATA_ENABLED", false)
	PulsarBrokerEntryMetadataInterceptors = gutil.GetEnvStr("PULSAR_BROKER_ENTRY_METADATA_INTERCEPTORS", "")
	PulsarBrokerDeleteInactiveTopicsEnabled = gutil.GetEnvBool("PULSAR_BROKER_DELETE_INACTIVE_TOPICS_ENABLED", false)
	PulsarTlsAllowInsecureConnection = gutil.GetEnvBool("PULSAR_TLS_ALLOW_INSECURE_CONNECTION", false)
	PulsarExposingBrokerEntryMetadataToClientEnabled = gutil.GetEnvBool("PULSAR_EXPOSING_BROKER_ENTRY_METADATA_TO_CLIENT_ENABLED", false)
	PulsarFunctionBrokerServerHost = gutil.GetEnvStr("", "localhost")
	PulsarFunctionBrokerWebServiceHost = gutil.GetEnvStr("", "localhost")
	ZkTlsEnable = gutil.GetEnvBool("ZOOKEEPER_TLS_ENABLE", false)
	BkTlsEnable = gutil.GetEnvBool("BOOKKEEPER_TLS_ENABLE", false)
	PulsarTlsEnable = gutil.GetEnvBool("PULSAR_TLS_ENABLE", false)
	PulsarFunctionTlsEnable = gutil.GetEnvBool("PULSAR_FUNCTION_TLS_ENABLE", false)
}
