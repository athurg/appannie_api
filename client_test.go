package appannie

import (
	"flag"
	"os"
	"testing"
)

var testClient *Client

func TestMain(m *testing.M) {
	flag.Parse()
	testClient = New(os.Getenv("API_KEY"), os.Getenv("APPANNIE_API_VERION"))
	os.Exit(m.Run())
}
