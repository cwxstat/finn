package earnings

import (
	"context"
	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
	"net/http"
)

func Earnings(client *finnhub.DefaultApiService, from, to string) (finnhub.EarningsCalendar, *http.Response, error) {
	earningsCalendar, resp, err := client.EarningsCalendar(context.Background()).
		From(from).To(to).Execute()
	return earningsCalendar, resp, err
}
