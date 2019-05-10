package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rhass99/katz/pkg/auth/local"
	"github.com/rhass99/katz/pkg/http/rpc/client"
)


func main() {

	conn, err := client.Run("localhost", "8080")
	if err != nil {
		log.Fatalf("failure while dialing: %v", err)
	}
	defer conn.Close()

	client := local.NewAuthClient(conn)
	response, err := client.SignUp(context.Background(), &local.SignUpRequest{
		Username: "ramihassanein",
		Email: "Alex@gmail.com",
		Password: "xxx",
		Firstname: "Bla Bla",
		Lastname: "Blo Blo",
	})
	fmt.Println(response)

	response, err = client.SignIn(context.Background(), &local.SignInRequest{
		Username: "ramihassanein",
		Password: "xxx",
	})
	fmt.Println(response)
}
