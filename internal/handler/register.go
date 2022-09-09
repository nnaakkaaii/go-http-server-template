package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/nnaakkaaii/go-http-server-template/gen/api"
	"github.com/nnaakkaaii/go-http-server-template/internal/handler/validator"
	"github.com/nnaakkaaii/go-http-server-template/internal/service"
	"github.com/nnaakkaaii/go-http-server-template/pkg/echoutil"
)

func (s *Server) PostRegister(ec echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req := &api.User{}
	if err := ec.Bind(req); err != nil {
		return echoutil.ErrBadRequest(ec, err)
	}

	if err := (&validator.User{req}).Register(); err != nil {
		return echoutil.ErrBadRequest(ec, err)
	}

	txn, err := s.db.Begin()
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	defer txn.Rollback()

	msg, err := service.Register(ctx, txn, req)
	if msg == nil || err != nil {
		return echoutil.ErrInternal(ec, err)
	}

	if err := txn.Commit(); err != nil {
		return echoutil.ErrInternal(ec, err)
	}

	return ec.JSON(http.StatusOK, msg)
}
