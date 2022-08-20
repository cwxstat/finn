package earnings

import (
	"fmt"
	"github.com/cwxstat/finn/client"
	"testing"
)

func TestEarnings(t *testing.T) {

	earningsCalendar, _, err := Earnings(client.Client())
	fmt.Println(earningsCalendar, err)

}
