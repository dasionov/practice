package dhcp

import (
	"github.com/vishvananda/netlink"

	v1 "kubevirt.io/api/core/v1"

	"practice/pkg/network/cache"
	netdriver "practice/pkg/network/driver"
	virtnetlink "practice/pkg/network/link"
)

type cacheCreator interface {
	New(filePath string) *cache.Cache
}

type BridgeConfigGenerator struct {
	handler          netdriver.NetworkHandler
	podInterfaceName string
	cacheCreator     cacheCreator
	launcherPID      string
	vmiSpecIfaces    []v1.Interface
	vmiSpecIface     *v1.Interface
	subdomain        string
}

func (d *BridgeConfigGenerator) Generate() (*cache.DHCPConfig, error) {
	dhcpConfig, err := cache.ReadDHCPInterfaceCache(d.cacheCreator, d.launcherPID, d.podInterfaceName)
	if err != nil {
		return nil, err
	}

	if dhcpConfig.IPAMDisabled {
		return dhcpConfig, nil
	}

	dhcpConfig.Name = d.podInterfaceName

	fakeBridgeIP := virtnetlink.GetFakeBridgeIP(d.vmiSpecIfaces, d.vmiSpecIface)
	fakeServerAddr, _ := netlink.ParseAddr(fakeBridgeIP)
	dhcpConfig.AdvertisingIPAddr = fakeServerAddr.IP

	newPodNicName := virtnetlink.GenerateNewBridgedVmiInterfaceName(d.podInterfaceName)
	podNicLink, err := d.handler.LinkByName(newPodNicName)
	if err != nil {
		return nil, err
	}
	dhcpConfig.Mtu = uint16(podNicLink.Attrs().MTU)
	dhcpConfig.Subdomain = d.subdomain

	return dhcpConfig, nil
}
