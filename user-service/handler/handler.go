package handler

import (
	"context"
	"errors"
	pb "github.com/samlina/microshop/user-service/proto/user"
	"github.com/samlina/microshop/user-service/repo"
	"github.com/samlina/microshop/user-service/service"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserService struct {
	Repo  repo.Repository
	Token service.Authable
}

func (srv *UserService) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	log.Println("Logging in with:", req.Email, req.Password)
	//获取用户信息
	user, err := srv.Repo.GetByEmail(req.Email)
	log.Println(user)
	if err != nil {
		return err
	}

	//校验用户输入密码是否与数据库存储密码匹配
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	//生成JWT token
	token, err := srv.Token.Encode(user)
	if err != nil {
		return err
	}

	res.Token = token
	return nil
}

func (srv *UserService) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	//校验用户请求中的token信息是否有效
	cliams, err := srv.Token.Decode(req.Token)

	if err != nil {
		return err
	}

	if cliams.User.Id == "" {
		return errors.New("无效的用户")
	}

	res.Valid = true

	return nil
}

func (srv *UserService) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := srv.Repo.Get(req.Id)
	if err != nil {
		return err
	}

	res.User = user
	return nil
}

func (srv *UserService) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := srv.Repo.GetAll()
	if err != nil {
		return err
	}

	res.Users = users
	return nil
}

func (srv *UserService) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	//对密码进行哈希加密
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPass)
	if err := srv.Repo.Create(req); err != nil {
		return err
	}
	res.User = req
	return nil
}
