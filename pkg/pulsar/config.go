package pulsar

import (
	"errors"
	"fmt"
	"github.com/paashzj/gutil"
	"net"
	"pulsar_mate_go/pkg/config"
	"pulsar_mate_go/pkg/path"
)

func Config() error {
	if config.Function {
		return configFunction()
	} else {
		return configBroker()
	}
}

func configBroker() error {
	if !config.ClusterEnable {
		return nil
	}
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
	return configProp.Write(path.PulsarConfig)
}

// GetInterfaceIpv4Addr useful links:
// https://stackoverflow.com/questions/27410764/dial-with-a-specific-address-interface-golang
// https://stackoverflow.com/questions/22751035/golang-distinguish-ipv4-ipv6
func GetInterfaceIpv4Addr(interfaceName string) (addr string, err error) {
	var (
		ief      *net.Interface
		addrs    []net.Addr
		ipv4Addr net.IP
	)
	if ief, err = net.InterfaceByName(interfaceName); err != nil { // get interface
		return
	}
	if addrs, err = ief.Addrs(); err != nil { // get addresses
		return
	}
	for _, addr := range addrs { // get ipv4 address
		if ipv4Addr = addr.(*net.IPNet).IP.To4(); ipv4Addr != nil {
			break
		}
	}
	if ipv4Addr == nil {
		return "", errors.New(fmt.Sprintf("interface %s don't have an ipv4 address\n", interfaceName))
	}
	return ipv4Addr.String(), nil
}
