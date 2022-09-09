package service

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"

	"github.com/nnaakkaaii/go-http-server-template/gen/api"
	"github.com/nnaakkaaii/go-http-server-template/gen/daocore"
	"github.com/nnaakkaaii/go-http-server-template/pkg/crypto"
)

func Register(ctx context.Context, txn *sql.Tx, user *api.User) (*api.Message, error) {
	id := uuid.NewString()
	u := &daocore.User{
		ID:        id,
		FirstName: *user.FirstName,
		LastName:  *user.LastName,
		Email:     *user.Email,
		Password:  crypto.Encrypto(*user.Password),
	}
	if err := daocore.InsertUser(ctx, txn, []*daocore.User{u}); err != nil {
		return nil, err
	}
	msg := "successfully registered"
	return &api.Message{&msg}, nil
}

func Login(ctx context.Context, txn *sql.Tx, user *api.User) (*api.Message, string, error) {
	u, err := daocore.SelectOneUserByEmail(ctx, txn, user.Email)
	if err != nil {
		return nil, "", err
	}
	// パスワード認証
	pw, err := crypto.Decrypto(u.Password)
	if err != nil {
		return nil, "", err
	}
	if user.Password == nil || pw != *user.Password {
		return nil, "", err
	}
	// セッション追加
	if err := daocore.DeleteOneUserSessionByUserID(ctx, txn, &u.ID); err != nil {
		return nil, "", err
	}
	id := uuid.New().String()
	sess := &daocore.UserSession{
		ID:     id,
		UserID: u.ID,
	}
	if err := daocore.InsertUserSession(ctx, txn, []*daocore.UserSession{sess}); err != nil {
		return nil, "", err
	}

	msg := "successfully logged in"
	return &api.Message{&msg}, id, nil
}

func Logout(ctx context.Context, txn *sql.Tx, user *daocore.User) (*api.Message, error) {
	log.Println("logout")
	if err := daocore.DeleteOneUserSessionByUserID(ctx, txn, &user.ID); err != nil {
		return nil, err
	}
	msg := "successfully logged out"
	return &api.Message{&msg}, nil
}
