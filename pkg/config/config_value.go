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
	PulsarAuthenticationEnabled                               bool
	PulsarAuthenticationProviders                             string
	PulsarAuthorizationEnabled                                bool
	PulsarTokenSecretKey                                      string
	PulsarSuperUserRoles                                      string
	PulsarClientAuthPlugin                                    string
	PulsarAllowAutoTopicCreation                              bool
	PulsarBrokerDeleteInactivePartitionedTopicMetadataEnabled bool
)

func init() {
	ClusterEnable = gutil.GetEnvBool("CLUSTER_ENABLE", false)
	ClusterInit = gutil.GetEnvBool("CLUSTER_INIT", false)
	Role = gutil.GetEnvStr("PULSAR_ROLE", "all")
	ZkAddress = gutil.GetEnvStr("ZK_ADDR", "")
	ClusterName = gutil.GetEnvStr("CLUSTER_NAME", "pulsar")
	Spec = gutil.GetEnvStr("SPEC", "SMALL")
	PulsarAuthenticationEnabled = gutil.GetEnvBool("PULSAR_AUTHENTICATION_ENABLED", false)
	PulsarAuthenticationProviders = gutil.GetEnvStr("PULSAR_AUTHENTICATION_PROVIDERS", "")
	PulsarAuthorizationEnabled = gutil.GetEnvBool("PULSAR_AUTHORIZATION_ENABLED", false)
	PulsarTokenSecretKey = gutil.GetEnvStr("PULSAR_TOKEN_SECRET_KEY", "")
	PulsarSuperUserRoles = gutil.GetEnvStr("PULSAR_SUPER_USER_ROLES", "")
	PulsarClientAuthPlugin = gutil.GetEnvStr("PULSAR_CLIENT_AUTH_PLUGIN", "")
	PulsarAllowAutoTopicCreation = gutil.GetEnvBool("PULSAR_ALLOW_AUTO_TOPIC_CREATION", true)
	PulsarBrokerDeleteInactivePartitionedTopicMetadataEnabled = gutil.GetEnvBool("PULSAR_BROKER_DELETE_INACTIVE_PARTITIONED_TOPIC_METADATA_ENABLED", false)
}
