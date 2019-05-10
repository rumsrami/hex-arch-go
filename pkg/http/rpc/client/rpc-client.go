package client

import (
	"fmt"
	"google.golang.org/grpc"
)

// Run returns an RPC client connection that can send messages to the RPC server
func Run(host string, port string) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", host, port), opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}