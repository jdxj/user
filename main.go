package main

import (
	"github.com/jdxj/logger"
	"github.com/jdxj/user/handler"
	user "github.com/jdxj/user/proto"
	"github.com/micro/micro/v3/service"
)

func main() {
	//err := config.Init("/home/jdxj/workspace/user/config.yaml")
	//if err != nil {
	//	panic(err)
	//}
	//
	//logger.NewPathMode(config.Log().Path, config.Mode())
	//
	//dbCfg := config.DB()
	//err = model.InitDB(dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.DBName)
	//if err != nil {
	//	panic(err)
	//}
	// Create service
	srv := service.New(
		service.Name("user"),
		service.Version("latest"),
	)

	// Register handler
	user.RegisterUserHandler(srv.Server(), new(handler.User))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Errorf("Run: %s", err)
	}
}
