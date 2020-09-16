package handler

import (
	"context"
	"database/sql"

	"github.com/jdxj/user/logger"
	"github.com/jdxj/user/model"
	user "github.com/jdxj/user/proto"
)

type User struct{}

func (u *User) Login(ctx context.Context, req *user.RequestLogin, resp *user.ResponseLogin) error {
	userInfo, err := model.LoginCheck(req.Name, req.Password)
	if err == nil {
		resp.Message = "ok"
		resp.UserId = int64(userInfo.ID)
		return nil
	}

	if err != sql.ErrNoRows {
		logger.Error("LoginCheck: %s", err)
		resp.Code = 101
		resp.Message = "internal error"
		return nil
	}

	resp.Code = 102
	resp.Message = "name or password error"
	return nil
}
