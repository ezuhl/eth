package data

import (
	"context"
	"github.com/ezuhl/eth/internal/data/model"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Account interface {
	GetUser(id int64) (user *model.User, err error)
	GetUserByLogin(username string, hashedPassword string) (user *model.User, err error)
	GetApiKey(apiKey int64) (api *model.Api, err error)
	GetApiKeyByUserId(userId int64) (api *model.Api, err error)
	CreateUser(ctx context.Context, userModel *model.User) error
	CreateApiKey(ctx context.Context, apiModel *model.Api) error
	CreateTransaction(ctx context.Context, transactionModel *model.Transaction) error
}

type account struct {
	Db *sqlx.DB
}

func NewAccount(db *sqlx.DB) Account {

	a := &account{db}

	return a

}

func (a *account) GetUser(id int64) (user *model.User, err error) {

	user = &model.User{}
	err = a.Db.Get(user, "SELECT * FROM account.user WHERE user_id =$1", id)

	//if err == sql.ErrNoRows{
	//	return nil,nil
	//}

	return
}

func (a *account) GetUserByLogin(username string, hashedPassword string) (user *model.User, err error) {

	user = &model.User{}
	err = a.Db.Get(user, "SELECT * FROM account.user WHERE user_id =$1 and password = $2", username, hashedPassword)

	return
}

func (a *account) GetApiKey(apiKey int64) (api *model.Api, err error) {

	api = &model.Api{}
	err = a.Db.Get(api, "SELECT * FROM account.apikey WHERE api_key =$1", apiKey)

	return
}

func (a *account) GetApiKeyByUserId(userId int64) (api *model.Api, err error) {

	api = &model.Api{}
	err = a.Db.Get(api, "SELECT * FROM account.apikey WHERE user_id  =$1", userId)

	return
}

func (a *account) CreateUser(ctx context.Context, userModel *model.User) error {

	_, err := a.Db.ExecContext(ctx, "INSERT INTO account.user (user_id, username, password) VALUES ($1, $2, $3)", userModel.UserID, userModel.Username, userModel.Password)

	if err != nil {
		return errors.Wrap(err, "CreateUser=>ExecContext")
	}

	return nil

}

func (a *account) CreateApiKey(ctx context.Context, apiModel *model.Api) error {

	_, err := a.Db.ExecContext(ctx, "INSERT INTO account.apikey (user_id, api_key) VALUES ($1, $2)", apiModel.UserID, apiModel.ApiKey)

	if err != nil {
		return errors.Wrap(err, "CreateApiKey=>ExecContext")
	}
	return nil

}

func (a *account) CreateTransaction(ctx context.Context, transactionModel *model.Transaction) error {

	_, err := a.Db.ExecContext(ctx, "INSERT INTO account.transactions (user_id, action) VALUES ($1, $2)", transactionModel.UserID, transactionModel.Action)

	if err != nil {
		return errors.Wrap(err, "CreateTransaction=>ExecContext")
	}
	return nil

}
