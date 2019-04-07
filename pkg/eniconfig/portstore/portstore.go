package portstore

import (
	log "github.com/cihub/seelog"
	"github.com/pkg/errors"
)

var UnknownPortError = errors.New("portstore: unknown port")
var PortUnAvailableError = errors.New("portstore: no port available")
var DoubleReleaseError = errors.New("portstore: the port already released")

type Port struct {
	InUse   bool
	Next    int
	PodName string
}

type PortMap struct {
	freePortsHdr   int // -1
	freePortsTail  int // -1
	Ports          []Port
	freeNumOfPorts int
	startPort      int
	maxPort        int
}

func PortMapInit(maxNumOfPorts int, startPort int) *PortMap {

	log.Debugf("PortMapInit: maxNumOfPorts=%v, startPort=%v", maxNumOfPorts, startPort)
	portMap := &PortMap{}

	portMap.Ports = make([]Port, maxNumOfPorts)
	portMap.freePortsHdr = 0
	portMap.freePortsTail = maxNumOfPorts - 1

	for i, _ := range portMap.Ports {
		portMap.Ports[i].Next = i + 1
		portMap.Ports[i].InUse = false
		log.Debugf("PortMapInit: i=%v, p=%v", i, portMap.Ports[i])
	}

	portMap.freeNumOfPorts = maxNumOfPorts
	portMap.startPort = startPort
	portMap.maxPort = maxNumOfPorts
	return portMap
}

func PortMapAllocPort(portMap *PortMap, podName string) (int, error) {

	log.Debugf("PortMapAllocPort: freeNumOfPort: %v, podName %v",
		portMap.freeNumOfPorts, podName)

	if portMap.freeNumOfPorts <= 0 {
		return 0, PortUnAvailableError
	}

	// take one of free list
	portMap.freeNumOfPorts--

	allocatedPort := portMap.freePortsHdr

	if portMap.freeNumOfPorts == 0 {
		portMap.freePortsHdr = -1
		portMap.freePortsTail = -1
	} else {
		portMap.freePortsHdr = portMap.Ports[allocatedPort].Next
		portMap.Ports[allocatedPort].InUse = true
		portMap.Ports[allocatedPort].Next = -1
		portMap.Ports[allocatedPort].PodName = podName
	}

	log.Debugf("PortMapAllocPort: allocatedPort=%v, freeNumOfPort=%v",
		allocatedPort, portMap.freeNumOfPorts)

	return portMap.startPort + allocatedPort, nil
}

func PortMapReleasePort(portMap *PortMap, port int) error {

	log.Debugf("PortMapReleasePort: port: %v", port)

	freedPort := port - portMap.startPort
	if (freedPort < 0) || (freedPort >= portMap.maxPort) {
		return UnknownPortError
	}

	if portMap.Ports[freedPort].InUse == false {
		return DoubleReleaseError
	}

	if portMap.freePortsHdr == -1 {
		portMap.freePortsHdr = freedPort
	} else {
		portMap.Ports[portMap.freePortsTail].Next = freedPort
	}
	portMap.freePortsTail = freedPort
	portMap.Ports[freedPort].InUse = false
	portMap.Ports[freedPort].Next = -1
	portMap.freeNumOfPorts++

	return nil
}
