package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/nnaakkaaii/go-http-server-template/gen/api"
	"github.com/nnaakkaaii/go-http-server-template/internal/handler/middleware"
	"github.com/nnaakkaaii/go-http-server-template/internal/handler/validator"
	"github.com/nnaakkaaii/go-http-server-template/internal/service"
	"github.com/nnaakkaaii/go-http-server-template/pkg/echoutil"
	"net/http"
)

func (s *Server) PostLogin(ec echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req := &api.User{}
	if err := ec.Bind(&req); err != nil {
		return echoutil.ErrBadRequest(ec, err)
	}

	if err := (&validator.User{req}).Login(); err != nil {
		return echoutil.ErrBadRequest(ec, err)
	}

	txn, err := s.db.Begin()
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	defer txn.Rollback()

	msg, id, err := service.Login(ctx, txn, req)
	if msg == nil || err != nil {
		return echoutil.ErrInternal(ec, err)
	}

	if err := txn.Commit(); err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	ec, err = middleware.SetCookie(ec, id)
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}

	return ec.JSON(http.StatusOK, msg)
}
