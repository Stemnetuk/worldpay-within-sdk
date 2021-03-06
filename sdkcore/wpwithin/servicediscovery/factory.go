package servicediscovery

import "net"

// NewScanner creates a new instance Scanner
func NewScanner(port, stepSleep int) (Scanner, error) {

	result := &scannerImpl{
		run:       false,
		stepSleep: stepSleep,
	}

	comm, err := NewUDPComm(port, "udp4")

	if err != nil {

		return nil, err
	}

	result.comm = comm

	return result, nil
}

// NewBroadcaster create a new instance of Broadcaster
func NewBroadcaster(host string, port int, stepSleep int) (Broadcaster, error) {

	result := &broadcasterImpl{

		stepSleep: stepSleep,
		run:       false,
		host:      net.IPv4bcast.To4().String(),
		port:      port,
	}

	comm, err := NewUDPComm(port, "udp4")

	if err != nil {

		return nil, err
	}

	result.comm = comm

	return result, nil
}
