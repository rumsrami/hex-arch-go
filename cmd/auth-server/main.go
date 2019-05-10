package main

import (
	"log"

	"github.com/rumsrami/hex-arch-go/pkg/auth/local"
	"github.com/rumsrami/hex-arch-go/pkg/http/rpc/server"
	_"github.com/rumsrami/hex-arch-go/pkg/storage/bolt"
)

func main() {
	lis, err := server.Run("tcp", "localhost", "8080")
	if err != nil {
		log.Fatalf("Error listening %v", err)
	}
	//s := new(bolt.Storage)
	
	local.NewGRPCService(lis)
}