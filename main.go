package main

import (
	"github.com/jdxj/user/config"
	"github.com/jdxj/user/handler"
	"github.com/jdxj/user/logger"
	"github.com/jdxj/user/model"
	"github.com/micro/micro/v3/service"
)

func main() {
	err := config.Init("config.yaml")
	if err != nil {
		panic(err)
	}

	logger.Init(config.Log().Path, config.Mode())

	dbCfg := config.DB()
	err = model.InitDB(dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.DBName)
	if err != nil {
		panic(err)
	}
	// Create service
	srv := service.New(
		service.Name("user"),
		//service.Version("latest"),
	)

	// Register handler
	//pb.RegisterUserHandler(srv.Server(), new(handler.User))
	srv.Handle(new(handler.User))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Error("Run: %s", err)
	}
}
