package main

import (
	"fmt"
	"github.com/micro/go-micro"
	database "github.com/samlina/microshop/user-service/db"
	"github.com/samlina/microshop/user-service/handler"
	pb "github.com/samlina/microshop/user-service/proto/user"
	repository "github.com/samlina/microshop/user-service/repo"
	"github.com/samlina/microshop/user-service/service"
	"log"
)

func main() {

	db, err := database.CreateConnection()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	defer db.Close()

	//和Laravel数据库迁移类似
	//每次启动服务时都会检查，如果数据表不存在则创建，已存在检查是否有修改
	db.AutoMigrate(&pb.User{})

	//初始化 Repo实例用于后续数据库操作
	repo := &repository.UserRepository{db}
	//初始化 token service
	token := &service.TokenService{repo}

	//以下是micro创建微服务流程
	srv := micro.NewService(
		micro.Name("microshop.user.service"),
		micro.Version("v1"), //新增接口版本参数
	)
	srv.Init()

	//注册处理器
	pb.RegisterUserServiceHandler(srv.Server(), &handler.UserService{repo, token})

	//启动用户服务
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
