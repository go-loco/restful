package rest_test

import (
	"net/http"
	"testing"

	"github.com/go-loco/restful/rest"
)

func TestMockup(t *testing.T) {

	defer rest.StopMockupServer()
	rest.StartMockupServer()

	myURL := "http://mytest.com/foo?val1=1&val2=2#fragment"

	myHeaders := make(http.Header)
	myHeaders.Add("Hello", "world")

	mock := rest.Mock{
		URL:          myURL,
		HTTPMethod:   http.MethodGet,
		ReqHeaders:   myHeaders,
		RespHTTPCode: http.StatusOK,
		RespBody:     "foo",
	}

	rest.AddMockups(&mock)

	v := rest.Get(myURL)
	if v.String() != "foo" {
		t.Fatal("Mockup Fail!")
	}

}
