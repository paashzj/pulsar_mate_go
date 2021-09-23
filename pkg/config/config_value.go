package config

import "github.com/paashzj/gutil"

// pulsar config
var (
	ClusterEnable                 bool
	ClusterInit                   bool
	Function                      bool
	ZkAddress                     string
	ClusterName                   string
	Spec                          string
	PulsarAuthenticationEnabled   bool
	PulsarAuthenticationProviders string
	PulsarAuthorizationEnabled    bool
	PulsarTokenSecretKey          string
	PulsarSuperUserRoles          string
	PulsarClientAuthPlugin        string
)

func init() {
	ClusterEnable = gutil.GetEnvBool("CLUSTER_ENABLE", false)
	ClusterInit = gutil.GetEnvBool("CLUSTER_INIT", false)
	Function = gutil.GetEnvBool("PULSAR_FUNCTION", false)
	ZkAddress = gutil.GetEnvStr("ZK_ADDR", "")
	ClusterName = gutil.GetEnvStr("CLUSTER_NAME", "pulsar")
	Spec = gutil.GetEnvStr("SPEC", "SMALL")
	PulsarAuthenticationEnabled = gutil.GetEnvBool("PULSAR_AUTHENTICATION_ENABLED", false)
	PulsarAuthenticationProviders = gutil.GetEnvStr("PULSAR_AUTHENTICATION_PROVIDERS", "")
	PulsarAuthorizationEnabled = gutil.GetEnvBool("PULSAR_AUTHORIZATION_ENABLED", false)
	PulsarTokenSecretKey = gutil.GetEnvStr("PULSAR_TOKEN_SECRET_KEY", "")
	PulsarSuperUserRoles = gutil.GetEnvStr("PULSAR_SUPER_USER_ROLES", "")
	PulsarClientAuthPlugin = gutil.GetEnvStr("PULSAR_CLIENT_AUTH_PLUGIN", "")
}
