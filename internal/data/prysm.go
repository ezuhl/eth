package data

import (
	"context"
	"github.com/davidebianchi/go-jsonclient"
	"github.com/ezuhl/eth/internal/data/model"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"time"
)

type PrysmClient interface {
	GetChainedHead() (*model.ChainHeadResponse, error)
}

type prysmClient struct {
	client *jsonclient.Client
}

func NewPrysmClient() PrysmClient {
	prysmHost := os.Getenv("PRYSM_JSON_HOST")

	opts := jsonclient.Options{
		BaseURL: prysmHost,
	}
	client, err := jsonclient.New(opts)
	if err != nil {
		panic("Error creating client")
	}

	p := &prysmClient{client}

	return p
}

func (p *prysmClient) GetChainedHead() (*model.ChainHeadResponse, error) {

	var data = map[string]interface{}{
		"some": "json format",
		"foo":  "bar",
		"that": float64(3),
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	req, err := p.client.NewRequestWithContext(ctx, http.MethodGet, "beacon/chainhead", data)
	if err != nil {
		panic("Error creating request")
	}

	type Response struct {
		my string
	}
	v := model.ChainHeadResponse{}
	// server response is: {"my": "data"}
	response, err := p.client.Do(req, &v)
	if err != nil {
		return nil, errors.Wrap(err, "could not query server")
	}

	if response.StatusCode != 200 {
		return nil, errors.New("not successful")
	}

	return &v, nil

}
