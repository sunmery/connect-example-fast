package service

import (
	v1 "connect-go-example/api/user/v1"
	"connect-go-example/internal/biz"

	"connectrpc.com/connect"
	"go.uber.org/fx"
)

var Module = fx.Module("service",
	fx.Provide(NewUserService),
)

// UserService 实现 Connect 服务
type UserService struct {
	uc *biz.UserUseCase
}

// 显式接口检查
var _ userv1connect.UserServiceHandler = (*UserService)(nil)

func NewUserService(uc *biz.UserUseCase) userv1connect.UserServiceHandler {
	return &UserService{
		uc: uc,
	}
}

func (s *UserService) SignIn(ctx context.Context, c *connect.Request[v1.SignInRequest]) (*connect.Response[v1.SignInResponse], error) {
	_, err := s.uc.SignIn(
		ctx,
		biz.SignInRequest{},
	)
	if err != nil {
		return nil, err
	}

	response := &v1.SignInResponse{}

	return connect.NewResponse(response), nil
}
