package handler

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/nnaakkaaii/go-http-server-template/gen/daocore"
	"github.com/nnaakkaaii/go-http-server-template/internal/handler/middleware"
	"github.com/nnaakkaaii/go-http-server-template/internal/service"
	"github.com/nnaakkaaii/go-http-server-template/pkg/echoutil"
	"net/http"
)

func (s *Server) PostLogout(ec echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	user, ok := ec.Get(middleware.ContextUserKey).(daocore.User)
	if !ok {
		return echoutil.ErrInternal(ec, errors.New("cannot dump context into user struct"))
	}

	txn, err := s.db.Begin()
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	defer txn.Rollback()

	msg, err := service.Logout(ctx, txn, &user)
	if msg == nil || err != nil {
		return echoutil.ErrInternal(ec, err)
	}

	if err := txn.Commit(); err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	return ec.JSON(http.StatusOK, *msg)
}
