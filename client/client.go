package client

import (
	"github.com/Finnhub-Stock-API/finnhub-go/v2"
	"os"
)

func Client() *finnhub.DefaultApiService {
	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", os.Getenv("FINNHUB_KEY"))
	client := finnhub.NewAPIClient(cfg).DefaultApi
	return client
}
