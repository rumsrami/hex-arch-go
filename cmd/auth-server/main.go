package main

import (
	"log"

	"github.com/rhass99/katz/pkg/auth/local"
	"github.com/rhass99/katz/pkg/http/rpc/server"
	"github.com/rhass99/katz/pkg/storage/bolt"
)

func main() {
	lis, err := server.Run("tcp", "localhost", "8080")
	if err != nil {
		log.Fatalf("Error listening %v", err)
	}
	s := new(bolt.Storage)
	
	local.NewGRPCService(lis)
}