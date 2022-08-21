package interval

import (
	"fmt"
	"testing"
	"time"
)

func TestInterval_Past(t *testing.T) {

	days := 10
	expected := time.Now().AddDate(0, 0, -days).Format("2006-01-02")
	i := NewInterval()
	i.Past(days)
	if i.StartDate() != expected {
		t.Errorf("Expected %s, got %s", expected, i.StartDate())
	}
	fmt.Println(i.StartDate())
}
