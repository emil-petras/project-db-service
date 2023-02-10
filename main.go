package main

import (
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/emil-petras/db-service/models"
	"github.com/emil-petras/db-service/servers"
	tokenProto "github.com/emil-petras/project-proto/token"
	userProto "github.com/emil-petras/project-proto/user"
)

func main() {
	godotenv.Load(".env")
	port := fmt.Sprintf(":%v", os.Getenv("PORT"))
	err := models.Connect()
	if err != nil {
		logrus.Error(err.Error())
		panic(err)
	}

	listener, err := net.Listen("tcp", port)
	if err != nil {
		logrus.Error(err.Error())
		panic(err)
	}

	s := grpc.NewServer()
	tokenProto.RegisterTokenServiceServer(s, &servers.TokenServer{})
	userProto.RegisterUserServiceServer(s, &servers.UserServer{})
	err = s.Serve(listener)
	if err != nil {
		logrus.Error(err.Error())
		panic(err)
	}
}
