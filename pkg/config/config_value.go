package config

import "github.com/paashzj/gutil"

// pulsar config
var (
	ClusterEnable bool
	ClusterInit   bool
	Function      bool
	ZkAddress     string
	ClusterName   string
)

func init() {
	ClusterEnable = gutil.GetEnvBool("CLUSTER_ENABLE", false)
	ClusterInit = gutil.GetEnvBool("CLUSTER_INIT", false)
	Function = gutil.GetEnvBool("PULSAR_FUNCTION", false)
	ZkAddress = gutil.GetEnvStr("ZK_ADDR", "")
	ClusterName = gutil.GetEnvStr("CLUSTER_NAME", "pulsar")
}
