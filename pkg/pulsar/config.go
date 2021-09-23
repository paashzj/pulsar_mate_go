package pulsar

import (
	"errors"
	"fmt"
	"net"
	"pulsar_mate_go/pkg/config"
	"pulsar_mate_go/pkg/util"
)

func Config() error {
	util.Logger().Info("begin to generate or refresh config")
	err := configClient()
	if err != nil {
		return err
	}
	if !config.ClusterEnable {
		if config.Function {
			return configFunctionStandalone()
		}
		return configBrokerStandalone()
	}
	if config.Function {
		return configFunctionCluster()
	} else {
		return configBrokerCluster()
	}
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
