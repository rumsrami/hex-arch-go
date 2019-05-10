package server

import (
	"net"
	"fmt"
)

// Run returns an RPC server listening on port 8080
func Run(protocol string, host string, port string) (net.Listener, error){
	lis, err := net.Listen(protocol, fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		return nil, err
	}
	return lis, nil
}