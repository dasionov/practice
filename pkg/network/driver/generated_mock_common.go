// Automatically generated by MockGen. DO NOT EDIT!
// Source: common.go

package driver

import (
	gomock "github.com/golang/mock/gomock"
	netlink "github.com/vishvananda/netlink"
	v1 "kubevirt.io/api/core/v1"

	cache "practice/pkg/network/cache"
)

// Mock of NetworkHandler interface
type MockNetworkHandler struct {
	ctrl     *gomock.Controller
	recorder *_MockNetworkHandlerRecorder
}

// Recorder for MockNetworkHandler (not exported)
type _MockNetworkHandlerRecorder struct {
	mock *MockNetworkHandler
}

func NewMockNetworkHandler(ctrl *gomock.Controller) *MockNetworkHandler {
	mock := &MockNetworkHandler{ctrl: ctrl}
	mock.recorder = &_MockNetworkHandlerRecorder{mock}
	return mock
}

func (_m *MockNetworkHandler) EXPECT() *_MockNetworkHandlerRecorder {
	return _m.recorder
}

func (_m *MockNetworkHandler) LinkByName(name string) (netlink.Link, error) {
	ret := _m.ctrl.Call(_m, "LinkByName", name)
	ret0, _ := ret[0].(netlink.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkHandlerRecorder) LinkByName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LinkByName", arg0)
}

func (_m *MockNetworkHandler) AddrList(link netlink.Link, family int) ([]netlink.Addr, error) {
	ret := _m.ctrl.Call(_m, "AddrList", link, family)
	ret0, _ := ret[0].([]netlink.Addr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkHandlerRecorder) AddrList(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddrList", arg0, arg1)
}

func (_m *MockNetworkHandler) ReadIPAddressesFromLink(interfaceName string) (string, string, error) {
	ret := _m.ctrl.Call(_m, "ReadIPAddressesFromLink", interfaceName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockNetworkHandlerRecorder) ReadIPAddressesFromLink(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadIPAddressesFromLink", arg0)
}

func (_m *MockNetworkHandler) RouteList(link netlink.Link, family int) ([]netlink.Route, error) {
	ret := _m.ctrl.Call(_m, "RouteList", link, family)
	ret0, _ := ret[0].([]netlink.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkHandlerRecorder) RouteList(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RouteList", arg0, arg1)
}

func (_m *MockNetworkHandler) LinkDel(link netlink.Link) error {
	ret := _m.ctrl.Call(_m, "LinkDel", link)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) LinkDel(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LinkDel", arg0)
}

func (_m *MockNetworkHandler) ParseAddr(s string) (*netlink.Addr, error) {
	ret := _m.ctrl.Call(_m, "ParseAddr", s)
	ret0, _ := ret[0].(*netlink.Addr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkHandlerRecorder) ParseAddr(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ParseAddr", arg0)
}

func (_m *MockNetworkHandler) StartDHCP(nic *cache.DHCPConfig, bridgeInterfaceName string, dhcpOptions *v1.DHCPOptions) error {
	ret := _m.ctrl.Call(_m, "StartDHCP", nic, bridgeInterfaceName, dhcpOptions)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) StartDHCP(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartDHCP", arg0, arg1, arg2)
}

func (_m *MockNetworkHandler) HasIPv4GlobalUnicastAddress(interfaceName string) (bool, error) {
	ret := _m.ctrl.Call(_m, "HasIPv4GlobalUnicastAddress", interfaceName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkHandlerRecorder) HasIPv4GlobalUnicastAddress(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HasIPv4GlobalUnicastAddress", arg0)
}

func (_m *MockNetworkHandler) HasIPv6GlobalUnicastAddress(interfaceName string) (bool, error) {
	ret := _m.ctrl.Call(_m, "HasIPv6GlobalUnicastAddress", interfaceName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkHandlerRecorder) HasIPv6GlobalUnicastAddress(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HasIPv6GlobalUnicastAddress", arg0)
}

func (_m *MockNetworkHandler) IsIpv4Primary() (bool, error) {
	ret := _m.ctrl.Call(_m, "IsIpv4Primary")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkHandlerRecorder) IsIpv4Primary() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsIpv4Primary")
}
