package earnings

import (
	"context"
	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
	"net/http"
)

func Earnings(client *finnhub.DefaultApiService) (finnhub.EarningsCalendar, *http.Response, error) {
	earningsCalendar, resp, err := client.EarningsCalendar(context.Background()).
		From("2022-08-15").To("2022-08-30").Execute()
	return earningsCalendar, resp, err
}
