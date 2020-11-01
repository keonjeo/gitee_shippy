package main

import (
	"fmt"

	"user_srv/db"
	pb "user_srv/proto/user"

	"github.com/micro/go-micro/v2"
)

func main() {
	defer db.Client().Close()

	repo := &UserRepository{db.Client()}

	tokenService := &TokenService{repo}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	pb.RegisterUserServiceHandler(srv.Server(), &userServiceHandler{repo, tokenService})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
