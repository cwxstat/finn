package client

import (
	"reflect"
	"testing"
)

func TestClient(t *testing.T) {
	client := Client()
	if reflect.TypeOf(client).String() != "*finnhub.DefaultApiService" {
		t.Errorf("Client() returned %v, want %v", reflect.TypeOf(client).String(), "*finnhub.DefaultApi")
	}
}
