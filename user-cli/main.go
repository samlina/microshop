package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	pb "user-service/proto/user"
)

func main() {


	service := micro.NewService(
		micro.Flags(
				cli.StringFlag{
					Name: "name",
					Usage: "Your Name",
				},
				cli.StringFlag{
					Name: "email",
					Usage: "Yout Email",
				},
				cli.StringFlag{
					Name: "password",
					Usage: "Your Password",
				},
			),
		)


	client :=
}
