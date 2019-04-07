package portstore

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPortMapInit(t *testing.T) {

	port := PortMapInit(10, 100)

	assert.Equal(t, port.freeNumOfPorts, 10)
	assert.Equal(t, port.startPort, 100)
	assert.Equal(t, port.maxPort, 10)
	assert.Equal(t, port.freePortsHdr, 0)
	assert.Equal(t, port.freePortsTail, 9)

	fmt.Printf("port: %v\n", port)
}

func TestPortMapAllocPort(t *testing.T) {

	portMap := PortMapInit(10, 100)

	port, err := PortMapAllocPort(portMap, "pod1")

	fmt.Printf("TestPortMapAllocPort: port: %v, err: %v\n", port, err)

	assert.Equal(t, portMap.freeNumOfPorts, 9)
	assert.Equal(t, portMap.freePortsHdr, 1)
	assert.Equal(t, portMap.freePortsTail, 9)

}

func TestPortMapAllocPortErr(t *testing.T) {
	portMap := PortMapInit(2, 100)

	port, err := PortMapAllocPort(portMap, "pod1")
	assert.NoError(t, err)
	fmt.Printf("TestPortMapAllocPort: port: %v, err: %v\n", port, err)

	port, err = PortMapAllocPort(portMap, "pod1")
	fmt.Printf("TestPortMapAllocPort: port: %v, err: %v\n", port, err)
	assert.NoError(t, err)

	port, err = PortMapAllocPort(portMap, "pod1")
	fmt.Printf("TestPortMapAllocPort: port: %v, err: %v\n", port, err)
	assert.Error(t, err)
}

func TestPortMapReleasePort(t *testing.T) {
	portMap := PortMapInit(10, 100)
	port, err := PortMapAllocPort(portMap, "pod1")
	fmt.Printf("TestPortMapReleasePort: port: %v, err: %v\n", port, err)

	port, err = PortMapAllocPort(portMap, "pod2")
	fmt.Printf("TestPortMapReleasePort: port: %v, err: %v\n", port, err)

	fmt.Printf("TestPortMapReleasePort: portMap: %v\n", portMap)

	err = PortMapReleasePort(portMap, port)
	fmt.Printf("TestPortMapReleasePort: portMap: %v\n", portMap)

	assert.Equal(t, portMap.freeNumOfPorts, 9)
}

func TestPortMapReleasePortErr(t *testing.T) {
	portMap := PortMapInit(10, 100)
	port, err := PortMapAllocPort(portMap, "pod1")
	fmt.Printf("TestPortMapReleasePort: port: %v, err: %v\n", port, err)

	port, err = PortMapAllocPort(portMap, "pod2")
	fmt.Printf("TestPortMapReleasePort: port: %v, err: %v\n", port, err)

	fmt.Printf("TestPortMapReleasePort: portMap: %v\n", portMap)

	err = PortMapReleasePort(portMap, port)
	fmt.Printf("TestPortMapReleasePort: portMap: %v\n", portMap)
	assert.NoError(t, err)

	err = PortMapReleasePort(portMap, port)
	fmt.Printf("TestPortMapReleasePort: portMap: %v error: %v\n", portMap, err)
	assert.Error(t, err)

	err = PortMapReleasePort(portMap, 200)
	fmt.Printf("TestPortMapReleasePort: portMap: %v, error: %v\n", portMap, err)
	assert.Error(t, err)
}
