package main

import (
	"github.com/jdxj/user/config"
	"github.com/jdxj/user/handler"
	"github.com/jdxj/user/logger"
	pb "github.com/jdxj/user/proto"

	"github.com/micro/micro/v3/service"
)

func main() {
	config.InitConfig("config.yaml")
	// Create service
	srv := service.New(
		service.Name("user"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterUserHandler(srv.Server(), new(handler.User))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Error("Run: %s", err)
	}
}
