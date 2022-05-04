package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/eth/internal/data"
	"github.com/eth/internal/data/model"
	"github.com/eth/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

type ethHandler struct {
	account   data.Account
	prysm     data.PrysmClient
	authToken (*jwtauth.JWTAuth)
}

func NewEthHandler(db *sqlx.DB) routes.Handler {

	e := &ethHandler{}

	a := data.NewAccount(db)
	e.account = a

	p := data.NewPrysmClient()
	e.prysm = p
	return e

}
func (e *ethHandler) SetAuthToken(authToken *jwtauth.JWTAuth) {
	e.authToken = authToken
}

func (e *ethHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := CreateAccountRequest{}
	resp := CreateAccountResponse{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := &model.User{}
	user.UserID = int64(seededRand.Int())
	user.Username = req.Username
	user.SetPassword(req.Password)

	err = e.account.CreateUser(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, tokenString, err := e.authToken.Encode(map[string]interface{}{"user_id": fmt.Sprintf("%d", user.UserID)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Token = tokenString

	respBytes, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(respBytes)

}
func (e *ethHandler) GetApiKey(w http.ResponseWriter, r *http.Request) {
	resp := &ApiKeyResponse{}
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, ok := claims["user_id"]; !ok {
		http.Error(w, "token is not valid", http.StatusBadRequest)
		return
	}

	log.Println("user id ", claims["user_id"])
	userId, err := strconv.ParseInt(claims["user_id"].(string), 10, 64)
	if err != nil {
		err = errors.Wrap(err, "could not parse user id claim")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err := e.account.GetUser(userId)
	if err != nil {
		err = errors.Wrapf(err, "could not find user with id  %d", userId)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userApi, err := e.account.GetApiKeyByUserId(userId)
	if err == nil {
		resp.ApiKey = userApi.ApiKey
		respBytes, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respBytes)
		return
	}

	api := &model.Api{}
	api.UserID = user.UserID
	api.ApiKey = seededRand.Int()

	err = e.account.CreateApiKey(r.Context(), api)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.ApiKey = api.ApiKey
	respBytes, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(respBytes)

}
func (e *ethHandler) GetAverageHeight(w http.ResponseWriter, r *http.Request) {

	resp := &AverageHeightResponse{}
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, ok := claims["user_id"]; !ok {
		http.Error(w, "token is not valid", http.StatusBadRequest)
		return
	}

	log.Println("user id ", claims["user_id"])
	userId, err := strconv.ParseInt(claims["user_id"].(string), 10, 64)
	if err != nil {
		err = errors.Wrap(err, "could not parse user id claim")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	apiKeyString := chi.URLParam(r, "token")
	if len(apiKeyString) == 0 {
		err = errors.New("not a valid api key")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	reqApiKey, err := strconv.ParseInt(apiKeyString, 10, 64)
	if err != nil {
		err = errors.Wrap(err, "could not parse user id claim")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userApi, err := e.account.GetApiKeyByUserId(userId)
	if err != nil {
		err = errors.Wrap(err, "could not find api key")
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	if int64(userApi.ApiKey) != reqApiKey {
		err = errors.Wrap(err, "bad access key")
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	chainHead, err := e.prysm.GetChainedHead()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	head, _ := chainHead.GetHeadSlot()
	finalized, _ := chainHead.GetFinalizedSlot()
	if finalized > 0 {
		secDifference := (head - finalized) * 12
		minuteRate := secDifference / 60 //
		resp.BlocksPerMinuteAVG = minuteRate
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(respBytes)
}
