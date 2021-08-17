package pulsar

import (
	"errors"
	"fmt"
	"github.com/paashzj/gutil"
	"io/ioutil"
	"net"
	"pulsar_mate_go/pkg/config"
	"pulsar_mate_go/pkg/path"
	"pulsar_mate_go/pkg/util"
	"strings"
)

func Config() error {
	if !config.ClusterEnable {
		return nil
	}
	configProp, err := initFromFile(path.PulsarOriginalConfig)
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

func initFromFile(file string) (*gutil.ConfigProperties, error) {
	configProp := gutil.ConfigProperties{}
	configProp.Init()
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(fileBytes), "\n")
	for _, line := range split {
		if strings.HasPrefix(line, "#") {
			continue
		}
		array := strings.Split(line, "=")
		if len(array) != 2 {
			util.Logger().Error(fmt.Sprintf("line error %s", line))
			continue
		}
		configProp.Set(array[0], array[1])
	}
	return &configProp, nil
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
