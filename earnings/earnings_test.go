package earnings

import (
	"fmt"
	"github.com/cwxstat/finn/client"
	"testing"
	"time"
)

func TestEarnings(t *testing.T) {

	from := time.Now().AddDate(0, 0, -10).Format("2006-01-02")
	to := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	earningsCalendar, _, err := Earnings(client.Client(), from, to)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(earningsCalendar, err)

}
